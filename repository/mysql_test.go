package repository

import (
	"database/sql"
	"errors"
	"testing"
	"trainee3/model/entity/mysql/mysql_trainee3"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fatih/structs"
	"github.com/stretchr/testify/assert"
	"github.com/tom-riddle-sr/struct_change_map"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMysql_Query_ErrRecordNotFound_Should_Return_Nil(t *testing.T) {
	expectDsActivity := mysql_trainee3.DSActivity{}
	actualDsActivity := mysql_trainee3.DSActivity{}

	db, mock, err := sqlmock.New() // 是一個假的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := convertToGormDB(db, mock) //用來轉換成gorm的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.
		ExpectQuery("SELECT (.+) FROM `DS_Activity` (.+)").
		WillReturnError(gorm.ErrRecordNotFound)

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.Query(gormDB, &actualDsActivity, "Open = ?", 1)

	assert.NoError(t, actualErr)                              // 預期是沒有錯誤
	assert.EqualValues(t, expectDsActivity, actualDsActivity) // 預期是一樣的
}
func TestMysql_Query_Error_Should_Return_Error(t *testing.T) {
	// expectDsActivity := mysql_trainee3.DSActivity{}
	actualDsActivity := mysql_trainee3.DSActivity{}

	db, mock, err := sqlmock.New() // 是一個假的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := convertToGormDB(db, mock) //用來轉換成gorm的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.
		ExpectQuery("SELECT (.+) FROM `DS_Activity` (.+)").
		WillReturnError(errors.New("any error"))

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.Query(gormDB, &actualDsActivity, "Open = ?", 1)

	assert.Error(t, actualErr)
	assert.EqualError(t, actualErr, "any error")
}
func TestMysql_Query_Success_Should_Return_Nil(t *testing.T) {
	expectDsActivity := mysql_trainee3.DSActivity{
		ID:   1,
		Name: "value2",
		Open: false,
	}
	actualDsActivity := mysql_trainee3.DSActivity{}

	db, mock, err := sqlmock.New() // 是一個假的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := convertToGormDB(db, mock) //用來轉換成gorm的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.
		ExpectQuery("SELECT (.+) FROM `DS_Activity` (.+)").
		WillReturnRows(sqlmock.NewRows([]string{"ID", "Name", "Open"}).AddRow(1, "value2", false))

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.Query(gormDB, &actualDsActivity, "Open = ?", 1)

	assert.NoError(t, actualErr)                              // 預期是沒有錯誤
	assert.EqualValues(t, expectDsActivity, actualDsActivity) // 預期是一樣的
}
func TestMysql_QueryAll_ErrRecordNotFound_Should_Return_Nil(t *testing.T) {
	expectDsActivity := mysql_trainee3.DSActivity{}
	actualDsActivity := mysql_trainee3.DSActivity{}
	rewardList := []mysql_trainee3.DSReward{}

	db, mock, err := sqlmock.New() // 是一個假的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := convertToGormDB(db, mock) //用來轉換成gorm的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.
		ExpectQuery("SELECT (.+) FROM `DS_Activity` (.+)").
		WillReturnError(gorm.ErrRecordNotFound)

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.QueryAll(gormDB, &actualDsActivity, &rewardList, "Open = ?", 1)

	assert.NoError(t, actualErr)                              // 預期是沒有錯誤
	assert.EqualValues(t, expectDsActivity, actualDsActivity) // 預期是一樣的
}
func TestMysql_QueryAll_Error_Should_Return_Error(t *testing.T) {
	actualDsActivity := mysql_trainee3.DSActivity{}
	rewardList := []mysql_trainee3.DSReward{}

	db, mock, err := sqlmock.New() // 是一個假的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := convertToGormDB(db, mock) //用來轉換成gorm的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.
		ExpectQuery("SELECT (.+) FROM `DS_Activity` (.+)").
		WillReturnError(errors.New("any error"))

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.QueryAll(gormDB, &actualDsActivity, &rewardList, "Open = ?", 1)

	assert.Error(t, actualErr)
	assert.EqualError(t, actualErr, "any error")
}
func TestMysql_QueryAll_Success_Should_Return_Nil(t *testing.T) {
	expectRewardList := []mysql_trainee3.DSReward{
		{
			ID:         1,
			ActivityID: 2,
			Weight:     3,
			Rewards:    "test rewards",
		},
	}
	actualRewardList := []mysql_trainee3.DSReward{}

	db, mock, err := sqlmock.New() // 是一個假的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := convertToGormDB(db, mock) //用來轉換成gorm的資料庫
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.
		ExpectQuery("SELECT (.+) FROM `DS_Rewards` (.+)").
		WillReturnRows(sqlmock.NewRows([]string{"Id", "ActivityId", "weight", "Rewards"}).
			AddRow(1, 2, 3, "test rewards"))

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.QueryAll(gormDB, &mysql_trainee3.DSReward{}, &actualRewardList, "Open = ?", 1)

	assert.NoError(t, actualErr)                              // 預期是沒有錯誤
	assert.EqualValues(t, expectRewardList, actualRewardList) // 預期是一樣的
}
func TestMysql_Update_Error_Should_Return_Error(t *testing.T) {
	actualDsActivity := mysql_trainee3.DSActivity{
		ID:   1,
		Name: "test",
		Open: false,
	}
	cols := structs.Map(&actualDsActivity)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := convertToGormDB(db, mock)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	expectErr := errors.New("any error")
	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `DS_Activity` (.+)$").
		WithArgs(1, "test", false, 1).
		WillReturnError(expectErr)
	mock.ExpectRollback()
	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.Update(gormDB, "id = ?", 1, actualDsActivity, cols)

	assert.Error(t, actualErr)
	assert.Equal(t, expectErr, actualErr)
}
func TestMysql_Update_Success_Should_Return_Nil(t *testing.T) {
	actualDsActivity := mysql_trainee3.DSActivity{
		ID:   1,
		Name: "test",
		Open: false,
	}
	cols := struct_change_map.New(&actualDsActivity)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := convertToGormDB(db, mock)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `DS_Activity` (.+)$").
		WithArgs(1, "test", false, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.Update(gormDB, "id = ?", 1, actualDsActivity, cols)

	assert.NoError(t, actualErr)
}
func TestMysql_UpdateColumns_Error_Should_Return_Error(t *testing.T) {
	actualDsActivity := mysql_trainee3.DSActivity{
		ID:   1,
		Name: "test",
		Open: false,
	}

	cols := map[string]interface{}{"Id": 1, "Name": "test", "Open": false}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := convertToGormDB(db, mock)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	expectErr := errors.New("any error")
	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `DS_Activity` (.+)$").
		WithArgs(1, "test", false, 1).
		WillReturnError(expectErr)
	mock.ExpectRollback()
	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.UpdateColumns(gormDB, &actualDsActivity, cols)

	assert.Error(t, actualErr)
	assert.Equal(t, expectErr, actualErr)
}
func TestMysql_UpdateColumns_Success_Should_Return_Nil(t *testing.T) {
	actualDsActivity := mysql_trainee3.DSActivity{
		ID:   1,
		Name: "test",
		Open: false,
	}

	cols := map[string]interface{}{"Id": 1, "Name": "test", "Open": false}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := convertToGormDB(db, mock)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `DS_Activity` (.+)$").
		WithArgs(1, "test", false, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.UpdateColumns(gormDB, &actualDsActivity, cols)

	assert.NoError(t, actualErr)
}
func TestMysql_Save_Error_Should_Return_Error(t *testing.T) {
	actualModel := mysql_trainee3.DSActivity{}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := convertToGormDB(db, mock)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	expectErr := errors.New("any error")

	// 期望开始一个事务
	mock.ExpectBegin()

	mock.ExpectExec("^INSERT INTO `DS_Activity` (.+)$").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(expectErr)

	mock.ExpectRollback()

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.Save(gormDB, &actualModel)

	assert.Error(t, actualErr)
	assert.EqualError(t, actualErr, expectErr.Error())
}
func TestMysql_Save_Success_Should_Return_Nil(t *testing.T) {
	actualModel := mysql_trainee3.DSActivity{}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := convertToGormDB(db, mock)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()

	mock.ExpectExec("^INSERT INTO `DS_Activity` (.+)$").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.Save(gormDB, &actualModel)

	assert.NoError(t, actualErr)
}

func convertToGormDB(db *sql.DB, mock sqlmock.Sqlmock) (*gorm.DB, error) {
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version()"}).AddRow("5.7.30"))
	return gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
}
