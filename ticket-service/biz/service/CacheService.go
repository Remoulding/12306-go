package service

import (
	"context"
	"errors"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"github.com/bytedance/sonic"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"time"
)

type Loader = func(ctx context.Context) (interface{}, error)

func GetCache(ctx context.Context, cacheKey string) (string, error) {
	// 查询redis
	redisData, err := configs.Cache.Get(ctx, cacheKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.WithContext(ctx).Warnf("SafeGet failed, err: %v", err)
		return "", err
	}
	return redisData, nil
}

func GetHashCache(ctx context.Context, cacheKey string) (map[string]string, error) {
	// 查询redis
	redisData, err := configs.Cache.HGetAll(ctx, cacheKey).Result()
	if err != nil {
		log.WithContext(ctx).Warnf("SafeGet failed, err: %v", err)
		return nil, err
	}
	return redisData, nil
}

func SetHashCache(ctx context.Context, cacheKey string, data map[string]string) error {
	err := configs.Cache.HMSet(ctx, cacheKey, data).Err()
	if err != nil {
		log.WithContext(ctx).Errorf("SetCache failed, err: %v", err)
		return err
	}
	return nil
}

func SetCache(ctx context.Context, cacheKey string, data interface{}) error {
	resp, _ := sonic.MarshalString(data)
	err := configs.Cache.Set(ctx, cacheKey, resp, configs.CacheTTL).Err()
	if err != nil {
		log.WithContext(ctx).Errorf("SetCache failed, err: %v", err)
		return err
	}
	return nil
}

func SafeLoad(ctx context.Context, cacheKey, lockKey string, loader Loader) (string, error) {
	// 查询redis
	redisData, err := GetCache(ctx, cacheKey)
	if err != nil {
		log.WithContext(ctx).Errorf("SafeGet get cache failed, err: %v", err)
		return "", err
	}
	if len(redisData) > 0 {
		return redisData, nil
	}
	// 加锁
	if lockKey != "" {
		mutex := configs.Rs.NewMutex(lockKey, redsync.WithExpiry(5*time.Second))
		err = mutex.Lock()
		if err != nil {
			log.WithContext(ctx).Errorf("SafeGet add mutex failed, err: %v", err)
		}
		defer func(mutex *redsync.Mutex) {
			if _, err = mutex.Unlock(); err != nil {
				log.WithContext(ctx).Errorf("SafeGet unlock failed, err: %v", err)
			}
		}(mutex)
	}
	redisData, err = GetCache(ctx, cacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		log.WithContext(ctx).Errorf("SafeGet get cache, err: %v", err)
		return "", err
	}
	if len(redisData) > 0 {
		return redisData, nil
	}
	loadData, err := loader(ctx)
	if err != nil {
		log.WithContext(ctx).Errorf("SafeGet load data failed, err: %v", err)
		return "", err
	}
	resp, _ := sonic.MarshalString(loadData)
	err = configs.Cache.Set(ctx, cacheKey, resp, configs.CacheTTL).Err()
	if err != nil {
		log.WithContext(ctx).Errorf("SafeGet set cache, err: %v", err)
		return "", err
	}
	return resp, nil
}
