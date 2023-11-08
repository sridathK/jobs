package repository

import "project/internal/model"

//go:generate mockgen -source=jobDao.go -destination=companyrepository_mock.go -package=repository
type Company interface {
	CreateCompany(model.Company) (model.Company, error)
	GetAllCompany() ([]model.Company, error)
	GetCompany(id int) (model.Company, error)
	CreateJob(j model.Job) (uint, error)
	GetJobs(id int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
	GetJobsByJobId(id uint) (model.Job, error)
	GetLocationById(id uint) (model.Location, error)
	GetQualificationById(id uint) (model.Qualification, error)
	GetTechnologyById(id uint) (model.Technology, error)
	GetShiftById(id uint) (model.Shift, error)
	GetJobTypeById(id uint) (model.JobType, error)
}

func (r *Repo) CreateCompany(u model.Company) (model.Company, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return model.Company{}, err
	}
	return u, nil
}

func (r *Repo) GetAllCompany() ([]model.Company, error) {
	var s []model.Company
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *Repo) GetCompany(id int) (model.Company, error) {
	var m model.Company
	id1 := uint64(id)
	tx := r.db.Where("id = ?", id1)
	err := tx.Find(&m).Error
	if err != nil {
		return model.Company{}, err
	}
	return m, nil

}

func (r *Repo) CreateJob(j model.Job) (uint, error) {
	err := r.db.Create(&j).Error
	if err != nil {
		return 0, err
	}
	return j.ID, nil
}

func (r *Repo) GetJobs(id int) ([]model.Job, error) {
	var m []model.Job

	tx := r.db.Where("uid = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil

}

func (r *Repo) GetAllJobs() ([]model.Job, error) {
	var s []model.Job
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *Repo) GetJobsByJobId(id uint) (model.Job, error) {
	var m model.Job

	tx := r.db.Preload("Qualification").Preload("JobLocations").Preload("Technology").Preload("Shift").Preload("JobType").Where("id = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return model.Job{}, err
	}
	return m, nil

}

func (r *Repo) GetLocationById(id uint) (model.Location, error) {
	var l model.Location

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&l).Error
	if err != nil {
		return model.Location{}, err
	}
	return l, nil

}

func (r *Repo) GetQualificationById(id uint) (model.Qualification, error) {
	var l model.Qualification

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&l).Error
	if err != nil {
		return model.Qualification{}, err
	}
	return l, nil
}

func (r *Repo) GetTechnologyById(id uint) (model.Technology, error) {
	var l model.Technology

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&l).Error
	if err != nil {
		return model.Technology{}, err
	}
	return l, nil
}

func (r *Repo) GetShiftById(id uint) (model.Shift, error) {
	var l model.Shift

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&l).Error
	if err != nil {
		return model.Shift{}, err
	}
	return l, nil
}

func (r *Repo) GetJobTypeById(id uint) (model.JobType, error) {
	var l model.JobType

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&l).Error
	if err != nil {
		return model.JobType{}, err
	}
	return l, nil
}
