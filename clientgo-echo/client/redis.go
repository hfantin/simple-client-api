package client

import (
	"fmt"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

type RedisClient struct {
	pool *redis.Pool
}

func (c *RedisClient) New() {
	log.Println("initializing redis")
	c.pool = initPool()
}

func initPool() *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
	return pool
}

func (c *RedisClient) Set(key string, value []byte) error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, []byte(value))
	if err != nil {
		log.Println("erro setar variavel: ", err)
	}
	conn.Do("EXPIRE", key, 60) //10 Minutes
	return err
}

func (c *RedisClient) Get(key string) ([]byte, error) {
	conn := c.pool.Get()
	defer conn.Close()
	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}
