package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"

	"project/internal/model"
	redisconn "project/internal/redisConn"

	//redisconn "project/internal/redisConn"
	"project/internal/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// type Service struct {
// 	r  repository.Users
// 	c  repository.Company
// 	re redisconn.Caching
// }

// GetCompanyById implements CompanyService.

// func NewService(r repository.Users, c repository.Company, re redisconn.Caching) (*Service, error) {
// 	if r == nil {
// 		return nil, errors.New("db connection not given")
// 	}

// 	return &Service{r: r, c: c, re: re}, nil

// }

type UserServiceImp struct {
	r  repository.Users
	re redisconn.Caching
}

func NewUserServiceImp(r repository.Users, re redisconn.Caching) (UsersService, error) {
	if r == nil {
		return nil, errors.New("db connection not given")
	}

	return &UserServiceImp{r: r, re: re}, nil

}

//go:generate mockgen -source=userService.go -destination=userservice_mock.go -package=services
type UsersService interface {
	UserSignup(nu model.UserSignup) (model.User, error)
	Userlogin(l model.UserLogin) (jwt.RegisteredClaims, error)
	UserForgetPassword(uf model.UserForgetPassword) (string, error)
	UserUpdatePassword(up model.UserUpdatePassword) (string, error)
}

func (s *UserServiceImp) UserSignup(nu model.UserSignup) (model.User, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg("error occured in hashing password")
		return model.User{}, errors.New("hashing password failed")
	}

	user := model.User{UserName: nu.UserName, Email: nu.Email, PasswordHash: string(hashedPass), Dob: nu.Dob}
	// database.CreateTable()
	cu, err := s.r.CreateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.User{}, errors.New("user creation failed")
	}

	return cu, nil

}
func (s *UserServiceImp) Userlogin(l model.UserLogin) (jwt.RegisteredClaims, error) {
	fu, err := s.r.FetchUserByEmail(l.Email)
	if err != nil {
		log.Error().Err(err).Msg("couldnot find user")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}
	fmt.Println(fu)
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
	fmt.Println(c)

	return c, nil

}

func (s *UserServiceImp) UserForgetPassword(uf model.UserForgetPassword) (string, error) {
	email := uf.Email
	dob := uf.DOB
	_, err := s.r.FetchUserByEmailAndDob(email, dob)
	if err != nil {
		log.Error().Err(err).Msg("couldnot find user")
		return "wrong creds given", errors.New("user not found")
	}

	otp := rand.Intn(10000)
	log.Info().Interface("----otp", otp).Send()
	log.Info().Interface("----email", uf.Email).Send()
	myString := strconv.Itoa(otp)

	context := context.Background()
	log.Info().Msg("hello")
	//err = s.re.AddToTheCacheOTP(ctx, uf.Email, otp)
	err = s.re.AddToTheCacheOTP(context, uf.Email, otp)
	if err != nil {
		log.Error().Err(err).Msg("couldnot add to cache")
		return "", errors.New("couldnt add to cache")
	}

	// Sender's email address and password
	from := "sridathkotturu7@gmail.com"
	password := "ncll nygy viwn upkr"

	// Recipient's email address
	to := "sridathkotturu6@gmail.com"

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	// Message content
	message := []byte("Subject: Test Email\n\nThis is a test email body." + myString)

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err = smtp.SendMail(smtpAddr, auth, from, []string{to}, message)
	if err != nil {
		log.Error().Err(err).Msg("couldnot send email")
		return "", err
	}

	return "sucessfull", nil

}

func (s *UserServiceImp) UserUpdatePassword(up model.UserUpdatePassword) (string, error) {
	context := context.Background()
	//string1, _ := s.re.GetTheCacheOTP(context, up.Email)

	hashed, err := bcrypt.GenerateFromPassword([]byte(up.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg("error occured in hashing password")
		return "hashing error", errors.New("hashing password failed")
	}
	string1, err := s.re.GetTheCacheOTP(context, up.Email)
	if err != nil {
		return "wouldnot get data from cache", err
	}

	if up.Password == up.RetypePassword && string1 == up.Otp {

		_, err := s.r.FetchUserByEmailAndUpdate(up.Email, string(hashed))
		if err != nil {
			log.Error().Msg("error occured in fetching")
			return "wrong creds/email ", err
		}
		return "successfull", nil
	}
	return "couldnot update,password/otp didnot match ", errors.New("wrong creds")
}
