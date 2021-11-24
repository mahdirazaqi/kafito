package otp

import (
	"context"
	"math/rand"
	"time"

	"github.com/kavenegar/kavenegar-go"
	"github.com/mahdirazaqi/kafito/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OTP struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Phone      string             `bson:"phone" json:"phone"`
	Code       string             `bson:"code" json:"code"`
	ExpireTime time.Time          `bson:"expire_time" json:"expire_time"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}

var (
	Kave     *kavenegar.Kavenegar
	Template string
)

func Init(token, template string) {
	Kave = kavenegar.New(token)
	Template = template
}

func (o *OTP) collection() *mongo.Collection {
	return database.Connection.Collection("otps")
}

func randomCode() (randomCode string) {
	rand.Seed(time.Now().UnixNano())
	digits := "1234567890"

	for i := 0; i < 6; i++ {
		randomCode += string(digits[rand.Intn(len(digits))])
	}

	return randomCode
}

func Generate(phone string) error {
	o := &OTP{
		ID:         primitive.NewObjectID(),
		Phone:      phone,
		Code:       randomCode(),
		ExpireTime: time.Now().Add(time.Minute * 2),
		CreatedAt:  time.Now(),
	}

	if _, err := o.collection().InsertOne(context.Background(), database.Bson(o)); err != nil {
		return err
	}

	if _, err := Kave.Verify.Lookup(o.Phone, Template, o.Code, nil); err != nil {
		return err
	}

	return nil
}

func Check(phone, code string) bool {
	opt := options.FindOne().SetSort([]bson.E{{Key: "created_at", Value: -1}})

	o := new(OTP)
	if err := o.collection().FindOne(context.Background(), bson.M{"phone": phone}, opt).Decode(o); err != nil {
		return false
	}

	if o.ExpireTime.UnixNano() < time.Now().UnixNano() {
		return false
	}

	return o.Code == code
}
