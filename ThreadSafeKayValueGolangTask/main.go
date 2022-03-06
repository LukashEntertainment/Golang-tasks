package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func main(){
	// Store.kvmap.Put("1", 1)
	// fmt.Println(Store.kvmap.Get("1"))
}

//ошибка компиляции
type KVStorage interface {
	Get(context.Context, key string) (interface{}, error)
	Put(context.Context, key string, val interface{}) error
	Delete(context.Context, key string) error
}

type Store struct {
	kvmap sync.Map
}

func (s *Store) Put(context.Context, key string, val interface{}) error {
	s.kvmap.Store(key, val)
}

func (s *Store) Get(context.Context, key string) (interface{}, error) {
	var result = s.kvmap.Load(key)
	if(result != nil){
		return result, nil
	}
	return result, errors.New("Key not found")
}

func (s *Store) Delete(context.Context, key string) error {
	if (s.Get(key) != nil){
		s.kvmap.Delete(key)
		return nil
	}
	else{
		return errors.New("Key not found")
	}
}

