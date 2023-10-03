package db

import (
	"bootcam1_users/custom_errors"
	"bootcam1_users/structures"
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type redisStorage struct {
	connection *redis.Client
	prefix     string
}

func NewRedisStorage(connection *redis.Client) Storage {
	redis := redisStorage{
		connection: connection,
		prefix:     "user_",
	}

	return &redis
}

func (rs *redisStorage) Get(id uuid.UUID) (interface{}, error) {
	val, errGet := rs.connection.Get(context.Background(), rs.prefix+id.String()).Result()
	if errGet != nil {
		return structures.User{}, custom_errors.Error_UserNotFound
	}

	return val, nil
}

func (rs *redisStorage) GetAll() ([]interface{}, error) {
	// Create a context
	ctx := context.TODO()

	// Use the SCAN command to retrieve keys matching the pattern
	pattern := rs.prefix + "*"
	// var cursor uint64
	keys := make([]string, 0)

	iter := rs.connection.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	//use the keys to retrieve the documents
	values, err := rs.connection.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, errors.New("error consultando reultados 2")

	}

	//parse from jsonString to User struct and return the users
	return values, nil
}

func (rs *redisStorage) Create(id uuid.UUID, entity interface{}) (interface{}, error) {
	//check if user exists
	_, errGet := rs.Get(id)
	if errGet == nil {
		return structures.User{}, custom_errors.Error_UuidAlreadyExists
	}
	//marshal to json string
	jsonVal, err := json.Marshal(entity)
	if err != nil {
		return structures.User{}, err
	}

	//save in database
	err = rs.connection.Set(context.Background(), "user_"+id.String(), jsonVal, 0).Err()
	if err != nil {
		panic(err)
	}

	//parse string to User
	return rs.Get(id)
}

func (rs *redisStorage) Update(id uuid.UUID, entity interface{}) (interface{}, error) {
	//check if user exists
	_, errGet := rs.Get(id)
	if errGet != nil {
		return structures.User{}, custom_errors.Error_UserNotFound
	}
	//marshal to json string
	jsonVal, err := json.Marshal(entity)
	if err != nil {
		return structures.User{}, err
	}

	//save in database
	err = rs.connection.Set(context.Background(), "user_"+id.String(), jsonVal, 0).Err()
	if err != nil {
		panic(err)
	}

	//parse string to User
	return rs.Get(id)
}

func (rs *redisStorage) Delete(id uuid.UUID) error {
	ctx := context.TODO()
	_, err := rs.Get(id)

	if err != nil {
		return custom_errors.Error_UserNotFound
	}
	_, errDeleting := rs.connection.Del(ctx, "user_"+id.String()).Result()
	return errDeleting
}

// // /////
