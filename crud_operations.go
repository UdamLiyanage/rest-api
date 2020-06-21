package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Operations interface {
		Create(model interface{}) (interface{}, error)
		Read() interface{}
		Index() interface{}
		Update() interface{}
		Delete() bool
	}

	Configuration struct {
		Collection *mongo.Collection
	}
)

func (config Configuration) Create(model interface{}) (interface{}, error) {
	doc, err := config.Collection.InsertOne(context.Background(), model)
	if err != nil {
		return nil, err
	} else {
		return doc, nil
	}
}

func (config Configuration) Read() interface{} {
	return nil
}

func (config Configuration) Index() interface{} {
	return nil
}

func (config Configuration) Update() interface{} {
	filter := bson.M{"urn": "test"}
	doc, err := config.Collection.UpdateOne(context.Background(), filter, "test")
	if err != nil {
		panic(err)
	} else {
		return doc
	}
}

func (config Configuration) Delete() bool {
	return false
}
