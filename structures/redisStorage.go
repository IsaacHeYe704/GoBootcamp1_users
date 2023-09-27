package structures

import (
	"bootcam1_users/custom_errors"
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
func (rs *redisStorage) Get(uuid uuid.UUID) (User, error) {
	//read string
	val, errGet := rs.connection.Get(context.Background(), "user_"+uuid.String()).Result()
	if errGet != nil {
		return User{}, custom_errors.Error_UserNotFound
	}
	return parseUser(val)
}
func (rs *redisStorage) GetAll() ([]User, error) {
	// Create a context
	ctx := context.TODO()

	// Use the SCAN command to retrieve keys matching the pattern
	pattern := "user_*"
	var cursor uint64
	keys := make([]string, 0)
	for {
		var scanKeys []string
		var err error

		// Execute the SCAN command
		scanKeys, cursor, err = rs.connection.Scan(ctx, cursor, pattern, 10).Result()
		if err != nil {
			panic(err)
		}

		keys = append(keys, scanKeys...)

		// Check if the cursor is 0, which means already iterated through all keys
		if cursor == 0 {
			break
		}
	}

	//use the keys to retrieve the documents
	values, err := rs.connection.MGet(ctx, keys...).Result()
	if err != nil {
		return []User{}, errors.New("error consultando reultados")

	}

	//parse from jsonString to User struct and return the users
	return parseUserArray(values)
}
func (rs *redisStorage) Create(user User) (User, error) {
	//check if user exists
	_, errGet := rs.Get(user.ID)
	if errGet == nil {
		return User{}, custom_errors.Error_UuidAlreadyExists
	}
	//marshal to json string
	jsonVal, err := json.Marshal(user)
	if err != nil {
		return User{}, err
	}

	//save in database
	err = rs.connection.Set(context.Background(), "user_"+user.ID.String(), jsonVal, 0).Err()
	if err != nil {
		panic(err)
	}

	//parse string to User
	return rs.Get(user.ID)
}
func (rs *redisStorage) Update(uuid uuid.UUID, newUser User) (User, error) {
	//check if user exists
	_, errGet := rs.Get(uuid)
	if errGet != nil {
		return User{}, custom_errors.Error_UserNotFound
	}
	//marshal to json string
	jsonVal, err := json.Marshal(newUser)
	if err != nil {
		return User{}, err
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

func parseUserArray(usersArr []interface{}) ([]User, error) {
	usersParsed := make([]User, 0)
	for _, userToParse := range usersArr {
		//parse this user
		user, _ := parseUser(fmt.Sprint(userToParse))
		usersParsed = append(usersParsed, user)
	}
	return usersParsed, nil
}
func parseUser(jsonVal string) (User, error) {
	data := User{}
	err := json.Unmarshal([]byte(jsonVal), &data)
	if err != nil {
		return User{}, custom_errors.Error_ParsingJson
	}
	return data, nil
}
