package repository

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/allegro/bigcache/v3"

	"avenue/app/model"
)

type IRepository interface {
	Upload(ctx context.Context, file *model.Upload) (string, error)
	Read(read *model.Chunk) ([]byte, error)
}

type UploadRepository struct {
	cache *bigcache.BigCache
}

func Execute(config bigcache.Config) *UploadRepository {
	cache, _ := bigcache.NewBigCache(config)
	return &UploadRepository{
		cache: cache,
	}
}

func (repository *UploadRepository) Upload(ctx context.Context, file *model.Upload) (string, error) {
	fmt.Println("file.ID: ", file.ID)
	fmt.Println("file: ", file.File)
	// repository.store.Store(file.ID, file)
	b, err := ioutil.ReadAll(file.File)
	if err != nil {
		panic(err)
	}
	repository.cache.Set(file.ID, b)

	return file.ID, nil
}

func (repository *UploadRepository) Read(read *model.Chunk) ([]byte, error) {
	fmt.Println("read.UploadID: ", read.UploadID)
	// file, ok := repository.store.Get(read.UploadID)
	// fmt.Println("file: ", file)
	entry, err := repository.cache.Get(read.UploadID)
	if err != nil {
		panic(err)
	}

	return entry, nil
}
