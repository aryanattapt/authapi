package service

import (
	"authapi/model"
	"authapi/repository"
	"authapi/utils"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(payload *model.UserInsertPayload) error {
	var goValidator = validator.New()
	if err := goValidator.Struct(payload); err != nil {
		return err
	}

	data, err := repository.GetUserByUsernameOrEmail(payload.Username, payload.Username)
	if err != nil {
		return err
	}

	if len(data) != 0 {
		return errors.New("user is exist")
	}

	payload.Password = utils.HashPasswordBCrypt(payload.Password)
	payload.IsActive = true
	if err := repository.SaveUser(*payload); err != nil {
		return err
	}

	return nil
}

func Signin(payload *model.SignInPayload) (error, interface{}) {
	data, err := repository.GetUserByUsernameOrEmail(payload.Username, payload.Username)
	if err != nil {
		return err, nil
	}

	if len(data) == 0 {
		return errors.New("user not exist"), nil
	}

	if !utils.ComparePasswordBCrypt(payload.Password, data[0]["password"].(string)) {
		return errors.New("invalid password"), nil
	}

	var jwtPayload = JWTPayload{
		StandardClaims: jwt.StandardClaims{
			Audience: utils.EncodeBase64(data[0]["_id"].(primitive.ObjectID).Hex()),
			Id:       utils.HashSHA1(utils.GenerateUUID()),
			Issuer:   "AUTH",
			Subject:  "AUTH",
		},
		SessionID: utils.HashMD5(utils.GenerateRandomNumber(32)),
	}

	jwtToken, err, _ := jwtPayload.CreateJWTToken()
	if err != nil {
		return err, nil
	}

	return nil, jwtToken
}
