package redis

import (
	"../../url-shortener/shortener"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"strconv"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisClient(redisURL string, redisPassword string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPassword,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepository(redisURL string, redisPassword string) (shortener.RedirectRepository, error) {
	repository := &redisRepository{}
	client, err := newRedisClient(redisURL, redisPassword)
	if err != nil {
		return nil, errors.Wrap(err, "Error in Redis repository.NewRedisRepository")
	}
	repository.client = client
	return repository, nil
}

func (r *redisRepository) generateKey(code string) string {
	return fmt.Sprintf("redirect:%s", code)
}

func (r redisRepository) Find(Code string) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	key := r.generateKey(Code)
	data, err := r.client.HGetAll(key).Result()

	if err != nil {
		return nil, errors.Wrap(err, "Error in Redis Find repository.Redirect.Find")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(shortener.ErrorRedirectNotFound, "repository.Redirect.Find")
	}
	createdAt, err := strconv.ParseInt(data["created_at"], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	redirect.Code = data["code"]
	redirect.URL = data["url"]
	redirect.CreatedAt = createdAt
	return redirect, nil
}

func (r redisRepository) Store(redirect *shortener.Redirect) error {
	key := r.generateKey(redirect.Code)
	data := map[string]interface{}{
		"code":       redirect.Code,
		"url":        redirect.URL,
		"created_at": redirect.CreatedAt,
	}
	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}
