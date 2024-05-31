package services

import (
	"context"
	"strconv"
	"time"
	"trainee3/database"
	apicode "trainee3/lib/apiCode"
	"trainee3/model/entity/mongo/mongo_trainee3"
	"trainee3/model/entity/mysql/mysql_trainee3"
	model_redis "trainee3/model/entity/redis"
	"trainee3/model/input"
	"trainee3/repository"
	"trainee3/tools/weighted_random"

	"github.com/bsm/redislock"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

const memberSignInLockKey = "memberSignInLockKey_"

type IActivity interface {
	SignIn(value input.MemberIdRequest) (string, apicode.Code)
}
type activity struct {
	repo *repository.Repo
	db   *database.Database
}

func NewActivity(repo *repository.Repo, db *database.Database) IActivity {
	return &activity{
		repo: repo,
		db:   db,
	}
}

func (s *activity) SignIn(value input.MemberIdRequest) (string, apicode.Code) {
	//redis lock
	locker := redislock.New(s.db.Redis)
	ctx := context.Background()
	key := memberSignInLockKey + strconv.Itoa(int(value.MemberID))
	lock, err := locker.Obtain(ctx, key, 100*time.Millisecond, nil)
	if err == redislock.ErrNotObtained {
		logrus.Error("Could not obtain lock!")
		return "", apicode.UnknownError

	} else if err != nil {
		logrus.Error("Could not obtain lock:", err)
		return "", apicode.UnknownError

	}

	defer lock.Release(ctx)

	// 檢查是否24小時內已經簽到
	var memberSignInRecord *mongo_trainee3.DS_MemberSignInRecord
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// $gte >= $lt <
	err1 := s.repo.Mongo.FindOne(s.db.Mongo.GetDB(), mongo_trainee3.CollectionName, bson.M{
		"memberId": value.MemberID,
		"recordTime": bson.M{
			"$gte": startOfDay,
			"$lt":  startOfDay.Add(24 * time.Hour),
		},
	}, &memberSignInRecord)
	if err1 != nil {
		logrus.Error("SignIn MongoDB FindOne error:", err1)
		return "", apicode.UnknownError
	}
	if memberSignInRecord != nil {
		logrus.Info("Already Sign In")
		return "", apicode.AlreadySignIn
	}

	// 獲取活動資訊
	var DS_Activity model_redis.DSInRedis
	if err := s.repo.Redis.Get(s.db.Redis, model_redis.DSActivityTableName, &DS_Activity); err != nil {
		logrus.Error("SignIn Redis-DS_Activity Get error:", err)
		return "", apicode.UnknownError
	}
	if DS_Activity.ID == 0 {
		logrus.Error("SignIn Redis-DS_Activity is not open")
		return "", apicode.ActivityNotOpen
	}

	// 隨機獲取獎勵
	weightedRandList := lo.Map(DS_Activity.Rewards, func(item mysql_trainee3.DSReward, _ int) weighted_random.WeightedRandom {
		return weighted_random.WeightedRandom{
			Object: item,
			Weight: item.Weight,
		}
	})

	randomNum := weighted_random.WeightedRandomList(weightedRandList).Gen()
	// 存入簽到紀錄
	if err := s.repo.Mongo.Insert(s.db.Mongo.GetDB(), mongo_trainee3.CollectionName, mongo_trainee3.DS_MemberSignInRecord{
		ID:         primitive.NewObjectID(),
		MemberID:   value.MemberID,
		Rewards:    DS_Activity.Rewards[randomNum].Rewards,
		RecordTime: time.Now(),
		ActivityID: DS_Activity.ID,
	}); err != nil {
		logrus.Error("SignIn MongoDB Insert error:", err)
		return "", apicode.UnknownError
	}
	time.Sleep(5 * time.Second)
	return DS_Activity.Rewards[randomNum].Rewards, apicode.Success
}
