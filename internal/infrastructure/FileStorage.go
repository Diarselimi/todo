package infrastructure

import (
	"cards/internal/domain"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type IStorage interface {
	GetList() (domain.TodoSlice, error)
	Save(list domain.TodoSlice)
}

type FileStorage struct {
	path      string
	separator string
}

func NewStorage() FileStorage {
	return FileStorage{
		path:      "testdata/todos",
		separator: "\n",
	}
}

func (s FileStorage) GetList() (domain.TodoSlice, error) {
	bs, err := os.ReadFile(s.path)
	if err != nil {
		fmt.Println("File could not be read!")
		return nil, errors.New("File could not be read.")
	}

	var result domain.TodoSlice
	err = json.Unmarshal(bs, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s FileStorage) Save(list domain.TodoSlice) {

	os.WriteFile(s.path, []byte(toJson(list)), 0666)
	fmt.Println("Finish", toJson(list))
}

func toJson(item any) string {
	res, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}
