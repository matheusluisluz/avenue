package repository

import (
	"avenue/app/model"
	"bytes"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UploadRepositoryTestSuite struct {
	suite.Suite

	assert     *assert.Assertions
	repository IRepository
}

func TestUploadRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UploadRepositoryTestSuite))
}

func (s *UploadRepositoryTestSuite) SetupTest() {
	cacheConfig := bigcache.DefaultConfig(10 * time.Minute)
	s.assert = assert.New(s.T())
	s.repository = Execute(cacheConfig)
}

func (s *UploadRepositoryTestSuite) TestCreateAndGet() {
	buffer := new(bytes.Buffer)
	upload := &model.Upload{
		FileName: "nasdaq_symbols.csv",
		File:     buffer,
	}

	id, err := s.repository.Upload(upload)

	s.assert.NoError(err)

	read := &model.Chunk{
		UploadID: id,
		Offset:   123,
		Limit:    123,
	}
	result, err := s.repository.Read(read)

	s.assert.Equal(upload.FileName, "nasdaq_symbols.csv")
	s.assert.NotNil(result)
}

func (s *UploadRepositoryTestSuite) TestReadWithFailure() {
	read := &model.Chunk{
		UploadID: "123",
	}
	_, err := s.repository.Read(read)

	s.assert.Error(err)
}
