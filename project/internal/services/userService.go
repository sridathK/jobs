package services

import (
	"errors"

	"project/internal/model"
	"project/internal/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	r repository.Users
	c repository.Company
}

func NewService(r repository.Users, c repository.Company) (Service, error) {
	if r == nil {
		return Service{}, errors.New("db connection not given")
	}

	return Service{r: r, c: c}, nil

}

func (s Service) UserSignup(nu model.UserSignup) (model.User, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg("error occured in hashing password")
		return model.User{}, errors.New("hashing password failed")
	}

	user := model.User{UserName: nu.UserName, Email: nu.Email, PasswordHash: string(hashedPass)}
	// database.CreateTable()
	cu, err := s.r.CreateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.User{}, errors.New("user creation failed")
	}

	return cu, nil

}
func (s Service) Userlogin(l model.UserLogin) (jwt.RegisteredClaims, error) {
	fu, err := s.r.FetchUserByEmail(l.Email)
	if err != nil {
		log.Error().Err(err).Msg("couldnot find user")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(fu.PasswordHash), []byte(l.Password))
	if err != nil {
		log.Error().Err(err).Msg("password of user incorrect")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}
	c := jwt.RegisteredClaims{
		Issuer:    "service project",
		Subject:   strconv.FormatUint(uint64(fu.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	return c, nil

}
