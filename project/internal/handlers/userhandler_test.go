package handlers

import (
	"context"
	"errors"
	"io"
	"net/http"
	"project/internal/model"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandlerMock struct{}

func (us *UserHandlerMock) UserSignup(nu model.UserSignup) (model.User, error) {
	if nu.UserName == "" {
		return model.User{}, errors.New("user signup failed")
	}
	return model.User{UserName: nu.UserName, Email: nu.Email}, nil

}
func initialization() {
	var c1 *gin.Context
	var c2 *gin.Context
	var c3 *gin.Context
	type key string
	const TraceIdKey key = "1"
	////////
	traceId := 1
	ctx := c1.Request.Context()
	ctx = context.WithValue(ctx, TraceIdKey, traceId)
	req := c1.Request.WithContext(ctx)
	c1.Request = req

	traceId1 := uuid.NewString()
	ctx1 := c2.Request.Context()
	ctx1 = context.WithValue(ctx1, TraceIdKey, traceId1)
	req = c2.Request.WithContext(ctx1)
	c2.Request = req

	Body := io.ReadCloser(model.UserSignup{
		UserName: "",
		Email:    "name@gmail.com",
		Password: "hfhhfhfh",
	})
	c2.Request.Body = Body
	Body.Close()

	traceId2 := uuid.NewString()
	ctx := c3.Request.Context()
	ctx = context.WithValue(ctx, TraceIdKey, traceId2)
	req := c3.Request.WithContext(ctx)
	c3.Request = req

	Body := []byte(model.UserSignup{
		"name":     "names",
		"email":    "name@gmail.com",
		"password": "hfhhfhfh",
	})
	c3.Request.Body = Body

}


func Test_handler_userSignin(t *testing.T) {
	type args struct {
		c *gin.Context
	
	}
	r1,_:=http.NewRequest(http.MethodPost,"",strings.NewReader(`{"name":     "names",
	"email":    "name@gmail.com",
	"password": "hfhhfhfh"}`))
	
	tests := []struct {
		name string
		h    *handler
		args args
	}{
		{name: "failure in traceid",
			h:    &handler{us: &UserHandlerMock{}},
			args: args{c: &gin.Context{Request:r1}}},
		},
		{
			name: " failure in validation",
			h:    &handler{us: &UserHandlerMock{}},
			args: args{c: c2},
		},
		{
			name: "failure",
			h:    &handler{us: &UserHandlerMock{}},
			args: args{c: c3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.userSignin(tt.args.c)

		})
	}
}
