package services

import (
	"errors"
	"project/internal/model"
	"reflect"
	"testing"
)

type MockUserSignup struct{}

func (m *MockUserSignup) CreateUser(m1 model.User) (model.User, error) {
	if m1.UserName == "" {
		return model.User{}, errors.New("incorrect input")
	}
	return model.User{UserName: m1.UserName, Email: m1.Email}, nil

}
func (m *MockUserSignup) FetchUserByEmail(string) (model.User, error) {

	return model.User{}, nil
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
