package auth

import (
	"barqi.com/user/database"
	"barqi.com/user/utils"
	"barqi.com/user/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	log "github.com/sirupsen/logrus"
	"context"
	"errors"
	"fmt"
)

type User struct {
	utils *utils.Utils
}

func (u *User) Login(username string, password string) (database.User, error){
	dbClient := database.Database.Client

	coll := dbClient.Database(database.Database.Databasename).Collection(common.ColUsers)

	var user database.User
	var err error

	filter := bson.D{
		{"username",username},
	}

	err = coll.FindOne(context.TODO(),filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Debug("user not exist")
		}else {
			log.Debugf("Someting went wrong: %v",err)
		}
	}else{
		log.Debug("User found")
	}

	isPasswordVerified := utils.VerifyPassword(password, user.Password)

	if !isPasswordVerified{
		log.Debug("Password / Username not match")
		err = errors.New(common.WrongUsernameOrPassword)
	}
	return user, err
}

func (u *User) Insert(username string, password string) error {
	dbClient := database.Database.Client
	coll := dbClient.Database(database.Database.Databasename).Collection(common.ColUsers)
	var user database.User
	hashedPassword,err := utils.HashPassword(password)
	user = database.User{username,hashedPassword}
	_, err = coll.InsertOne(context.TODO(),user)

	if err == nil {
		log.Debug("New user created")
	}

	return err
}

func (u *User) GetAll(username string) ([]database.UserInformation, error){
	dbClient := database.Database.Client
	coll :=  dbClient.Database(database.Database.Databasename).Collection(common.ColUsers)
	findOptions := options.Find()
	findOptions.SetBatchSize(10)

	filter := bson.D{}
	if len(username) > 0 {
		usernameRegex := fmt.Sprintf("%s.*",username)
		filter = bson.D{{"username", bson.D{{"$regex",usernameRegex}}}}
	}

	cursor,err := coll.Find(context.TODO(),filter,findOptions)
	defer cursor.Close(context.TODO())

	var users []database.UserInformation
	for cursor.Next(context.TODO()){
		var user database.UserInformation
		if err := cursor.Decode(&user); err != nil {
			log.Debug("User cursor [ERROR]: ",err.Error())
		}
		users = append(users,user)
	}

	if err := cursor.Err(); err != nil {
		log.Debug(err)
	}

	if len(users) == 0 {
		log.Debug("Not returning any users")
		users = []database.UserInformation{}
	}

	return users,err
}

func (u *User) GetByID(id string) (database.UserInformation, error) {
	objID, err := u.utils.ValidateObjectID(id)
	if err != nil {
		log.Debug("ID is not valid: ",err.Error())
		return database.UserInformation{}, err
	}

	filter := bson.D{{"_id",objID}}
	var user database.UserInformation
	dbClient := database.Database.Client
	coll :=  dbClient.Database(database.Database.Databasename).Collection(common.ColUsers)

	err =  coll.FindOne(context.TODO(),filter).Decode(&user)
	if err != nil{
		log.Debug("GetByID [ERROR]: ", err.Error())
	}

	return user, err
}

func (u *User) DeleteByID(id string) error {
	objID, err := u.utils.ValidateObjectID(id)

	filter := bson.D{{"_id",objID}}
	dbClient := database.Database.Client
	coll :=  dbClient.Database(database.Database.Databasename).Collection(common.ColUsers)
	result, err := coll.DeleteOne(context.TODO(),filter)

	if err != nil {
		log.Debug("Delete process failed ",err.Error())
		return err
	}

	if result.DeletedCount == 1 {
		log.Debug("Document deleted successfully")
	}else {
		err = errors.New(common.UserNotExist)
		log.Debug("Document not found or not deleted")
	}

	if err != nil {
		log.Debug("GetByID [ERROR]: ", err.Error())
	}

	return err
}

func (u *User) Update(user database.User, id string) error {
	objID, err := u.utils.ValidateObjectID(id)

	filter := bson.D{{"_id",objID}}
	dbClient := database.Database.Client
	coll :=  dbClient.Database(database.Database.Databasename).Collection(common.ColUsers)

	hashedPassword,err := utils.HashPassword(user.Password)
	update := bson.D{
		{"$set",bson.D{
			{"username",user.Username},
			{"password",hashedPassword},
		}},
	}

	result, err := coll.UpdateOne(context.TODO(),filter,update)

	if result.ModifiedCount == 1 {
		log.Debug("Document deleted successfully")
	}else {
		err = errors.New(common.UserNotExist)
		log.Debug("Document not found or not deleted")
	}
	
	return err
}