package service

import (
	"bytes"
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"avenue/app/model"
	"avenue/app/repository"
)

type GinMock struct {
	mock.Mock
}

type UploadServiceTestSuite struct {
	suite.Suite

	ctx        context.Context
	c          *gin.Context
	r          GinMock
	assert     *assert.Assertions
	repository *repository.MockUploadRepository
	service    IService
}

func TestUploadServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UploadServiceTestSuite))
}

func (s *UploadServiceTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.assert = assert.New(s.T())
	s.c, _ = gin.CreateTestContext(httptest.NewRecorder())
	s.r = GinMock{}
	s.repository = &repository.MockUploadRepository{}
	s.service = Execute(
		s.repository,
	)
}

func (s *UploadServiceTestSuite) TearDownTest() {
	s.repository.AssertExpectations(s.T())
}

func (s *UploadServiceTestSuite) TestUploadMem() {
	buffer := new(bytes.Buffer)
	upload := &model.Upload{
		FileName: "file.cvs",
		File:     buffer,
	}

	s.repository.
		On("Upload", mock.AnythingOfType("*model.Upload")).
		Return("asd", nil)

	result, err := s.service.UploadMem(s.c, upload)
	fmt.Println("result: ", result)
	s.assert.NoError(err)
	s.assert.Equal(result.Success, true)
}
