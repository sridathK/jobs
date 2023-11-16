package handlers

import (
	"project/internal/auth"
	"project/internal/middlewear"
	"project/internal/services"

	"github.com/gin-gonic/gin"
)

func Api(a *auth.Auth, su services.UsersService, sc services.CompanyService) *gin.Engine {
	r := gin.New()
	h1, _ := NewUserHandler(a, su)
	h2, _ := NewJobHandler(a, sc)
	//h, _ := NewHandler(a, su, sc)
	//h := handler{a: a, us: s, cs: s}
	m, _ := middlewear.NewMiddleWear(a)
	r.Use(m.Log(), gin.Recovery())
	r.POST("/api/register", h1.userSignup)
	r.POST("/api/login", h1.userLogin)
	r.POST("/api/companies", m.Auth(h2.companyCreation))
	r.GET("/api/companies", m.Auth(h2.getAllCompany))
	r.GET(" /api/company/:company_id", m.Auth(h2.getCompanyById))
	r.POST("/api/companies/:company_id/jobs", h2.postJobByCompany)
	r.GET("/api/companies/:company_id/jobs", m.Auth(h2.getJobByCompany))
	r.GET("/api/jobs", m.Auth(h2.getAllJob))
	r.GET("/api/jobs/:job_id", m.Auth(h2.getJobByJobId))
	r.POST("/api/jobs/processing", h2.processingJobInput)
	return r
}
