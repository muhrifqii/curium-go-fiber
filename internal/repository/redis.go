package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	StorageRedis struct {
		db *redis.Client
	}
)

func NewStorageRedis(rdb *redis.Client) *StorageRedis {
	return &StorageRedis{
		db: rdb,
	}
}

// modified from https://github.com/gofiber/storage/blob/main/redis/redis.go

// Get value by key
func (s *StorageRedis) Get(key string) ([]byte, error) {
	if len(key) <= 0 {
		return nil, nil
	}
	val, err := s.db.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return val, err
}

// Set key with value
func (s *StorageRedis) Set(key string, val []byte, exp time.Duration) error {
	if len(key) <= 0 || len(val) <= 0 {
		return nil
	}
	return s.db.Set(context.Background(), key, val, exp).Err()
}

// Delete key by key
func (s *StorageRedis) Delete(key string) error {
	if len(key) <= 0 {
		return nil
	}
	return s.db.Del(context.Background(), key).Err()
}

// Reset all keys
func (s *StorageRedis) Reset() error {
	return s.db.FlushDB(context.Background()).Err()
}

// Close the database
func (s *StorageRedis) Close() error {
	return s.db.Close()
}

// Return database client
func (s *StorageRedis) Conn() redis.UniversalClient {
	return s.db
}

// Return all the keys
func (s *StorageRedis) Keys() ([][]byte, error) {
	var keys [][]byte
	var cursor uint64
	var err error

	for {
		var batch []string

		if batch, cursor, err = s.db.Scan(context.Background(), cursor, "*", 10).Result(); err != nil {
			return nil, err
		}

		for _, key := range batch {
			keys = append(keys, []byte(key))
		}

		if cursor == 0 {
			break
		}
	}

	if len(keys) == 0 {
		return nil, nil
	}

	return keys, nil
}
