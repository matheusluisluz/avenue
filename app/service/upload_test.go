package service

import (
	"context"
	"mime/multipart"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"avenue/app/model"
	"avenue/app/repository"
)

type UploadServiceTestSuite struct {
	suite.Suite

	ctx        context.Context
	c          *gin.Context
	assert     *assert.Assertions
	repository *repository.MockUploadRepository
	file       multipart.File
	service    UploadService
}

func TestUploadServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UploadServiceTestSuite))
}

func (s *UploadServiceTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.assert = assert.New(s.T())
	s.repository = &repository.MockUploadRepository{}
	s.service = UploadService(
		s.service,
	)
}

func (s *UploadServiceTestSuite) TearDownTest() {
	s.repository.AssertExpectations(s.T())
}

func (s *UploadServiceTestSuite) TestUpload() {

	upload := &model.Upload{
		FileName: "file.txt",
	}

	s.repository.
		On("Upload", s.c, mock.AnythingOfType("*domain.Upload")).
		Return("123", nil)

	result, _ := s.service.Upload(s.c, upload)

	// s.assert.NoError(err)
	s.assert.Equal(result.Success, true)
}
