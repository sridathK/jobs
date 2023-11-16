package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"project/internal/model"
	redisconn "project/internal/redisConn"
	"project/internal/repository"
	"slices"
	"strconv"
	"sync"

	"github.com/rs/zerolog/log"
)

type CompanyServiceImp struct {
	c  repository.Company
	re redisconn.Caching
}

func NewCompanyServiceImp(c repository.Company, re redisconn.Caching) (CompanyService, error) {
	if c == nil {
		return nil, errors.New("db connection not given")
	}

	return &CompanyServiceImp{c: c, re: re}, nil

}

//go:generate mockgen -source=companyService.go -destination=companyservice_mock.go -package=services
type CompanyService interface {
	CompanyCreate(nc model.CreateCompany) (model.Company, error)
	GetAllCompanies() ([]model.Company, error)
	GetCompanyById(id int) (model.Company, error)
	JobCreate(nj model.CreateJob, id uint64) (model.JobResponse, error)
	GetJobsByCompanyId(id int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
	GetJobByJobId(id uint) (model.Job, error)
	ProcessingJobDetails([]model.JobRequest) ([]model.Job, error)
}

func (s *CompanyServiceImp) CompanyCreate(nc model.CreateCompany) (model.Company, error) {
	company := model.Company{CompanyName: nc.CompanyName, Adress: nc.Adress, Domain: nc.Domain}
	cu, err := s.c.CreateCompany(company)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create company")
		return model.Company{}, errors.New("company creation failed")
	}

	return cu, nil
}

func (s *CompanyServiceImp) GetAllCompanies() ([]model.Company, error) {

	AllCompanies, err := s.c.GetAllCompany()
	if err != nil {
		log.Error().Err(err).Msg("couldnot get companies")
		return nil, err
	}
	return AllCompanies, nil

}

func (s *CompanyServiceImp) GetCompanyById(id int) (model.Company, error) {
	if id > 10 {
		return model.Company{}, errors.New("id cannnot be greater")
	}
	Companies, err := s.c.GetCompany(id)
	if err != nil {
		log.Error().Err(err).Msg("couldnot get company")
		return model.Company{}, err
	}
	return Companies, nil

}

func (s *CompanyServiceImp) JobCreate(nj model.CreateJob, id uint64) (model.JobResponse, error) {
	// var response model.JobResponse
	var locations []model.Location
	var qualifications []model.Qualification
	var technologies []model.Technology
	var shifts []model.Shift
	var JobTypes []model.JobType
	for _, v := range nj.JobLocations {
		location, _ := s.c.GetLocationById(v)
		locations = append(locations, location)
	}
	for _, v := range nj.Qualification {
		qualification, _ := s.c.GetQualificationById(v)
		qualifications = append(qualifications, qualification)
	}

	for _, v := range nj.Technology {
		technology, _ := s.c.GetTechnologyById(v)
		technologies = append(technologies, technology)
	}

	for _, v := range nj.Shift {
		shift, _ := s.c.GetShiftById(v)
		shifts = append(shifts, shift)
	}

	for _, v := range nj.JobType {
		jobType, _ := s.c.GetJobTypeById(v)
		JobTypes = append(JobTypes, jobType)
	}

	job := model.Job{JobTitle: nj.JobTitle, JobSalary: nj.JobSalary, Uid: id, MinNoticePeriod: nj.MinNoticePeriod, MaxNoticePeriod: nj.MaxNoticePeriod, Budget: nj.Budget, WorkMode: nj.WorkMode,
		Description: nj.Description, MinExperience: nj.MinExperience, MaxExperience: nj.MaxExperience, JobLocations: locations, Technology: technologies, Qualification: qualifications, Shift: shifts, JobType: JobTypes}

	cu, err := s.c.CreateJob(job)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create job")
		return model.JobResponse{}, errors.New("job creation failed")
	}
	res := model.JobResponse{ID: cu.ID}
	return res, nil
}

func (s *CompanyServiceImp) GetJobsByCompanyId(id int) ([]model.Job, error) {
	if id > 10 {
		return nil, errors.New("id cannnot be greater")
	}
	AllCompanies, err := s.c.GetJobs(id)
	if err != nil {
		return nil, errors.New("job retreval failed")
	}
	return AllCompanies, nil
}

func (s *CompanyServiceImp) GetAllJobs() ([]model.Job, error) {

	AllJobs, err := s.c.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return AllJobs, nil

}

func (s *CompanyServiceImp) GetJobByJobId(id uint) (model.Job, error) {
	if id > 10 {
		return model.Job{}, errors.New("id cannnot be greater")
	}
	Companies, err := s.c.GetJobsByJobId(id)
	if err != nil {
		log.Error().Err(err).Msg("couldnot get company")
		return model.Job{}, err
	}
	return Companies, nil

}

// func (s *Service) ProcessingJobDetails(m []model.JobRequest) ([]model.Job, error) {
// 	var job []model.Job
// 	var allJobs []model.Job
// 	allJobs, err := s.c.GetAllJobs()

// 	if err != nil {
// 		return nil, errors.New("couldnot retrieve 'all jobs' from db")
// 	}
// 	c1 := make(chan bool)
// 	c2 := make(chan bool)
// 	wg := new(sync.WaitGroup)
// 	for _, v1 := range m {
// 		wg.Add(1)
// 		go func(jr model.JobRequest) {
// 			defer wg.Done()
// 			for _, v := range allJobs {
// 				//wg.Add(1)
// 				go func(v model.Job) {
// 					//defer wg.Done()
// 					result := Processing5Data(v, jr)
// 					c1 <- result
// 				}(v)
// 				//wg.Add(1)
// 				go func(v model.Job) {
// 					//defer wg.Done()
// 					result := ProcessingOther5Data(v, jr)
// 					c2 <- result
// 				}(v)
// 				output1 := <-c1
// 				output2 := <-c2
// 				if output1 && output2 {
// 					job = append(job, v)
// 				}

// 			}
// 		}(v1)
// 	}
// 	wg.Wait()
// 	close(c1)
// 	return job, nil
// }

func (s *CompanyServiceImp) ProcessingJobDetails(m []model.JobRequest) ([]model.Job, error) {
	wg := new(sync.WaitGroup)
	// wg1 := new(sync.WaitGroup)
	c1 := make(chan model.Job)
	var JobsResult []model.Job
	var Jobs model.Job
	//var JobsMap map[uint]model.Job                       Using maps as memory storage
	//JobsMap := make(map[uint]model.Job)
	// for _, v := range m {
	// 	_, ok := JobsMap[v.JobId]
	// 	if !ok {
	// 		Jobs, _ := s.c.GetJobsByJobId(v.JobId)
	// 		JobsMap[v.JobId] = Jobs
	// 	}
	// }
	context := context.Background()

	for _, v := range m {
		// using redis as In memory database.
		//jobId := string(v.JobId)
		//jobId := strconv.FormatUint(uint64(v.JobId), 10)
		_, err := s.re.GetTheCacheData(context, v.JobId)
		//_, err := redis.Get(context, jobId).Result()
		if err != nil {
			Jobs, err := s.c.GetJobsByJobId(v.JobId)
			if err != nil {
				fmt.Println("error getting data from databse", err)
				return nil, errors.New("couldnot get data")
			}
			//JobsJson, err := json.Marshal(Jobs)

			err = s.re.AddToTheCache(context, v.JobId, Jobs)
			//err = redis.Set(context, jobId, JobsJson, 0).Err()
			if err != nil {
				return nil, err
			}
		}
	}

	go func() {
		for _, v := range m {
			wg.Add(1)
			go func(v model.JobRequest) {
				//jobId := strconv.FormatUint(uint64(v.JobId), 10)
				defer wg.Done()
				//val, err := redis.Get(context, jobId).Result()
				val, err := s.re.GetTheCacheData(context, v.JobId)
				if err != nil {
					log.Error().Msg("couldnot get from redis")
					return
				}
				// wg1.Add(1)
				//err := json.NewDecoder(val).Decode(&Jobs)
				err = json.Unmarshal([]byte(val), &Jobs)
				if err != nil {
					fmt.Println("Error Unmarshalling JSON:", err)
					return
				}
				result1 := Processing5Data(Jobs, v)
				result2 := ProcessingOther5Data(Jobs, v)
				//wg1.Wait()

				if result1 && result2 {
					//JobsResult = append(JobsResult, Jobs)
					c1 <- Jobs
				}

			}(v)

		}
		wg.Wait()
		close(c1)
	}()
	for val := range c1 {
		JobsResult = append(JobsResult, val)
	}

	return JobsResult, nil
}

func Processing5Data(v model.Job, m model.JobRequest) bool {
	minNP, err := strconv.Atoi(v.MinNoticePeriod)
	if err != nil {
		panic("string to int error")
	}
	maxNP, err := strconv.Atoi(v.MaxNoticePeriod)
	if err != nil {
		panic("string to int error")
	}
	Budget, err := strconv.Atoi(v.Budget)
	if err != nil {
		panic("string to int error")
	}

	minEXP, err := strconv.Atoi(v.MinExperience)
	if err != nil {
		panic("string to int error")
	}
	maxEXP, err := strconv.Atoi(v.MaxExperience)
	if err != nil {
		panic("string to int error")
	}

	if !(minNP <= m.NoticePeriod && m.NoticePeriod <= maxNP) {
		return false
	}
	if !(0 < m.Budget && m.Budget <= Budget) {
		return false
	}
	if !(minEXP <= m.Experience && m.Experience <= maxEXP) {
		return false
	}
	if !(containsLocation(v.JobLocations, m.JobLocations)) {
		return false
	}
	if !(containsQualification(v.Qualification, m.Qualification)) {
		return false
	}

	return true

}
func ProcessingOther5Data(v model.Job, m model.JobRequest) bool {
	if !(containsTechn(v.Technology, m.Technology)) {
		return false
	}
	if !(containsShift(v.Shift, m.Shift)) {
		return false
	}
	if !(containsJobType(v.JobType, m.JobType)) {
		return false
	}
	if !(v.WorkMode == m.WorkMode) {
		return false
	}

	return true
}

func containsLocation(job []model.Location, request []string) bool {

	var lv []string

	for _, a := range job {
		lv = append(lv, a.Name)
	}
	for _, b := range request {
		if !slices.Contains(lv, b) {
			return false
		}
	}
	return true
}

func containsQualification(job []model.Qualification, request []string) bool {
	var lv []string

	for _, a := range job {
		lv = append(lv, a.Name)
	}
	for _, b := range request {
		if !slices.Contains(lv, b) {
			return false
		}
	}
	return true
}

func containsTechn(job []model.Technology, request []string) bool {
	var lv []string

	for _, a := range job {
		lv = append(lv, a.Name)
	}
	for _, b := range request {
		if !slices.Contains(lv, b) {
			return false
		}
	}
	return true
}

func containsShift(job []model.Shift, request []string) bool {
	var lv []string

	for _, a := range job {
		lv = append(lv, a.Name)
	}
	for _, b := range request {
		if !slices.Contains(lv, b) {
			return false
		}
	}
	return true
}

func containsJobType(job []model.JobType, request []string) bool {
	var lv []string

	for _, a := range job {
		lv = append(lv, a.Name)
	}
	for _, b := range request {
		if !slices.Contains(lv, b) {
			return false
		}
	}
	return true
}
