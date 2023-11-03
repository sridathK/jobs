package services

import (
	"errors"
	"project/internal/model"
	"project/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_CompanyCreate(t *testing.T) {
	type args struct {
		nc model.CreateCompany
	}
	tests := []struct {
		name string
		//s       *Service
		args             args
		want             model.Company
		wantErr          bool
		mockRepoResponse func() (model.Company, error)
	}{
		{
			name: "success",
			want: model.Company{CompanyName: "tcs", Adress: "bangalore", Domain: "software"},
			args: args{
				nc: model.CreateCompany{CompanyName: "tcs", Adress: "bangalore", Domain: "software"},
			},
			wantErr: false,
			mockRepoResponse: func() (model.Company, error) {
				return model.Company{CompanyName: "tcs", Adress: "bangalore", Domain: "software"}, nil
			},
		},
		{
			name: "invalid input-failure case",
			want: model.Company{},
			args: args{
				nc: model.CreateCompany{CompanyName: "", Adress: "bangalore", Domain: "software"},
			},
			wantErr: true,
			mockRepoResponse: func() (model.Company, error) {
				return model.Company{}, errors.New("test error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockCompanyRepo := repository.NewMockCompany(mc)
			mockUserRepo := repository.NewMockUsers(mc)

			companyModel := model.Company{CompanyName: tt.args.nc.CompanyName, Adress: tt.args.nc.Adress, Domain: tt.args.nc.Domain}
			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().CreateCompany(companyModel).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockUserRepo, mockCompanyRepo)
			got, err := s.CompanyCreate(tt.args.nc)

			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CompanyCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.CompanyCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetCompany(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		//s       *Service
		args             args
		want             model.Company
		wantErr          bool
		mockRepoResponse func() (model.Company, error)
	}{
		{
			name: "success",
			want: model.Company{CompanyName: "tcs", Adress: "bangalore", Domain: "software"},
			args: args{
				id: 1,
			},
			wantErr: false,
			mockRepoResponse: func() (model.Company, error) {
				return model.Company{CompanyName: "tcs", Adress: "bangalore", Domain: "software"}, nil
			},
		},
		{
			name: "failure",
			want: model.Company{},
			args: args{
				id: 12,
			},
			wantErr: true,
			mockRepoResponse: func() (model.Company, error) {
				return model.Company{}, errors.New("id cannnot be greater")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockCompanyRepo := repository.NewMockCompany(mc)
			mockUserRepo := repository.NewMockUsers(mc)

			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().GetCompany(tt.args.id).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockUserRepo, mockCompanyRepo)
			got, err := s.GetCompany(tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAllCompanies(t *testing.T) {
	tests := []struct {
		name string
		//s       *Service
		want             []model.Company
		wantErr          bool
		mockRepoResponse func() ([]model.Company, error)
	}{
		{
			name:    "success",
			want:    []model.Company{{CompanyName: "tcs", Adress: "bangalore", Domain: "sde"}, {CompanyName: "Tek", Adress: "hyd", Domain: "sde"}},
			wantErr: false,
			mockRepoResponse: func() ([]model.Company, error) {
				return []model.Company{{CompanyName: "tcs", Adress: "bangalore", Domain: "sde"}, {CompanyName: "Tek", Adress: "hyd", Domain: "sde"}}, nil
			},
		},
		{
			name:    "failure",
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]model.Company, error) {
				return []model.Company{}, errors.New("no company created to get company's")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockCompanyRepo := repository.NewMockCompany(mc)
			mockUserRepo := repository.NewMockUsers(mc)

			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().GetAllCompany().Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockUserRepo, mockCompanyRepo)
			got, err := s.GetAllCompanies()

			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAllCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAllCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_JobCreate(t *testing.T) {
	type args struct {
		nj model.CreateJob
		id uint64
	}
	tests := []struct {
		name string
		//s       *Service
		args             args
		want             model.Job
		wantErr          bool
		mockRepoResponse func() (model.Job, error)
	}{{
		name: "success",
		want: model.Job{JobTitle: "tcs", JobSalary: "bangalore", Uid: 1},
		args: args{
			nj: model.CreateJob{JobTitle: "tcs", JobSalary: "bangalore"},
			id: 1,
		},
		wantErr: false,
		mockRepoResponse: func() (model.Job, error) {
			return model.Job{JobTitle: "tcs", JobSalary: "bangalore", Uid: 1}, nil
		},
	},
		{name: "failure",
			want: model.Job{},
			args: args{
				nj: model.CreateJob{JobTitle: "", JobSalary: "bangalore"},
				id: 1,
			},
			wantErr: true,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{}, errors.New("test error")
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockCompanyRepo := repository.NewMockCompany(mc)
			mockUserRepo := repository.NewMockUsers(mc)
			job := model.Job{JobTitle: tt.args.nj.JobTitle, JobSalary: tt.args.nj.JobSalary, Uid: tt.args.id}
			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().CreateJob(job).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockUserRepo, mockCompanyRepo)
			got, err := s.JobCreate(tt.args.nj, tt.args.id)

			if (err != nil) != tt.wantErr {
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
		name string
		//s       *Service
		args             args
		want             []model.Job
		wantErr          bool
		mockRepoResponse func() ([]model.Job, error)
	}{
		{
			name: "success",
			want: []model.Job{{JobTitle: "tcs", JobSalary: "bangalore", Uid: 1}},
			args: args{
				id: 1,
			},
			wantErr: false,
			mockRepoResponse: func() ([]model.Job, error) {
				return []model.Job{{JobTitle: "tcs", JobSalary: "bangalore", Uid: 1}}, nil
			},
		},
		{
			name: "failure",
			want: nil,
			args: args{
				id: 12,
			},
			wantErr: true,
			mockRepoResponse: func() ([]model.Job, error) {
				return nil, errors.New("id cannnot be greater")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockCompanyRepo := repository.NewMockCompany(mc)
			mockUserRepo := repository.NewMockUsers(mc)

			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().GetJobs(tt.args.id).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockUserRepo, mockCompanyRepo)
			got, err := s.GetJobs(tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAllJobs(t *testing.T) {
	tests := []struct {
		name string
		//s       *Service
		want             []model.Job
		wantErr          bool
		mockRepoResponse func() ([]model.Job, error)
	}{
		{
			name:    "success",
			want:    []model.Job{{JobTitle: "sde", JobSalary: "234", Uid: 1}, {JobTitle: "aws", JobSalary: "890", Uid: 2}},
			wantErr: false,
			mockRepoResponse: func() ([]model.Job, error) {
				return []model.Job{{JobTitle: "sde", JobSalary: "234", Uid: 1}, {JobTitle: "aws", JobSalary: "890", Uid: 2}}, nil
			},
		},
		{
			name:    "failure",
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]model.Job, error) {
				return []model.Job{}, errors.New("no jobs created to get jobs")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockCompanyRepo := repository.NewMockCompany(mc)
			mockUserRepo := repository.NewMockUsers(mc)

			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().GetAllJobs().Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockUserRepo, mockCompanyRepo)
			got, err := s.GetAllJobs()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAllJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAllJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
