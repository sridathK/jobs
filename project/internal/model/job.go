package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	JobTitle        string          `json:"job_title" validate:"required"`
	JobSalary       string          `json:"job_salary" validate:"required"`
	Company         Company         `gorm:"ForeignKey:uid"`
	Uid             uint64          `JSON:"uid, omitempty"`
	MinNoticePeriod string          `json:"min_np"  validate:"required"`
	MaxNoticePeriod string          `json:"max_np"  validate:"required"`
	Budget          string          `json:"budget" validate:"required"`
	Description     string          ` json:"description" validate:"required"`
	MinExperience   string          ` json:"Min_exp" validate:"required"`
	MaxExperience   string          `json:"Max_exp" validate:"required"`
	JobLocations    []Location      `gorm:"many2many:job_location" json:"job_locations"`
	Technology      []Technology    `gorm:"many2many:job_techs" json:"technologies"`
	WorkMode        string          ` json:"work_mode" validate:"required"`
	Qualification   []Qualification `gorm:"many2many:job_qualifications" json:"job_quals"`
	Shift           []Shift         `gorm:"many2many:shift" json:"job_shifts"`
	JobType         []JobType       `gorm:"many2many:jobtypes" json:"job_type"`
}

type Location struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name" gorm:"unique"`
}

type Technology struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name" gorm:"unique"`
}

type Qualification struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name" gorm:"unique"`
}

type Shift struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name" gorm:"unique"`
}

type JobType struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name" gorm:"unique"`
}

type Company struct {
	gorm.Model
	CompanyName string `json:"company_name"  gorm:"unique"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

type CreateCompany struct {
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

type CreateJob struct {
	JobTitle  string `json:"job_title" validate:"required"`
	JobSalary string `json:"job_salary" validate:"required"`
	// Company         Company         `gorm:"ForeignKey:uid"`
	// Uid             uint64          `JSON:"uid, omitempty"`
	MinNoticePeriod string `json:"min_np"  validate:"required"`
	MaxNoticePeriod string `json:"max_np"  validate:"required"`
	Budget          string `json:"budget" validate:"required"`
	Description     string ` json:"description" validate:"required"`
	MinExperience   string ` json:"Min_exp" validate:"required"`
	MaxExperience   string `json:"Max_exp" validate:"required"`
	JobLocations    []uint `json:"job_locations"`
	Technology      []uint `json:"technologies"`
	WorkMode        string `json:"work_mode"`
	Qualification   []uint `json:"qualification"`
	Shift           []uint `json:"shift"`
	JobType         []uint `json:"job_type"`
}

type JobRequest struct {
	JobTitle      string   `json:"job_title" validate:"required"`
	JobSalary     string   `json:"job_salary" validate:"required"`
	NoticePeriod  int      `json:"np"  validate:"required"`
	JobId         uint     `json:"job_id"  validate:"required"`
	Budget        int      ` json:"budget" validate:"required"`
	JobLocations  []string ` json:"job_locations" validate:"required"`
	Technology    []string ` json:"tech_stack" validate:"required"`
	WorkMode      string   ` json:"work_mode" validate:"required"`
	Description   string   ` json:"description" validate:"required"`
	Experience    int      ` json:"exp" validate:"required"`
	Qualification []string `json:"qualification" validate:"required"`
	Shift         []string ` json:"shift" validate:"required"`
	JobType       []string ` json:"job_type" validate:"required"`
}

type JobResponse struct {
	ID uint
}
