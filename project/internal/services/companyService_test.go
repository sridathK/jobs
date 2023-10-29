package services

import (
	"errors"
	"project/internal/model"
	"reflect"
	"testing"
)

type MockJobs struct{}

func (m *MockJobs) CreateJob(j model.Job) (model.Job, error) {
	if j.JobTitle == "" {
		return model.Job{}, errors.New("job creation failed")
	}
	return model.Job{JobTitle: j.JobTitle, JobSalary: j.JobSalary, Uid: j.Uid}, nil
}

func (m *MockJobs) GetJobs(id int) ([]model.Job, error) {
	if id == 5 {
		return []model.Job{{JobTitle: "developer", JobSalary: "20k", Uid: 5}}, nil
	}

	return nil, errors.New("job retreval failed")
}
func (m *MockJobs) CreateCompany(model.Company) (model.Company, error) { return model.Company{}, nil }
func (m *MockJobs) GetAllCompany() ([]model.Company, error)            { return nil, nil }
func (m *MockJobs) GetCompany(id int) (model.Company, error)           { return model.Company{}, nil }
func (m *MockJobs) GetAllJobs() ([]model.Job, error)                   { return nil, nil }

func TestService_JobCreate(t *testing.T) {
	type args struct {
		nj model.CreateJob
		id uint64
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    model.Job
		wantErr error
	}{
		{name: "CASE Sucess",
			s:       &Service{c: &MockJobs{}},
			args:    args{nj: model.CreateJob{JobTitle: "developer", JobSalary: "20k"}, id: 5},
			want:    model.Job{JobTitle: "developer", JobSalary: "20k", Uid: 5},
			wantErr: nil,
		},
		{name: "CASE Failure",
			s:       &Service{c: &MockJobs{}},
			args:    args{nj: model.CreateJob{JobTitle: "", JobSalary: ""}, id: 5},
			want:    model.Job{},
			wantErr: errors.New("job creation failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.JobCreate(tt.args.nj, tt.args.id)
			if err != tt.wantErr {
				t.Errorf("Service.JobCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.JobCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetJobs(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    []model.Job
		wantErr error
	}{
		{name: "CASE Sucess",
			s:       &Service{c: &MockJobs{}},
			args:    args{id: 5},
			want:    []model.Job{{JobTitle: "developer", JobSalary: "20k", Uid: 5}},
			wantErr: nil,
		},
		{name: "CASE Failure",
			s:       &Service{c: &MockJobs{}},
			args:    args{id: 7},
			want:    nil,
			wantErr: errors.New("job retreval failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetJobs(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("Service.GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
