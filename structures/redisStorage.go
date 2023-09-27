package structures

import (
	"bootcam1_users/custom_errors"
	"context"
	"log/slog"

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

	return User{}, nil
}
func (rs *redisStorage) GetAll() (map[uuid.UUID]User, error) {
	return make(map[uuid.UUID]User), nil
}
func (rs *redisStorage) Create(user User) (User, error) {
	err := rs.connection.Set(context.Background(), user.ID.String(), User{}, 0)
	if err != nil {
		slog.Error(err.Err().Error())
		return User{}, custom_errors.Error_UserNotFound
	}

	return User{}, nil
}
func (rs *redisStorage) Update(uuid.UUID, User) (User, error) {
	return User{}, nil

}
func (rs *redisStorage) Delete(uuid.UUID) error {
	return nil
}

// err := rdb.Set(ctx, "key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := rdb.Get(ctx, "key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := rdb.Get(ctx, "key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
