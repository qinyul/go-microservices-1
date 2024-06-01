package utils 

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"barqi.com/user/common"
	_ "barqi.com/user/docs"
	"time"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Utils struct {
}

func (u *Utils) GenerateJWT(username string, role string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":username,
		"iss":common.Config.Issuer,
		"aud":role,
		"exp":time.Now().Add(time.Hour).Unix(),
		"iat":time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(common.Config.JwtSecretPassword))
	if err != nil {
		log.Debugf("Failed to create token: %v",err)
	}else{	
		log.Debug("Token Created")
	}

	return tokenString, err
}

func (u *Utils) ValidateObjectID(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Debug("Error parsing object: ",err)
	}

	return objID, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}