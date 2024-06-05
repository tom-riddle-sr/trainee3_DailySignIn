package repository

import (
	"database/sql"
	"errors"
	"testing"
	"trainee3/model/entity/mysql/mysql_trainee3"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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
	// expectDsActivity := mysql_trainee3.DSActivity{}
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
	rewardList := []mysql_trainee3.DSReward{}
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
	actualErr := mysqlRepo.QueryAll(gormDB, &actualDsActivity, &rewardList, "Open = ?", 1)

	assert.NoError(t, actualErr)                              // 預期是沒有錯誤
	assert.EqualValues(t, expectDsActivity, actualDsActivity) // 預期是一樣的
}
func TestMysql_Update_Error_Should_Return_Error(t *testing.T) {
	actualDsActivity := mysql_trainee3.DSActivity{}

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
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewErrorResult(expectErr))
	mock.ExpectRollback()

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.Update(gormDB, actualDsActivity)

	assert.Error(t, actualErr)

	assert.Contains(t, actualErr.Error(), expectErr.Error())
	assert.EqualError(t, actualErr, expectErr.Error())

}

// func TestMysql_QueryAll_Success_Should_Return_Nil(t *testing.T) {
// 	rewardList := []mysql_trainee3.DSReward{}
// 	expectDsActivity := mysql_trainee3.DSActivity{
// 		ID:   1,
// 		Name: "value2",
// 		Open: false,
// 	}
// 	actualDsActivity := mysql_trainee3.DSActivity{}

// 	db, mock, err := sqlmock.New() // 是一個假的資料庫
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	gormDB, err := convertToGormDB(db, mock) //用來轉換成gorm的資料庫
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	mock.
// 		ExpectQuery("SELECT (.+) FROM `DS_Activity` (.+)").
// 		WillReturnRows(sqlmock.NewRows([]string{"ID", "Name", "Open"}).AddRow(1, "value2", false))

// 	mysqlRepo := Mysql{}
// 	actualErr := mysqlRepo.QueryAll(gormDB, &actualDsActivity, &rewardList, "Open = ?", 1)

//		assert.NoError(t, actualErr)                              // 預期是沒有錯誤
//		assert.EqualValues(t, expectDsActivity, actualDsActivity) // 預期是一樣的
//	}

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

	// // 检查所有预期的操作是否都被调用
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }
}

func convertToGormDB(db *sql.DB, mock sqlmock.Sqlmock) (*gorm.DB, error) {
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version()"}).AddRow("5.7.30"))
	return gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
}
