package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	Title           string             `bson:"title,omitempty"`
	Rating          int                `bson:"rating,omitempty"`
	TotalCollection string             `bson:"totalCollection,omitempty"`
	Starring        []string           `bson:"starring,omitempty"`
	ReleasedDate    string             `bson:"releasedDate,omitempty"`
	Categories      []string           `bson:"categories,omitempty"`
	Img             string             `bson:"img,omitempty"`
}
