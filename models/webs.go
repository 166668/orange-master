package models

import (
	"code.google.com/p/go-uuid/uuid"
	"time"
)

type Web struct {
	Uuid    string
	Url     string `bson:"url"`
	Pattern string `bson:"pattern"`
	Name    string `bson:"name`
}

type Web_Status struct {
	_id        string `bson:"_id"`
	Url        string `bson:"url"`
	Pattern    string `bson:"pattern"`
	Name       string `bson:"name`
	Web_status bool   `bson:"web_status"`
	Time       time.Time
}

func (this *Web) Gen_Web_UUID() {
	this.Uuid = uuid.New()
}
