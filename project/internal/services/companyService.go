package services

import (
	"errors"
	"fmt"
	"project/internal/model"
	"slices"
	"strconv"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=companyService.go -destination=companyservice_mock.go -package=services
type CompanyService interface {
	CompanyCreate(nc model.CreateCompany) (model.Company, error)
	GetAllCompanies() ([]model.Company, error)
	GetCompanyById(id int) (model.Company, error)
	JobCreate(nj model.CreateJob, id uint64) (uint, error)
	GetJobsByCompanyId(id int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
	GetJobByJobId(id uint) (model.Job, error)
	ProcessingJobDetails([]model.JobRequest) ([]model.Job, error)
}

func (s *Service) CompanyCreate(nc model.CreateCompany) (model.Company, error) {
	company := model.Company{CompanyName: nc.CompanyName, Adress: nc.Adress, Domain: nc.Domain}
	cu, err := s.c.CreateCompany(company)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create company")
		return model.Company{}, errors.New("company creation failed")
	}

	return cu, nil
}

func (s *Service) GetAllCompanies() ([]model.Company, error) {

	AllCompanies, err := s.c.GetAllCompany()
	if err != nil {
		log.Error().Err(err).Msg("couldnot get companies")
		return nil, err
	}
	return AllCompanies, nil

}

func (s *Service) GetCompanyById(id int) (model.Company, error) {
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

func (s *Service) JobCreate(nj model.CreateJob, id uint64) (uint, error) {
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
		return 0, errors.New("job creation failed")
	}

	return cu, nil
}

func (s *Service) GetJobsByCompanyId(id int) ([]model.Job, error) {
	if id > 10 {
		return nil, errors.New("id cannnot be greater")
	}
	AllCompanies, err := s.c.GetJobs(id)
	if err != nil {
		return nil, errors.New("job retreval failed")
	}
	return AllCompanies, nil
}

func (s *Service) GetAllJobs() ([]model.Job, error) {

	AllJobs, err := s.c.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return AllJobs, nil

}

func (s *Service) GetJobByJobId(id uint) (model.Job, error) {
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

func (s *Service) ProcessingJobDetails(m []model.JobRequest) ([]model.Job, error) {
	//wg := new(sync.WaitGroup)
	var JobsResult []model.Job
	// var JobRequest []model.JobRequest
	fmt.Println(s.c.GetJobsByJobId(6))
	for _, v := range m {
		Jobs, err := s.c.GetJobsByJobId(v.JobId)
		if err != nil {
			return nil, errors.New("couldnot retrieve 'all jobs' from db")
		}
		for _, v := range m {
			result1 := Processing5Data(Jobs, v)
			result2 := ProcessingOther5Data(Jobs, v)

			if result1 && result2 {
				JobsResult = append(JobsResult, Jobs)
				//return JobsResult, nil
			}
		}

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
