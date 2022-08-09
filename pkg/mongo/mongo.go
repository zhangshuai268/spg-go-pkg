package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Engine struct {
	Orm *mongo.Database
}

// InitMongoEngin 初始化mongodb引擎
//  appUri: mongodb://User:PassWord@127.0.0.1:27017
//  dbName: 数据库名称
func InitMongoEngin(appUri string, dbName string) (*Engine, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(appUri))
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	return &Engine{Orm: db}, nil
}
