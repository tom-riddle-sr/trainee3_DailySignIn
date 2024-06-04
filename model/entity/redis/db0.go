package redis

import "trainee3/model/entity/mysql/mysql_trainee3"

const (
	DSActivityTableName string = "DS_Activity"
)

type DSInRedis struct {
	ID      int32  `json:"Id"`
	Name    string `json:"Name"`
	Open    bool   `json:"Open"`
	Rewards []mysql_trainee3.DSReward
}
