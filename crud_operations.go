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
		Update(urn string) (interface{}, error)
		Delete() bool
	}

	Configuration struct {
		Collection *mongo.Collection
	}
)

func (config Configuration) Create(model interface{}) (interface{}, error) {
	doc, err := config.Collection.InsertOne(context.Background(), model)
	return doc, err
}

func (config Configuration) Read() interface{} {
	return nil
}

func (config Configuration) Index() interface{} {
	return nil
}

func (config Configuration) Update(urn string) (interface{}, error) {
	filter := bson.M{"urn": urn}
	doc, err := config.Collection.UpdateOne(context.Background(), filter, "test")
	return doc, err
}

func (config Configuration) Delete() bool {
	return false
}
