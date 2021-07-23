package repository

import (
	"avenue/app/model"

	"github.com/stretchr/testify/mock"
)

type MockUploadRepository struct {
	mock.Mock
}

func (m *MockUploadRepository) Read(read *model.Chunk) ([]byte, error) {
	args := m.Called(read)

	upload, _ := args.Get(0).([]byte)

	return upload, args.Error(1)
}

func (m *MockUploadRepository) Upload(file *model.Upload) (string, error) {
	args := m.Called(file)

	return args.String(0), args.Error(1)
}
