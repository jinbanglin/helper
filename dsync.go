package helper

import (
  "time"
  "github.com/bsm/redis-lock"
  "github.com/go-redis/redis"
  "github.com/spf13/viper"
)

var gRdsClient *redis.Client

func NewDLock(topic string, timeOut time.Duration) *lock.Locker {
  if gRdsClient == nil {
    gRdsClient = redis.NewClient(&redis.Options{
      Network: "tcp",
      Addr:    viper.GetStringSlice("redis.addr")[0],
    })
  }

  return lock.New(gRdsClient, topic, &lock.Options{LockTimeout: timeOut, RetryCount: 10})
}
