package services

import (
	"errors"
	"project/internal/model"
	redisconn "project/internal/redisConn"
	"project/internal/repository"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

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
			//mockUserRepo := repository.NewMockUsers(mc)
			mockRedis := redisconn.NewMockCaching(mc)

			companyModel := model.Company{CompanyName: tt.args.nc.CompanyName, Adress: tt.args.nc.Adress, Domain: tt.args.nc.Domain}
			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().CreateCompany(companyModel).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewCompanyServiceImp(mockCompanyRepo, mockRedis)
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
			//mockUserRepo := repository.NewMockUsers(mc)
			mockRedis := redisconn.NewMockCaching(mc)

			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().GetCompany(tt.args.id).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewCompanyServiceImp(mockCompanyRepo, mockRedis)
			got, err := s.GetCompanyById(tt.args.id)

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
			//	mockUserRepo := repository.NewMockUsers(mc)\
			mockRedis := redisconn.NewMockCaching(mc)

			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().GetAllCompany().Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewCompanyServiceImp(mockCompanyRepo, mockRedis)
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

// func TestService_JobCreate(t *testing.T) {
// 	type args struct {
// 		nj model.CreateJob
// 		id uint64
// 	}
// 	tests := []struct {
// 		name string
// 		//s       *Service
// 		args             args
// 		want             model.Job
// 		wantErr          bool
// 		mockRepoResponse func() (model.Job, error)
// 	}{{
// 		name: "success",
// 		want: model.Job{JobTitle: "tcs", JobSalary: "bangalore", Uid: 1},
// 		args: args{
// 			nj: model.CreateJob{JobTitle: "tcs", JobSalary: "bangalore"},
// 			id: 1,
// 		},
// 		wantErr: false,
// 		mockRepoResponse: func() (model.Job, error) {
// 			return model.Job{JobTitle: "tcs", JobSalary: "bangalore", Uid: 1}, nil
// 		},
// 	},
// 		{name: "failure",
// 			want: model.Job{},
// 			args: args{
// 				nj: model.CreateJob{JobTitle: "", JobSalary: "bangalore"},
// 				id: 1,
// 			},
// 			wantErr: true,
// 			mockRepoResponse: func() (model.Job, error) {
// 				return model.Job{}, errors.New("test error")
// 			}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockCompanyRepo := repository.NewMockCompany(mc)

// 			mockUserRepo := repository.NewMockUsers(mc)
// 			job := model.Job{JobTitle: tt.args.nj.JobTitle, JobSalary: tt.args.nj.JobSalary, Uid: tt.args.id}
// 			if tt.mockRepoResponse != nil {
// 				mockCompanyRepo.EXPECT().CreateJob(job).Return(tt.mockRepoResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockUserRepo, mockCompanyRepo)
// 			got, err := s.JobCreate(tt.args.nj, tt.args.id)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.JobCreate() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.JobCreate() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

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
			//mockUserRepo := repository.NewMockUsers(mc)
			mockRedis := redisconn.NewMockCaching(mc)

			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().GetJobs(tt.args.id).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewCompanyServiceImp(mockCompanyRepo, mockRedis)
			got, err := s.GetJobsByCompanyId(tt.args.id)

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
			//mockUserRepo := repository.NewMockUsers(mc)
			mockRedis := redisconn.NewMockCaching(mc)

			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().GetAllJobs().Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewCompanyServiceImp(mockCompanyRepo, mockRedis)
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

func TestService_JobCreate(t *testing.T) {
	type args struct {
		nj model.CreateJob
		id uint64
	}
	tests := []struct {
		name string
		//	s       *Service
		args              args
		want              model.JobResponse
		wantErr           bool
		mockRepoResponse  func() (model.Job, error)
		mockRepoResponseL func() (model.Location, error)
		mockRepoResponseQ func() (model.Qualification, error)
		mockRepoResponseT func() (model.Technology, error)
		mockRepoResponseS func() (model.Shift, error)
		mockRepoResponseJ func() (model.JobType, error)
	}{
		{
			name: "success",
			want: model.JobResponse{ID: 0},
			args: args{
				nj: model.CreateJob{JobTitle: "tcs", JobSalary: "bangalore", MinNoticePeriod: "0", MaxNoticePeriod: "60", Budget: "800000", Description: "gorm", MinExperience: "1", MaxExperience: "3", JobLocations: []uint{1}, Qualification: []uint{1}, Technology: []uint{1}, Shift: []uint{1}, JobType: []uint{1}},
				id: 2,
			},
			wantErr:           false,
			mockRepoResponseL: func() (model.Location, error) { return model.Location{ID: 1, Name: "bangalore"}, nil },
			mockRepoResponseQ: func() (model.Qualification, error) { return model.Qualification{ID: 1, Name: "bcs-cse"}, nil },
			mockRepoResponseT: func() (model.Technology, error) { return model.Technology{ID: 1, Name: "java"}, nil },
			mockRepoResponseS: func() (model.Shift, error) { return model.Shift{ID: 1, Name: "morning-shift"}, nil },
			mockRepoResponseJ: func() (model.JobType, error) { return model.JobType{ID: 1, Name: "fulltime"}, nil },
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{JobTitle: "tcs", JobSalary: "bangalore", MinNoticePeriod: "0", MaxNoticePeriod: "60", Budget: "800000", Description: "gorm", MinExperience: "1", MaxExperience: "3"}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockCompanyRepo := repository.NewMockCompany(mc)
			//mockUserRepo := repository.NewMockUsers(mc)
			mockRedis := redisconn.NewMockCaching(mc)

			// job := model.Job{JobTitle: tt.args.nj.JobTitle, JobSalary: tt.args.nj.JobSalary, MinNoticePeriod: tt.args.nj.MinNoticePeriod,
			// 	MaxNoticePeriod: tt.args.nj.MinNoticePeriod, Budget: tt.args.nj.Budget,
			// 	Description: tt.args.nj.Description, MinExperience: tt.args.nj.MinExperience, MaxExperience: tt.args.nj.MaxExperience,
			// 	Uid: tt.args.id,JobLocations:[]model.Location{ID1,"hyderabad"}}

			//ID := uint(1)
			if tt.mockRepoResponseL != nil {
				mockCompanyRepo.EXPECT().GetLocationById(gomock.Any()).Return(tt.mockRepoResponseL()).AnyTimes()
			}

			if tt.mockRepoResponseQ != nil {
				mockCompanyRepo.EXPECT().GetQualificationById(gomock.Any()).Return(tt.mockRepoResponseQ()).AnyTimes()
			}

			if tt.mockRepoResponseT != nil {
				mockCompanyRepo.EXPECT().GetTechnologyById(gomock.Any()).Return(tt.mockRepoResponseT()).AnyTimes()
			}

			if tt.mockRepoResponseS != nil {
				mockCompanyRepo.EXPECT().GetShiftById(gomock.Any()).Return(tt.mockRepoResponseS()).AnyTimes()
			}

			if tt.mockRepoResponseJ != nil {
				mockCompanyRepo.EXPECT().GetJobTypeById(gomock.Any()).Return(tt.mockRepoResponseJ()).AnyTimes()
			}
			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().CreateJob(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewCompanyServiceImp(mockCompanyRepo, mockRedis)
			// gotL,err:=s.c.GetLocationById(tt.args.nj.JobLocations[0])
			// gotQ,err:=s.c.GetQualificationById(tt.args.nj.Qualification[0])
			// gotT,err:=s.c.GetTechnologyById(tt.args.nj.Technology[0])
			// gotS,err:=s.c.GetShById(tt.args.nj.Technology[0])
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

func TestService_ProcessingJobDetails(t *testing.T) {
	type args struct {
		m []model.JobRequest
	}
	tests := []struct {
		name string
		// s       *Service
		args                 args
		want                 []model.Job
		wantErr              bool
		mockRepoResponse     func() (model.Job, error)
		mockRedisGetResponse func() (string, error)
		mockRedisSetResponse func() error
	}{
		{
			name: "success",
			args: args{
				m: []model.JobRequest{{JobId: 5, JobTitle: "tcs", JobSalary: "234", NoticePeriod: 20, Budget: 800000, Description: "gorm",
					Experience: 2, JobLocations: []string{"hyderabad"}, Qualification: []string{"bcs-cse"}, WorkMode: "Remote",
					Technology: []string{"java"}, Shift: []string{"morning-shift"}, JobType: []string{"fulltime"}}},
			},
			want: []model.Job{{JobTitle: "tcs", JobSalary: "234", MinNoticePeriod: "0", MaxNoticePeriod: "30", Budget: "800000", Description: "gorm", MinExperience: "1", MaxExperience: "3", WorkMode: "Remote", JobLocations: []model.Location{{ID: 1, Name: "hyderabad"},
				{ID: 3, Name: "pune"}}, Technology: []model.Technology{{ID: 1, Name: "java"},
				{ID: 2, Name: "dsa"}}, Qualification: []model.Qualification{{ID: 1, Name: "bcs-cse"}}, Shift: []model.Shift{{ID: 1, Name: "morning-shift"}}, JobType: []model.JobType{{ID: 1, Name: "fulltime"}}, Uid: 1}},

			wantErr: false,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{JobTitle: "tcs", JobSalary: "234", MinNoticePeriod: "0", MaxNoticePeriod: "30", Budget: "800000", Description: "gorm", MinExperience: "1", MaxExperience: "3", WorkMode: "Remote", JobLocations: []model.Location{{ID: 1, Name: "hyderabad"},
					{ID: 3, Name: "pune"}}, Technology: []model.Technology{{ID: 1, Name: "java"},
					{ID: 2, Name: "dsa"}}, Qualification: []model.Qualification{{ID: 1, Name: "bcs-cse"}}, Shift: []model.Shift{{ID: 1, Name: "morning-shift"}}, JobType: []model.JobType{{ID: 1, Name: "fulltime"}}, Uid: 1}, nil
			},
			mockRedisGetResponse: func() (string, error) { return "", errors.New("e") },
			mockRedisSetResponse: func() error { return nil },
		},
		// {
		// 	name: "failure",
		// 	args: args{
		// 		m: []model.JobRequest{{JobId: 5, JobTitle: "tcs", JobSalary: "234", NoticePeriod: 20, Budget: 800000, Description: "gorm",
		// 			Experience: 2, JobLocations: []string{}, Qualification: []string{"mcs-cse"}, WorkMode: "Remote",
		// 			Technology: []string{"java"}, Shift: []string{"morning-shift"}, JobType: []string{"fulltime"}}},
		// 	},
		// 	want:    nil,
		// 	wantErr: false,
		// 	mockRepoResponse: func() (model.Job, error) {
		// 		return model.Job{JobTitle: "tcs", JobSalary: "234", MinNoticePeriod: "0", MaxNoticePeriod: "30", Budget: "800000", Description: "gorm", MinExperience: "1", MaxExperience: "3", WorkMode: "Remote", JobLocations: []model.Location{{ID: 1, Name: "hyderabad"},
		// 			{ID: 3, Name: "pune"}}, Technology: []model.Technology{{ID: 1, Name: "java"},
		// 			{ID: 2, Name: "dsa"}}, Qualification: []model.Qualification{{ID: 1, Name: "bcs-cse"}}, Shift: []model.Shift{{ID: 1, Name: "morning-shift"}}, JobType: []model.JobType{{ID: 1, Name: "fulltime"}}, Uid: 1}, nil
		// 	},
		// },
		{
			name: "success2",
			args: args{
				m: []model.JobRequest{{JobId: 5, JobTitle: "tcs", JobSalary: "234", NoticePeriod: 20, Budget: 800000, Description: "gorm",
					Experience: 2, JobLocations: []string{"hyderabad"}, Qualification: []string{"bcs-cse"}, WorkMode: "Remote",
					Technology: []string{"java"}, Shift: []string{"morning-shift"}, JobType: []string{"fulltime"}}},
			},
			want: []model.Job{{JobTitle: "tcs", JobSalary: "234", MinNoticePeriod: "0", MaxNoticePeriod: "30", Budget: "800000", Description: "gorm", MinExperience: "1", MaxExperience: "3", WorkMode: "Remote", JobLocations: []model.Location{{ID: 1, Name: "hyderabad"},
				{ID: 3, Name: "pune"}}, Technology: []model.Technology{{ID: 1, Name: "java"},
				{ID: 2, Name: "dsa"}}, Qualification: []model.Qualification{{ID: 1, Name: "bcs-cse"}}, Shift: []model.Shift{{ID: 1, Name: "morning-shift"}}, JobType: []model.JobType{{ID: 1, Name: "fulltime"}}, Uid: 1}},

			wantErr: false,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{JobTitle: "tcs", JobSalary: "234", MinNoticePeriod: "0", MaxNoticePeriod: "30", Budget: "800000", Description: "gorm", MinExperience: "1", MaxExperience: "3", WorkMode: "Remote", JobLocations: []model.Location{{ID: 1, Name: "hyderabad"},
					{ID: 3, Name: "pune"}}, Technology: []model.Technology{{ID: 1, Name: "java"},
					{ID: 2, Name: "dsa"}}, Qualification: []model.Qualification{{ID: 1, Name: "bcs-cse"}}, Shift: []model.Shift{{ID: 1, Name: "morning-shift"}}, JobType: []model.JobType{{ID: 1, Name: "fulltime"}}, Uid: 1}, nil
			},
			mockRedisGetResponse: func() (string, error) { return "", nil },
			mockRedisSetResponse: func() error { return nil },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockCompanyRepo := repository.NewMockCompany(mc)
			//mockUserRepo := repository.NewMockUsers(mc)
			mockRedis := redisconn.NewMockCaching(mc)

			if tt.mockRepoResponse != nil {
				mockCompanyRepo.EXPECT().GetJobsByJobId(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			if tt.mockRedisGetResponse != nil {
				mockRedis.EXPECT().GetTheCacheData(gomock.Any(), gomock.Any()).Return(tt.mockRedisGetResponse()).AnyTimes()
			}

			if tt.mockRedisSetResponse != nil {
				mockRedis.EXPECT().AddToTheCache(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockRedisSetResponse()).AnyTimes()
			}

			s, _ := NewCompanyServiceImp(mockCompanyRepo, mockRedis)
			got, err := s.ProcessingJobDetails(tt.args.m)

			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ProcessingJobDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, got, tt.want)
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Service.ProcessingJobDetails() = %v, want %v", got, tt.want)
			// }
		})
	}
}
