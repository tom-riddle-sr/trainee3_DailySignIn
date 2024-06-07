package services

import (
	"encoding/json"
	"trainee3/database"
	"trainee3/database/mysql"
	apicode "trainee3/lib/apiCode"
	"trainee3/model/entity/mysql/mysql_trainee3"
	model_redis "trainee3/model/entity/redis"
	"trainee3/repository"

	"github.com/sirupsen/logrus"
)

type ICache interface {
	ServicesRefresh() apicode.Code
}

type cache struct {
	repo *repository.Repo
	db   *database.Database
}

func NewCache(repo *repository.Repo, db *database.Database) ICache {
	return &cache{
		repo: repo,
		db:   db,
	}
}

func (s *cache) ServicesRefresh() apicode.Code {
	var (
		activityModel mysql_trainee3.DSActivity
		rewardModel   mysql_trainee3.DSReward
		rewardList    []mysql_trainee3.DSReward
	)
	db := s.db.Mysql.GetDB(mysql.Trainee3)
	if err := s.repo.Mysql.Query(db, &activityModel, "Open = ?", 1); err != nil {
		return apicode.UnknownError
	}
	// 沒有開啟的活動
	if activityModel.ID == 0 {
		return apicode.Success
	}

	if err := s.repo.Mysql.QueryAll(db, &rewardModel, &rewardList, "ActivityId = ?", activityModel.ID); err != nil {
		return apicode.UnknownError
	}

	if len(rewardList) == 0 {
		return apicode.UnknownError
	}

	json, _ := json.Marshal(model_redis.DSInRedis{
		ID:      activityModel.ID,
		Name:    activityModel.Name,
		Open:    activityModel.Open,
		Rewards: rewardList,
	})

	if err := s.repo.Redis.Set(s.db.Redis, model_redis.DSActivityTableName, json, 0); err != nil {
		logrus.Error("redis.Set error:", err)
		return apicode.UnknownError
	}
	return apicode.Success
}
