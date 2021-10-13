package skyhookcomp

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"
)

type Record struct {
	ID    string
	Name  string
	Value string
}

type Store struct {
	cache  *cache.Cache
	client redis.UniversalClient
}

func NewStore(target string) *Store {
	client := redis.NewUniversalClient(
		&redis.UniversalOptions{
			Addrs: []string{target},
		},
	)
	cache := cache.New(
		&cache.Options{
			Redis: client,
			Marshal: func(v interface{}) ([]byte, error) {
				b, err := msgpack.Marshal(v)
				if err != nil {
					log.Printf("ERROR: writing: %q | %v | error: %v", b, b, err)
				} else {
					log.Printf("DEBUG: writing: %q | %v", b, b)
				}
				return b, err
			},
			Unmarshal: func(b []byte, v interface{}) error {
				err := msgpack.Unmarshal(b, v)
				if err != nil {
					log.Printf("ERROR: reading: %q | %v | error: %v", b, b, err)
				}
				return err
			},
		},
	)
	return &Store{
		client: client,
		cache:  cache,
	}
}

func (s *Store) Get(ctx context.Context, id string) (*Record, error) {
	r := &Record{}
	k := fmt.Sprintf("record:%s", id)
	err := s.cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   k,
		Value: r,
		TTL:   time.Minute,
		Do: func(i *cache.Item) (interface{}, error) {
			return &Record{
				ID:    i.Key,
				Name:  fmt.Sprintf("name:%s", i.Key),
				Value: fmt.Sprintf("value:%s ts:%v", i.Key, time.Now()),
			}, nil
		},
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}
