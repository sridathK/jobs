package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"project/internal/auth"
	"project/internal/middlewear"
	"project/internal/model"
	"project/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type JobHandler struct {
	a  *auth.Auth
	cs services.CompanyService
}

func NewJobHandler(a *auth.Auth, cs services.CompanyService) (JobHandlerIn, error) {
	if cs == nil {
		return nil, errors.New("service implementation not given")
	}

	return &JobHandler{a: a, cs: cs}, nil

}

type JobHandlerIn interface {
	companyCreation(c *gin.Context)
	getAllCompany(c *gin.Context)
	getCompanyById(c *gin.Context)
	postJobByCompany(c *gin.Context)
	getJobByCompany(c *gin.Context)
	getAllJob(c *gin.Context)
	getJobByJobId(c *gin.Context)
	processingJobInput(c *gin.Context)
}

func (h *JobHandler) companyCreation(c *gin.Context) {

	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)

	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// _, ok = ctx.Value(middlewear.TokenIdKey).(jwt.RegisteredClaims)
	// if !ok {
	// 	//	log.Error().Str("Trace Id", traceid).Msg("login first")
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
	// 	return
	// }

	var companyCreation model.CreateCompany
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&companyCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&companyCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid input"})
		return
	}
	// Regclaims:=ctx.Value(middlewear.TokenIdKey)
	// companyCreation.User= int(Regclaims.Subject)

	us, err := h.cs.CompanyCreate(companyCreation)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("company creation problem in db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "compoany creation failed"})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *JobHandler) getAllCompany(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)

	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	us, err := h.cs.GetAllCompanies()
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("get all company problem from db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": " could not get all companies"})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *JobHandler) getCompanyById(c *gin.Context) {

	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	id, erro := strconv.Atoi(c.Param("company_id"))
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	us, err := h.cs.GetCompanyById(id)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("get company problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)
}

func (h *JobHandler) postJobByCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in  handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	id, erro := strconv.ParseUint(c.Param("company_id"), 10, 32)
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	var jobCreation model.CreateJob
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&jobCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&jobCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	us, err := h.cs.JobCreate(jobCreation, id)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("job creatuion problem in db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *JobHandler) getJobByCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in  handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	id, erro := strconv.Atoi(c.Param("company_id"))
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	us, err := h.cs.GetJobsByCompanyId(id)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("getting jobs problem fro db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *JobHandler) getAllJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	us, err := h.cs.GetAllJobs()
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("geting all jobs problem from db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)

}
func (h *JobHandler) getJobByJobId(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	id, erro := strconv.Atoi(c.Param("job_id"))
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	us, err := h.cs.GetJobByJobId(uint(id))
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("geting all jobs problem from db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)
}

func (h *JobHandler) processingJobInput(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in  handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var JobRequest []model.JobRequest

	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&JobRequest)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	// validate := validator.New()
	// err = validate.Struct(&JobRequest)
	// if err != nil {
	// 	log.Error().Err(err).Msg("error in validating ")
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
	// 	return
	// }

	us, err := h.cs.ProcessingJobDetails(JobRequest)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("geting all jobs problem from db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)
}
