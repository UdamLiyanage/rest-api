package main

import "go.mongodb.org/mongo-driver/mongo"

type (
	Operation interface {
		create() interface{}
		read() interface{}
		update() interface{}
		delete() bool
	}

	Configuration struct {
		Collection *mongo.Collection
	}
)

func (config Configuration) create() interface{} {
	return nil
}

func (config Configuration) read() interface{} {
	return nil
}

func (config Configuration) update() interface{} {
	return nil
}

func (config Configuration) delete() bool {
	return false
}
