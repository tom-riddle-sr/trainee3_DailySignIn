package repository

import (
	"errors"
	"testing"
	"time"

	model "trainee3/model/entity/redis"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedis_Set_Error_Should_Return_Error(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "test"
	value := "sdfsds"
	expiration := int32(2)
	expectErr := errors.New("my custom error")
	mock.ExpectSet(key, value, time.Duration(expiration)*time.Second).SetErr(expectErr)

	repository := Redis{}
	err := repository.Set(db, key, value, expiration)
	assert.Error(t, err)
}

func TestRedis_Set_Success_Should_Return_Nil(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "test"
	value := "value"
	expiration := int32(2)
	expectVal := "success"

	mock.ExpectSet(key, value, time.Duration(expiration)*time.Second).SetVal(expectVal)

	repository := Redis{}
	err := repository.Set(db, key, value, expiration)
	assert.NoError(t, err)

}
func TestRedis_Get_Error_Should_Return_Error(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "test"
	expectErr := errors.New("my custom error")
	mock.ExpectGet(key).SetErr(expectErr)
	var DS_Activity model.DSInRedis
	repository := Redis{}
	err := repository.Get(db, key, &DS_Activity)
	assert.Error(t, err)

}
func TestRedis_Get_JSON_UnMarshal_Err_Should_Return_Error(t *testing.T) {
	db, mock := redismock.NewClientMock()
	expectKey := "test"
	expectValue := "let umarshal fail"

	mock.ExpectGet(expectKey).SetVal(expectValue)

	var dsActivity model.DSInRedis
	repository := Redis{}

	err := repository.Get(db, expectKey, &dsActivity)
	assert.Error(t, err)
}

func TestRedis_Get_JSON_UnMarshal_Success_Should_Return_Nil(t *testing.T) {
	db, mock := redismock.NewClientMock()
	expectKey := "test"
	expectValue := `{"id":1,"name":"test"}`

	mock.ExpectGet(expectKey).SetVal(expectValue)

	var dsActivity model.DSInRedis
	repository := Redis{}

	err := repository.Get(db, expectKey, &dsActivity)
	assert.NoError(t, err)
}
