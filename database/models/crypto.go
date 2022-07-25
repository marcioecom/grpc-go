package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Crypto struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Up    int64              `bson:"up,omitempty"`
	Down  int64              `bson:"down,omitempty"`
	Total int64              `bson:"total,omitempty"`
}
