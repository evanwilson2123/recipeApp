package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    RecipeID  primitive.ObjectID `json:"recipe_id,omitempty" bson:"recipe_id,omitempty"`
    AuthorID  primitive.ObjectID `json:"author_id,omitempty" bson:"author_id,omitempty"`
    Content   string             `json:"content,omitempty" bson:"content,omitempty"`
    CreatedAt int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
