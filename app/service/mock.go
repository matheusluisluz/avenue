package service

import (
	"avenue/app/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockUploadService struct {
	mock.Mock
}

func (m *MockUploadService) UploadTest(c *gin.Context, upload *model.Upload) (*model.UploadResponse, error) {
	args := m.Called(c, upload)
	result := args.Get(0).(*model.UploadResponse)
	return result, args.Error(1)
}

func (m *MockUploadService) Upload(c *gin.Context, upload *model.Upload) (*model.UploadResponse, error) {
	args := m.Called(c, upload)
	result := args.Get(0).(*model.UploadResponse)
	return result, args.Error(1)
}
