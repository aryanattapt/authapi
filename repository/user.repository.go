package repository

import (
	"authapi/model"
	"authapi/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	mongoDBUserRepository = utils.MongoDBDatabase{DatabaseName: "goapi", CollectionName: "users"}
)

func GetUserByUsernameOrEmail(username string, email string) (data []map[string]interface{}, err error) {
	mongoDBUserRepository.Filter = bson.M{
		"isactive": true,
		"$or": bson.A{
			bson.M{"username": username},
			bson.M{"email": email},
		},
	}
	data, err = mongoDBUserRepository.GetMongoDB()
	return
}

func GetUserById(userid string) (data []map[string]interface{}, err error) {
	id, _ := primitive.ObjectIDFromHex(userid)
	mongoDBUserRepository.Filter = bson.M{
		"isactive": true,
		"_id":      id,
	}
	data, err = mongoDBUserRepository.GetMongoDB()
	return
}

func SaveUser(payload model.UserInsertPayload) (err error) {
	payload.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	payload.IsActive = true
	data, _ := utils.StructToMap(payload)
	mongoDBUserRepository.Payload = data
	err = mongoDBUserRepository.InsertMongoDB()
	return
}

func UpdateUser(payload model.UserUpdatePayload) (err error) {
	id, _ := primitive.ObjectIDFromHex(payload.ID)
	mongoDBUserRepository.Filter = bson.M{"_id": id}

	data, _ := utils.StructToMap(payload)
	mongoDBUserRepository.Payload = data
	err = mongoDBUserRepository.UpdateMongoDB()
	return
}

func GetAllUser(searchPayload map[string]interface{}) (data []map[string]interface{}, err error) {
	idPayload, ok := searchPayload["_id"].(string)
	if ok {
		id, _ := primitive.ObjectIDFromHex(idPayload)
		delete(searchPayload, "_id")
		mongoDBUserRepository.Filter = bson.D{{Key: "_id", Value: id}}
	} else {
		mongoDBUserRepository.Filter = searchPayload
	}
	data, err = mongoDBUserRepository.GetMongoDB()
	return
}
