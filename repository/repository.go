package repository

type Repo struct {
	Mysql IMysql
	Mongo IMongo
	Redis IRedis
}

func NewRepo(mysql IMysql, mongo IMongo, redis IRedis) *Repo {
	return &Repo{
		Mysql: mysql,
		Mongo: mongo,
		Redis: redis,
	}
}
