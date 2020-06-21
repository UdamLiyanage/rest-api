package main

import "go.mongodb.org/mongo-driver/mongo"

type (
	Operations interface {
		Create() interface{}
		Read() interface{}
		Index() interface{}
		Update() interface{}
		Delete() bool
	}

	Configuration struct {
		Collection *mongo.Collection
	}
)

func (config Configuration) Create() interface{} {
	return nil
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
