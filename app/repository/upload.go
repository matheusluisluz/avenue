package repository

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/allegro/bigcache/v3"

	"avenue/app/model"
)

var (
	result *model.Upload
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
	fmt.Println("file: ", file)
	fmt.Println("file.ID: ", file.ID)
	// fmt.Println("file.File: ", file.File)
	//TODO build a struct with bytes of file and store it
	// pointer := fmt.Sprintf("%v", file)
	// fmt.Println("file: ", pointer)
	// bytesFiles := []byte(file)
	b, err := ioutil.ReadAll(file.File)
	if err != nil {
		panic(err)
	}

	// b, _ := json.Marshal(&model.Upload{
	// 	ID:       file.ID,
	// 	FileName: file.FileName,
	// 	Path:     file.Path,
	// 	File:     file.File,
	// })
	repository.cache.Set(file.ID, b)

	return file.ID, nil
}

func (repository *UploadRepository) Read(read *model.Chunk) ([]byte, error) {
	fmt.Println("read.UploadID: ", read.UploadID)
	entry, err := repository.cache.Get(read.UploadID)
	if err != nil {
		panic(err)
	}
	// fmt.Println("entry: ", entry)
	// b, err := ioutil.ReadAll(file.File)
	// if err != nil {
	// 	panic(err)
	// }

	// json.Unmarshal(entry, &result)
	// fmt.Println("json.Unmarshal(entry, &languages): ", result)
	// handlerA := result.File
	// fmt.Println("result.File: ", handlerA)
	// handlerB := &result.File
	// fmt.Println("&result.File: ", handlerB)
	// handlerC := *result
	// fmt.Println("*result.File: ", handlerC.File)

	// message := entry
	// ioutil.WriteFile("teste", message, 0644)

	return entry, nil
}
