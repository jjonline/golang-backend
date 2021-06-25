package initializer

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jjonline/go-mod-library/logger"
	"github.com/jjonline/golang-backend/config"
	"runtime"
	"time"
)

func initRedis() *redis.Client {
	cnf := config.Config.Redis
	r := redis.NewClient(&redis.Options{
		Network:            "tcp",
		Addr:               fmt.Sprintf("%s:%d", cnf.Host, cnf.Port),
		Password:           cnf.Password,
		DB:                 cnf.Database,
		DialTimeout:        time.Duration(cnf.ConnectTimeout) * time.Second,
		ReadTimeout:        time.Duration(cnf.ReadTimeout) * time.Second,
		WriteTimeout:       time.Duration(cnf.WriteTimeout) * time.Second,
		PoolSize:           cnf.PoolMaxOpen * runtime.NumCPU(),
		MinIdleConns:       cnf.PoolMaxIdle,
		IdleTimeout:        time.Duration(cnf.PoolMaxTime) * time.Second,
		IdleCheckFrequency: 10 * time.Minute,
	})
	r.AddHook(&logger.RedisHook{})
	return r
}