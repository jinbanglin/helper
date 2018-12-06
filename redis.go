package helper

import (
  "os"
  "time"

  "github.com/go-redis/redis"
  "github.com/spf13/viper"
  "strconv"
)

//you have to config your toml file like this:

//[redis]
//addr                  =["127.0.0.1:6379"]
//heartbeatfrequency    =500
//db                    =0
//password              =""
//maxretries            =3
//dialtimeout           =3
//readtimeout           =3
//writetimeout          =3
//poolsize              =4096
//pooltimeout           =30
//idletimeout           =500
//idlecheckfrequency    =500

var GRedisRing *redis.Ring

func RedisChaos() {
  addr := make(map[string]string)
  for k, v := range viper.GetStringSlice("redis.addr") {
    addr["shard"+strconv.Itoa(k+1)] = v
  }
  GRedisRing = redis.NewRing(&redis.RingOptions{
    Addrs:              addr,
    HeartbeatFrequency: time.Duration(viper.GetInt("redis.heartbeatfrequency")) * time.Millisecond,
    OnConnect:          nil,
    DB:                 viper.GetInt("redis.db"),
    Password:           viper.GetString("redis.password"),
    MaxRetries:         viper.GetInt("redis.maxretries"),
    DialTimeout:        time.Duration(viper.GetInt("redis.dialtimeout")) * time.Second,
    ReadTimeout:        time.Duration(viper.GetInt("redis.readtimeout")) * time.Second,
    WriteTimeout:       time.Duration(viper.GetInt("redis.writetimeout")) * time.Second,
    PoolSize:           viper.GetInt("redis.poolsize"),
    PoolTimeout:        time.Duration(viper.GetInt("redis.pooltimeout")) * time.Second,
    IdleTimeout:        time.Duration(viper.GetInt("redis.idletimeout")) * time.Millisecond,
    IdleCheckFrequency: time.Duration(viper.GetInt("redis.idlecheckfrequency")) * time.Millisecond,
  })
  if err := GRedisRing.Ping().Err(); err != nil {
    panic(err)
    os.Exit(1)
  }
}
