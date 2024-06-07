package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
    ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title        string             `json:"title,omitempty" bson:"title,omitempty"`
    Ingredients  []string           `json:"ingredients,omitempty" bson:"ingredients,omitempty"`
    Instructions string             `json:"instructions,omitempty" bson:"instructions,omitempty"`
    AuthorID     primitive.ObjectID `json:"author_id,omitempty" bson:"author_id,omitempty"`
    CreatedAt    int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}