package user

import (
	"strings"

	"github.com/ttacon/libphonenumber"
	"go.mongodb.org/mongo-driver/bson"
)

func ParsePhoneNumber(phone string, region string) (string, error) {
	number, err := libphonenumber.Parse(phone, strings.ToUpper(region))
	if err != nil {
		return "", err
	}

	new := libphonenumber.Format(number, libphonenumber.INTERNATIONAL)

	return strings.ReplaceAll(new, " ", ""), nil
}

func Auth(phone string) (*User, error) {
	u, err := FindOne(bson.M{"phone": phone})
	if err != nil {
		u = &User{
			Name:  phone,
			Phone: phone,
		}

		if err := u.Save(); err != nil {
			return nil, err
		}
	}

	return u, nil
}
