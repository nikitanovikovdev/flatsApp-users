package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

)

type MongoConfig struct {
	Host     string
	Port     string
}

func NewMongoDB(c *MongoConfig) (*mongo.Client, error) {
	//uri := "mongodb://localhost:27017"
	uri := fmt.Sprintf("mongodb://%v:%v", c.Host, c.Port)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}