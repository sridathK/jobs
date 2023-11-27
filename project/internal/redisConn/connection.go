package redisconn

import (
	"context"
	"encoding/json"
	"project/internal/model"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RDBLayer struct {
	rdb *redis.Client
}

//go:generate mockgen -source=connection.go -destination=connection_mock.go -package=redisconn
type Caching interface {
	AddToTheCache(ctx context.Context, jid uint, jobData model.Job) error
	GetTheCacheData(ctx context.Context, jid uint) (string, error)
	AddToTheCacheOTP(ctx context.Context, email string, otp int) error
	GetTheCacheOTP(ctx context.Context, email string) (string, error)
}

func NewRDBLayer(rdb *redis.Client) *RDBLayer {
	return &RDBLayer{
		rdb: rdb,
	}
}

func (c *RDBLayer) AddToTheCache(ctx context.Context, jid uint, jobData model.Job) error {
	jobID := strconv.FormatUint(uint64(jid), 10)
	val, err := json.Marshal(jobData)
	if err != nil {
		return err
	}
	err = c.rdb.Set(ctx, jobID, val, 1*time.Minute).Err()
	return err
}

func (c *RDBLayer) AddToTheCacheOTP(ctx context.Context, email string, otp int) error {

	myString := strconv.FormatInt(int64(otp), 10)
	// if err != nil {
	// 	return err
	// }
	err := c.rdb.Set(ctx, email, myString, 10*time.Minute).Err()

	return err
}

func (c *RDBLayer) GetTheCacheOTP(ctx context.Context, email string) (string, error) {
	str, err := c.rdb.Get(ctx, email).Result()

	if err != nil {
		return "", err
	}
	return str, nil
}

func (c *RDBLayer) GetTheCacheData(ctx context.Context, jid uint) (string, error) {
	jobID := strconv.FormatUint(uint64(jid), 10)
	str, err := c.rdb.Get(ctx, jobID).Result()
	return str, err
}
