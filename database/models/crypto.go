package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Crypto struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name,omitempty"`
	Up    int64              `json:"up" bson:"up,omitempty"`
	Down  int64              `json:"down" bson:"down,omitempty"`
	Total int64              `json:"total" bson:"total,omitempty"`
}
