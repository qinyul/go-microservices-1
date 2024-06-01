package database


import (
	"barqi.com/user/common"
	"barqi.com/user/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	log "github.com/sirupsen/logrus"
	"fmt"
	"context"
)
type MongoDB struct {
	Databasename string
	ClusterName string
	Client *mongo.Client
}

func (db *MongoDB) Init() error {
	var err error
	db.Databasename = common.Config.MgDbName
	db.ClusterName = common.Config.ClusterName
	username := common.Config.MgDbUsername
	password := common.Config.MgDbPassword

	
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s.as86ndc.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0",username,password,db.ClusterName) 

	db.Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	var result bson.M 
	if err := db.Client.Database(db.Databasename).RunCommand(context.TODO(),bson.D{{"ping",1}}).Decode(&result); err != nil {
		log.Debug("Cannot connect to DB")
		panic(err)
	}
	
	log.Debug("Successfully connected to DB")

	return db.initData()
}

func (db *MongoDB) initData() error {
	coll := db.Client.Database(db.Databasename).Collection(common.ColUsers)

	filter := bson.D{{"username","admin"}}
	count, err := coll.CountDocuments(context.TODO(),filter)

	if err != nil {
		panic(err)
	}
	
	log.Debugf("Collection Counted: %d",count)

	if count < 1 {
		adminPassword, err := utils.HashPassword("admin") 
		if err != nil {
			panic(err)
		}
		var user User
		user = User{"admin", adminPassword}
		log.Debug("Creating admin user")
		indexModel := mongo.IndexModel{
			Keys:bson.D{{Key:"username",Value:1}},
			Options: options.Index().SetUnique(true),
		}
		_, err = coll.Indexes().CreateOne(context.TODO(),indexModel)
		result, err := coll.InsertOne(context.TODO(), user)
		if err != nil {
			panic(err)
		}
		log.Debug("Admin Created %v",result)
	}
	return err
}

func (db *MongoDB) Close() {
	if db.Client != nil {
		log.Debug("Closing MongoDB Client")
		db.Client.Disconnect(context.TODO())
	}
}