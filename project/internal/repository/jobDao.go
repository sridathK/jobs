package repository

import (
	"errors"
	"project/internal/model"

	"gorm.io/gorm"
)

type CompanyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) (Company, error) {
	if db == nil {
		return nil, errors.New("db connection not given")
	}

	return &CompanyRepo{db: db}, nil

}

//go:generate mockgen -source=jobDao.go -destination=companyrepository_mock.go -package=repository
type Company interface {
	CreateCompany(model.Company) (model.Company, error)
	GetAllCompany() ([]model.Company, error)
	GetCompany(id int) (model.Company, error)
	CreateJob(j model.Job) (model.Job, error)
	GetJobs(id int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
	GetJobsByJobId(id uint) (model.Job, error)
	GetLocationById(id uint) (model.Location, error)
	GetQualificationById(id uint) (model.Qualification, error)
	GetTechnologyById(id uint) (model.Technology, error)
	GetShiftById(id uint) (model.Shift, error)
	GetJobTypeById(id uint) (model.JobType, error)
}

func (r *CompanyRepo) CreateCompany(u model.Company) (model.Company, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return model.Company{}, err
	}
	return u, nil
}

func (r *CompanyRepo) GetAllCompany() ([]model.Company, error) {
	var s []model.Company
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *CompanyRepo) GetCompany(id int) (model.Company, error) {
	var m model.Company
	id1 := uint64(id)
	tx := r.db.Where("id = ?", id1)
	err := tx.Find(&m).Error
	if err != nil {
		return model.Company{}, err
	}
	return m, nil

}

func (r *CompanyRepo) CreateJob(j model.Job) (model.Job, error) {
	err := r.db.Create(&j).Error
	if err != nil {
		return model.Job{}, err
	}
	return j, nil
}

func (r *CompanyRepo) GetJobs(id int) ([]model.Job, error) {
	var m []model.Job

	tx := r.db.Where("uid = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil

}

func (r *CompanyRepo) GetAllJobs() ([]model.Job, error) {
	var s []model.Job
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *CompanyRepo) GetJobsByJobId(id uint) (model.Job, error) {
	var m model.Job

	tx := r.db.Preload("Qualification").Preload("JobLocations").Preload("Technology").Preload("Shift").Preload("JobType").Where("id = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return model.Job{}, err
	}
	return m, nil

}

func (r *CompanyRepo) GetLocationById(id uint) (model.Location, error) {
	var l model.Location

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&l).Error
	if err != nil {
		return model.Location{}, err
	}
	return l, nil

}

func (r *CompanyRepo) GetQualificationById(id uint) (model.Qualification, error) {
	var l model.Qualification

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&l).Error
	if err != nil {
		return model.Qualification{}, err
	}
	return l, nil
}

func (r *CompanyRepo) GetTechnologyById(id uint) (model.Technology, error) {
	var l model.Technology

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&l).Error
	if err != nil {
		return model.Technology{}, err
	}
	return l, nil
}

func (r *CompanyRepo) GetShiftById(id uint) (model.Shift, error) {
	var l model.Shift

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&l).Error
	if err != nil {
		return model.Shift{}, err
	}
	return l, nil
}

func (r *CompanyRepo) GetJobTypeById(id uint) (model.JobType, error) {
	var l model.JobType

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&l).Error
	if err != nil {
		return model.JobType{}, err
	}
	return l, nil
}
