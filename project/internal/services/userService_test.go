package services

import (
	"errors"
	"project/internal/model"
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MockUserSignup struct{}

func (m *MockUserSignup) CreateUser(m1 model.User) (model.User, error) {
	if m1.UserName == "" {
		return model.User{}, errors.New("incorrect input")
	}
	return model.User{UserName: m1.UserName, Email: m1.Email}, nil

}
func (m *MockUserSignup) FetchUserByEmail(s string) (model.User, error) {
	if s == "" {
		return model.User{}, errors.New("incorrect input")
	}

	return model.User{UserName: "names", Email: "name@gmail.com", PasswordHash: "$2a$10$votXUqKwkXe6l5.2aVKSU.08QEPzZYuXy47OP7JuHebrZSppBlYSW"}, nil
}
func TestService_UserSignup(t *testing.T) {
	// type args struct {
	// 	nu model.UserSignup
	// }
	tests := []struct {
		name    string
		s       *Service
		nu      model.UserSignup
		want    model.User
		wantErr error
	}{
		{name: "checking mocked  sucess",
			s:       &Service{r: &MockUserSignup{}},
			nu:      model.UserSignup{UserName: "names", Email: "name@gmail.com", Password: "hfhhfhfh"},
			want:    model.User{UserName: "names", Email: "name@gmail.com"},
			wantErr: nil,
		},

		{name: "checking mocked failure ",
			s:       &Service{r: &MockUserSignup{}},
			nu:      model.UserSignup{UserName: "", Email: "", Password: ""},
			want:    model.User{},
			wantErr: errors.New("user creation failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UserSignup(tt.nu)
			if err != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ///////////    test case for login ///////////////////////
func TestService_Userlogin(t *testing.T) {

	tests := []struct {
		name    string
		s       *Service
		nu      model.UserLogin
		want    jwt.RegisteredClaims
		wantErr error
	}{
		{name: "checking  sucess case",
			s:       &Service{r: &MockUserSignup{}},
			nu:      model.UserLogin{Email: "name@gmail.com", Password: "hfhhfhfh"},
			want:    jwt.RegisteredClaims{Issuer: "service project", Subject: "0", Audience: jwt.ClaimStrings{"users"}, ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), IssuedAt: jwt.NewNumericDate(time.Now())},
			wantErr: nil,
		},
		{name: "checking  failure case",
			s:       &Service{r: &MockUserSignup{}},
			nu:      model.UserLogin{Email: "", Password: ""},
			want:    jwt.RegisteredClaims{},
			wantErr: errors.New("= user login failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Userlogin(tt.nu)
			if err != tt.wantErr {
				t.Errorf("Service.Userlogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Userlogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
