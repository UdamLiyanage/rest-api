package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Operations interface {
		Create(model interface{}) interface{}
		Read() interface{}
		Index() interface{}
		Update() interface{}
		Delete() bool
	}

	Configuration struct {
		Collection *mongo.Collection
	}
)

func (config Configuration) Create(model interface{}) interface{} {
	doc, err := config.Collection.InsertOne(context.Background(), model)
	if err != nil {
		return err
	} else {
		return doc
	}
}

func (config Configuration) Read() interface{} {
	return nil
}

func (config Configuration) Index() interface{} {
	return nil
}

func (config Configuration) Update() interface{} {
	return nil
}

func (config Configuration) Delete() bool {
	return false
}
