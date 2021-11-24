package user

import (
	"context"
	"time"

	"github.com/mahdirazaqi/kafito/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Bio       string             `bson:"bio" json:"bio"`
	Phone     string             `bson:"phone" json:"phone"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func (*User) collection() *mongo.Collection {
	return database.Connection.Collection("users")
}

func Count(filter bson.M) int {
	u := new(User)
	count, _ := u.collection().CountDocuments(context.Background(), filter)
	return int(count)
}

func Find(filter bson.M, page, limit int, sorts ...bson.E) []User {
	opt := options.Find()
	opt.SetSort(sorts)

	if limit > 0 {
		opt.SetLimit(int64(limit))
	}

	if page > 1 {
		opt.SetSkip(int64((page - 1) * limit))
	}

	u := new(User)
	cursor, err := u.collection().Find(context.Background(), filter, opt)
	if err != nil {
		return nil
	}

	users := []User{}
	for cursor.Next(context.Background()) {
		u := new(User)
		if err := cursor.Decode(u); err != nil {
			continue
		}

		users = append(users, *u)
	}

	return users
}

func FindOne(filter bson.M) (*User, error) {
	u := new(User)
	if err := u.collection().FindOne(context.Background(), filter).Decode(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Insert() error {
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()

	_, err := u.collection().InsertOne(context.Background(), database.Bson(u))
	return err
}

func (u *User) Update() error {
	_, err := u.collection().UpdateOne(context.Background(), bson.M{"_id": u.ID}, bson.M{"$set": database.Bson(u)})
	return err
}

func (u *User) Save() error {
	if u.ID.IsZero() {
		return u.Insert()
	}

	return u.Update()
}

func (u *User) Delete() error {
	_, err := u.collection().DeleteOne(context.Background(), bson.M{"_id": u.ID})
	return err
}
