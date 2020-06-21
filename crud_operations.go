package main

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Operations interface {
		Create(model interface{}) (interface{}, error)
		Read() interface{}
		Index(urn string) (interface{}, error)
		Update(urn string) (interface{}, error)
		Delete(urn string) (bool, error)
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

func (config Configuration) Index(urn string) (interface{}, error) {
	var doc interface{}
	filter := bson.M{"urn": urn}
	err := config.Collection.FindOne(context.Background(), filter).Decode(&doc)
	return doc, err
}

func (config Configuration) Update(urn string) (interface{}, error) {
	filter := bson.M{"urn": urn}
	doc, err := config.Collection.UpdateOne(context.Background(), filter, "test")
	return doc, err
}

func (config Configuration) Delete(urn string) (bool, error) {
	filter := bson.M{"urn": urn}
	delRes, err := config.Collection.DeleteOne(context.Background(), filter)
	if err != nil || delRes.DeletedCount == 0 {
		err = errors.New("delete error: delete count less than 1")
		return false, err
	}
	return true, nil
}
