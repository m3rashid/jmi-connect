package likes

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Topic struct {
	likeID    primitive.ObjectID `bson:"_id"`
	userID    primitive.ObjectID `bson:"userID"`
	postID    primitive.ObjectID `bson:"postID"`
	createdAt primitive.DateTime `bson:"createdAt"`
	updatedAt primitive.DateTime `bson:"updatedAt"`
}
