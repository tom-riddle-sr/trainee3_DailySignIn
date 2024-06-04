package services

import (
	"encoding/json"
	"errors"
	"testing"
	"trainee3/database"
	"trainee3/database/mysql"
	apicode "trainee3/lib/apiCode"
	"trainee3/mocks"
	"trainee3/model/entity/mysql/mysql_trainee3"
	model_redis "trainee3/model/entity/redis"
	"trainee3/repository"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_cache_ServicesRefresh_Query_Error_Should_Unknown_Error(t *testing.T) {
	// given
	db := &gorm.DB{}
	dbName := mysql.Trainee3
	activityModel := mysql_trainee3.DSActivity{}

	// expect
	expect := apicode.UnknownError
	// mock
	mockMysqlDB := mocks.NewIMysqlDB(t)
	mockMysqlDB.EXPECT().GetDB(dbName).Return(db)

	mockDB := &database.Database{
		Mysql: mockMysqlDB,
		Mongo: nil,
		Redis: nil,
	}

	mockMysqlRepo := mocks.NewIMysql(t)
	mockMysqlRepo.EXPECT().Query(db, &activityModel, "Open = ?", 1).Return(errors.New("any error"))

	mockRepo := &repository.Repo{
		Mysql: mockMysqlRepo,
		Mongo: nil,
		Redis: nil,
	}

	cache := &cache{
		repo: mockRepo,
		db:   mockDB,
	}

	actual := cache.ServicesRefresh()
	assert.Equal(t, actual, expect)
}
func Test_cache_ServicesRefresh_ActivityModel_ID_Is_Zero_Should_Success(t *testing.T) {
	// given
	db := &gorm.DB{}
	dbName := mysql.Trainee3
	activityModel := mysql_trainee3.DSActivity{}
	// expect
	expect := apicode.Success

	// mock
	mockMysqlDB := mocks.NewIMysqlDB(t)
	mockMysqlDB.EXPECT().GetDB(dbName).Return(db)
	mockDB := &database.Database{
		Mysql: mockMysqlDB,
		Mongo: nil,
		Redis: nil,
	}

	mockMysqlRepo := mocks.NewIMysql(t)
	mockMysqlRepo.EXPECT().Query(db, &activityModel, "Open = ?", 1).Return(nil)

	mockRepo := &repository.Repo{
		Mysql: mockMysqlRepo,
		Mongo: nil,
		Redis: nil,
	}

	cache := &cache{
		repo: mockRepo,
		db:   mockDB,
	}

	actual := cache.ServicesRefresh()

	assert.Equal(t, actual, expect)
}

func Test_cache_ServicesRefresh_QueryAll_Error_Should_Unknown_Error(t *testing.T) {
	// given
	db := &gorm.DB{}
	dbName := mysql.Trainee3
	activityModel := mysql_trainee3.DSActivity{}
	rewardModel := mysql_trainee3.DSReward{}
	var rewardList []mysql_trainee3.DSReward
	activityModelID := int32(1)

	// expect
	expect := apicode.UnknownError

	// mock
	mockMysqlDB := mocks.NewIMysqlDB(t)
	mockMysqlDB.EXPECT().GetDB(dbName).Return(db)
	mockDB := &database.Database{
		Mysql: mockMysqlDB,
		Mongo: nil,
		Redis: nil,
	}

	mockMysqlRepo := mocks.NewIMysql(t)
	mockMysqlRepo.EXPECT().Query(db, &activityModel, "Open = ?", 1).RunAndReturn(
		func(db *gorm.DB, model interface{}, condition string, values ...interface{}) error {
			// 型別轉換
			dsActivity := (model).(*mysql_trainee3.DSActivity)
			dsActivity.ID = 1
			return nil
		})

	mockMysqlRepo.EXPECT().QueryAll(db, &rewardModel, &rewardList, "ActivityId = ?", activityModelID).Return(errors.New("any error"))

	mockRepo := &repository.Repo{
		Mysql: mockMysqlRepo,
		Mongo: nil,
		Redis: nil,
	}

	cache := &cache{
		repo: mockRepo,
		db:   mockDB,
	}

	actual := cache.ServicesRefresh()

	assert.Equal(t, actual, expect)
}

func Test_cache_ServicesRefresh_RewardList_Length_Is_Zero_Should_Unknown_Error(t *testing.T) {
	// given
	db := &gorm.DB{}
	dbName := mysql.Trainee3
	activityModel := mysql_trainee3.DSActivity{}
	rewardModel := mysql_trainee3.DSReward{}
	var rewardList []mysql_trainee3.DSReward
	activityModelID := int32(1)

	// expect
	expect := apicode.UnknownError

	// mock
	mockMysqlDB := mocks.NewIMysqlDB(t)
	mockMysqlDB.EXPECT().GetDB(dbName).Return(db)
	mockDB := &database.Database{
		Mysql: mockMysqlDB,
		Mongo: nil,
		Redis: nil,
	}

	mockMysqlRepo := mocks.NewIMysql(t)
	mockMysqlRepo.EXPECT().Query(db, &activityModel, "Open = ?", 1).RunAndReturn(
		func(db *gorm.DB, model interface{}, condition string, values ...interface{}) error {
			// 型別轉換
			dsActivity := (model).(*mysql_trainee3.DSActivity)
			dsActivity.ID = 1
			return nil
		})

	mockMysqlRepo.EXPECT().QueryAll(db, &rewardModel, &rewardList, "ActivityId = ?", activityModelID).
		RunAndReturn(func(db *gorm.DB, model interface{}, result *[]mysql_trainee3.DSReward, condition string, values ...interface{}) error {
			return nil
		})

	mockRepo := &repository.Repo{
		Mysql: mockMysqlRepo,
		Mongo: nil,
		Redis: nil,
	}

	cache := &cache{
		repo: mockRepo,
		db:   mockDB,
	}

	actual := cache.ServicesRefresh()

	assert.Equal(t, actual, expect)
}

func Test_cache_ServicesRefresh_Redis_Set_Should_Unknown_Error(t *testing.T) {
	// given
	gormDB := &gorm.DB{}
	redisClient := &redis.Client{}
	dbName := mysql.Trainee3
	activityModel := mysql_trainee3.DSActivity{}
	rewardModel := mysql_trainee3.DSReward{}
	var rewardList []mysql_trainee3.DSReward

	// expect
	expect := apicode.UnknownError
	expectDSRewardList := []mysql_trainee3.DSReward{
		{
			ID: 1,
		},
	}
	expectDSActivity := mysql_trainee3.DSActivity{
		ID:   1,
		Name: "test",
		Open: false,
	}

	expectRedisDataJson, _ := json.Marshal(model_redis.DSInRedis{
		ID:      expectDSActivity.ID,
		Name:    expectDSActivity.Name,
		Open:    expectDSActivity.Open,
		Rewards: expectDSRewardList,
	})

	// mock
	mockMysqlDB := mocks.NewIMysqlDB(t)
	mockMysqlDB.EXPECT().GetDB(dbName).Return(gormDB)
	mockDB := &database.Database{
		Mysql: mockMysqlDB,
		Mongo: nil,
		Redis: redisClient,
	}

	mockMysqlRepo := mocks.NewIMysql(t)
	mockMysqlRepo.EXPECT().Query(gormDB, &activityModel, "Open = ?", 1).RunAndReturn(
		func(db *gorm.DB, model interface{}, condition string, values ...interface{}) error {
			// 型別轉換
			dsActivity := (model).(*mysql_trainee3.DSActivity)
			dsActivity.ID = expectDSActivity.ID
			dsActivity.Name = expectDSActivity.Name
			dsActivity.Open = expectDSActivity.Open
			return nil
		})

	mockMysqlRepo.EXPECT().QueryAll(gormDB, &rewardModel, &rewardList, "ActivityId = ?", expectDSActivity.ID).
		RunAndReturn(func(db *gorm.DB, model interface{}, result *[]mysql_trainee3.DSReward, condition string, values ...interface{}) error {
			*result = expectDSRewardList
			return nil
		})

	mockRedisRepo := mocks.NewIRedis(t)
	mockRedisRepo.EXPECT().Set(redisClient, model_redis.DSActivityTableName, expectRedisDataJson, int32(0)).Return(errors.New("any error"))

	mockRepo := &repository.Repo{
		Mysql: mockMysqlRepo,
		Mongo: nil,
		Redis: mockRedisRepo,
	}

	cache := &cache{
		repo: mockRepo,
		db:   mockDB,
	}

	actual := cache.ServicesRefresh()

	assert.Equal(t, actual, expect)
}

func Test_cache_ServicesRefresh_Success(t *testing.T) {
	// given
	gormDB := &gorm.DB{}
	redisClient := &redis.Client{}
	dbName := mysql.Trainee3
	activityModel := mysql_trainee3.DSActivity{}
	rewardModel := mysql_trainee3.DSReward{}
	var rewardList []mysql_trainee3.DSReward

	// expect
	expect := apicode.Success
	expectDSRewardList := []mysql_trainee3.DSReward{
		{
			ID: 1,
		},
	}

	expectDSActivity := mysql_trainee3.DSActivity{
		ID:   1,
		Name: "test",
		Open: false,
	}

	expectRedisDataJson, _ := json.Marshal(model_redis.DSInRedis{
		ID:      expectDSActivity.ID,
		Name:    expectDSActivity.Name,
		Open:    expectDSActivity.Open,
		Rewards: expectDSRewardList,
	})

	// mock
	mockMysqlDB := mocks.NewIMysqlDB(t)
	mockMysqlDB.EXPECT().GetDB(dbName).Return(gormDB)
	mockDB := &database.Database{
		Mysql: mockMysqlDB,
		Mongo: nil,
		Redis: redisClient,
	}

	mockMysqlRepo := mocks.NewIMysql(t)
	mockMysqlRepo.EXPECT().Query(gormDB, &activityModel, "Open = ?", 1).RunAndReturn(
		func(db *gorm.DB, model interface{}, condition string, values ...interface{}) error {
			// 型別轉換
			dsActivity := (model).(*mysql_trainee3.DSActivity)
			dsActivity.ID = expectDSActivity.ID
			dsActivity.Name = expectDSActivity.Name
			dsActivity.Open = expectDSActivity.Open
			return nil
		})

	mockMysqlRepo.EXPECT().QueryAll(gormDB, &rewardModel, &rewardList, "ActivityId = ?", expectDSActivity.ID).
		RunAndReturn(func(db *gorm.DB, model interface{}, result *[]mysql_trainee3.DSReward, condition string, values ...interface{}) error {
			*result = expectDSRewardList
			return nil
		})

	mockRedisRepo := mocks.NewIRedis(t)
	mockRedisRepo.EXPECT().
		Set(redisClient, model_redis.DSActivityTableName, expectRedisDataJson, int32(0)).
		Return(nil)

	mockRepo := &repository.Repo{
		Mysql: mockMysqlRepo,
		Mongo: nil,
		Redis: mockRedisRepo,
	}

	cache := &cache{
		repo: mockRepo,
		db:   mockDB,
	}

	actual := cache.ServicesRefresh()

	assert.Equal(t, actual, expect)
}
