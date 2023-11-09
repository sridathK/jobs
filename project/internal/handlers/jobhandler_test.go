package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"project/internal/middlewear"
	"project/internal/model"
	"project/internal/services"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func Test_handler_companyCreation(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{name: "input validation",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"company_name":"",
			"domain": "hfhhfhfh"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"invalid input"}`,
		}, {
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"company_name":"namevs",
				"company_adress":    "name@gmail.com",
				"domain": "hfhhfhfh"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().CompanyCreate(gomock.Any()).Return(model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"","company_adress":"","domain":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.companyCreation(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_getAllCompany(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		}, {
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetAllCompanies().Return([]model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.getAllCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_getCompanyById(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		}, {
			name: "error while fetching companies from service",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "abc"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetCompanyById(gomock.Any()).Return(model.Company{}, errors.New("test service error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetCompanyById(gomock.Any()).Return(model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"","company_adress":"","domain":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.getCompanyById(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_postJob(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "missing param",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{name: "input validation",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"job_title":"",
			"job_salary":    "16786",}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		}, {
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"job_title":"rvrvrv",
				"job_salary":    "16786"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().JobCreate(gomock.Any(), gomock.Any()).Return(model.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"job_title":"","job_salary":"","Company":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"","company_adress":"","domain":""},"Uid":0}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.postJobByCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_getJobByCompany(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		}, {
			name: "error while fetching jobs from service",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "abc"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetJobsByCompanyId(gomock.Any()).Return([]model.Job{}, errors.New("test service error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetJobsByCompanyId(gomock.Any()).Return([]model.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.getJobByCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_getAllJob(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		}, {
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetAllJobs().Return([]model.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.getAllJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_processingJobInput(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedResponse   string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "error decoding",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"job_title":"rvrvrv",
				"job_salary":    "16786"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().ProcessingJobDetails(gomock.Any()).Return([]model.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`[{
					"job_title": "Go Developer",
					"job_salary":"200000",
					  "np":20,
					  "job_id":5,
					"budget": 200000,
					"description": "gRPC",
					"exp":2,
					"job_locations": ["bangalo"],
					"tech_stack": [
						"java"  ],
					"work_mode": "Remote",
					"qualification": [ "bcs-cse"
					],
					"shift": [ "morning-shift"],
					"job_type": [ "fulltime"]
				},
				 {"job_title": "Java Developer",
					"job_salary":"200000",
					  "np":20,
					  "job_id":6,
					"budget": 200000,
					"description": "gRPC",
					"exp":2,
					"job_locations": ["hyderaba"],
					"tech_stack": ["java" ],
					"work_mode": "Remote",
					"qualification": [ "bcs-cse"],
					"shift": [ "morning-shift"],
					"job_type": [ "fulltime"]}]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().ProcessingJobDetails(gomock.Any()).Return([]model.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.processingJobInput(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_postJobByCompany(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.postJobByCompany(tt.args.c)
		})
	}
}
