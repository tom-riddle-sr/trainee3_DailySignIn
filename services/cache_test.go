package services

import (
	"errors"
	"testing"
	"trainee3/database"
	"trainee3/database/mysql"
	apicode "trainee3/lib/apiCode"
	"trainee3/mocks"
	"trainee3/model/entity/mysql/mysql_trainee3"
	"trainee3/repository"

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
	//EXPECT() 代表預期這個方法會被呼叫
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
	mockMysqlRepo.EXPECT().Query(db, &activityModel, "Open = ?", 1).RunAndReturn(
		func(db *gorm.DB, model interface{}, condition string, values ...interface{}) error {
			dsActivity := (model).(*mysql_trainee3.DSActivity)
			dsActivity.ID = 1
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
