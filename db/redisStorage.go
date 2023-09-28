package db

import (
	"bootcam1_users/custom_errors"
	"bootcam1_users/structures"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type redisStorage struct {
	connection *redis.Client
}

func NewRedisStorage(connection *redis.Client) Storage {
	redis := redisStorage{connection: connection}

	return &redis
}
func (rs *redisStorage) Get(uuid uuid.UUID) (structures.User, error) {
	//read string
	val, errGet := rs.connection.Get(context.Background(), "user_"+uuid.String()).Result()
	if errGet != nil {
		return structures.User{}, custom_errors.Error_UserNotFound
	}
	return parseUser(val)
}
func (rs *redisStorage) GetAll() ([]structures.User, error) {
	// Create a context
	ctx := context.TODO()

	// Use the SCAN command to retrieve keys matching the pattern
	pattern := "user_*"
	// var cursor uint64
	keys := make([]string, 0)

	iter := rs.connection.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return []structures.User{}, errors.New("error consultando reultados")
	}

	//use the keys to retrieve the documents
	values, err := rs.connection.MGet(ctx, keys...).Result()
	if err != nil {
		return []structures.User{}, errors.New("error consultando reultados")

	}

	//parse from jsonString to User struct and return the users
	return parseUserArray(values)
}
func (rs *redisStorage) Create(user structures.User) (structures.User, error) {
	//check if user exists
	_, errGet := rs.Get(user.ID)
	if errGet == nil {
		return structures.User{}, custom_errors.Error_UuidAlreadyExists
	}
	//marshal to json string
	jsonVal, err := json.Marshal(user)
	if err != nil {
		return structures.User{}, err
	}

	//save in database
	err = rs.connection.Set(context.Background(), "user_"+user.ID.String(), jsonVal, 0).Err()
	if err != nil {
		panic(err)
	}

	//parse string to User
	return rs.Get(user.ID)
}
func (rs *redisStorage) Update(uuid uuid.UUID, newUser structures.User) (structures.User, error) {
	//check if user exists
	_, errGet := rs.Get(uuid)
	if errGet != nil {
		return structures.User{}, custom_errors.Error_UserNotFound
	}
	//marshal to json string
	jsonVal, err := json.Marshal(newUser)
	if err != nil {
		return structures.User{}, err
	}

	//save in database
	err = rs.connection.Set(context.Background(), "user_"+newUser.ID.String(), jsonVal, 0).Err()
	if err != nil {
		panic(err)
	}

	//parse string to User
	return rs.Get(newUser.ID)

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

func parseUserArray(usersArr []interface{}) ([]structures.User, error) {
	usersParsed := make([]structures.User, 0)
	for _, userToParse := range usersArr {
		//parse this user
		user, _ := parseUser(fmt.Sprint(userToParse))
		usersParsed = append(usersParsed, user)
	}
	return usersParsed, nil
}
func parseUser(jsonVal string) (structures.User, error) {
	data := structures.User{}
	err := json.Unmarshal([]byte(jsonVal), &data)
	if err != nil {
		return structures.User{}, custom_errors.Error_ParsingJson
	}
	return data, nil
}
