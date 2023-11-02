package handlers

import (
	"errors"
	"net/http"
	"project/internal/model"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandlerMock struct{}

func (us *UserHandlerMock) UserSignup(nu model.UserSignup) (model.User, error) {
	if nu.UserName == "" {
		return model.User{}, errors.New("user signup failed")
	}
	return model.User{UserName: nu.UserName, Email: nu.Email}, nil

}
func (us *UserHandlerMock) Userlogin(l model.UserLogin) (jwt.RegisteredClaims, error) {
	return jwt.RegisteredClaims{}, nil
}
func (us *UserHandlerMock) CompanyCreate(nc model.CreateCompany) (model.Company, error) {
	return model.Company{}, nil
}
func (us *UserHandlerMock) GetAllCompanies() ([]model.Company, error) {
	return nil, nil
}
func (us *UserHandlerMock) GetCompany(id int) (model.Company, error) {
	return model.Company{}, nil
}
func (us *UserHandlerMock) JobCreate(nj model.CreateJob, id uint64) (model.Job, error) {
	return model.Job{}, nil
}
func (us *UserHandlerMock) GetJobs(id int) ([]model.Job, error) {
	return nil, nil
}
func (us *UserHandlerMock) GetAllJobs() ([]model.Job, error) {
	return nil, nil
}

func Test_handler_userSignin(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	r1, _ := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"name":     "names",
	"email":    "name@gmail.com",
	"password": "hfhhfhfh"}`))

	tests := []struct {
		name string
		h    *handler
		args args
	}{
		{name: "failure in traceid",
			h:    &handler{us: &UserHandlerMock{}},
			args: args{c: &gin.Context{Request: r1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.userSignin(tt.args.c)

		})
	}
}
