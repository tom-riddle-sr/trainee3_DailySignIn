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
		WillReturnError(errors.New("any error"))

	mysqlRepo := Mysql{}
	actualErr := mysqlRepo.Query(gormDB, &actualDsActivity, "Open = ?", 1)

	assert.Equal(t, expectErr, actualErr)                     // 預期是沒有錯誤
	assert.EqualValues(t, expectDsActivity, actualDsActivity) // 預期是一樣的
}

func convertToGormDB(db *sql.DB, mock sqlmock.Sqlmock) (*gorm.DB, error) {
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version()"}).AddRow("5.7.30"))
	return gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
}
