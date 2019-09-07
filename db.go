package helper

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ChaosDB() {
	MongodbInstance()
	RedisInstance()
}

var GMongo *mongo.Database

func MongodbInstance() {
	clientOptions := &options.ClientOptions{
		Hosts: []string{viper.GetString("mongodb.addr")},
	}
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	//defer client.Disconnect(context.TODO())
	GMongo = client.Database(viper.GetString("mongodb.database"))
}

var GRedis *redis.Client

func RedisInstance() {
	GRedis = redis.NewClient(&redis.Options{
		Addr:               viper.GetString("redis.addr"),
		OnConnect:          func(cn *redis.Conn) error { return cn.ClientSetName("on_connect").Err() },
		DB:                 viper.GetInt("redis.db"),
		Password:           viper.GetString("redis.password"),
		MaxRetries:         viper.GetInt("redis.maxretries"),
		DialTimeout:        time.Duration(viper.GetInt("redis.dialtimeout")) * time.Second,
		ReadTimeout:        time.Duration(viper.GetInt("redis.readtimeout")) * time.Second,
		WriteTimeout:       time.Duration(viper.GetInt("redis.writetimeout")) * time.Second,
		PoolSize:           viper.GetInt("redis.poolsize"),
		PoolTimeout:        time.Duration(viper.GetInt("redis.pooltimeout")) * time.Second,
		IdleTimeout:        time.Duration(viper.GetInt("redis.idletimeout")) * time.Minute,
		IdleCheckFrequency: time.Duration(viper.GetInt("redis.idlecheckfrequency")) * time.Millisecond,
	})
	if err := GRedis.Ping().Err(); err != nil {
		panic(err)
	}
}
