package mongo_trainee3

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionName = "DS_MemberSignInRecord"
)

type DS_MemberSignInRecord struct {
	ID         primitive.ObjectID `bson:"_id"`
	MemberID   int32              `bson:"memberId"`
	Rewards    string             `bson:"rewards"`
	RecordTime time.Time          `bson:"recordTime"`
	ActivityID int32              `bson:"activityId"`
}
