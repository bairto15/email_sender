package celery

import (
	"log"
	"manager/send"
	"time"

	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

func Worker() {
	// create redis connection pool
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://localhost:6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	// initialize celery client
	cli, err := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		5,
	)
	if err != nil {
		log.Print(err)
	}

	// task
	task := func(name string, surname string, email string, birthday string, link string) int {
		user := send.User{Name: name, Surname: surname, Email: email, Birthday: birthday, Link: link}
		send.Send(user)
		return 0
	}

	cli.Register("send", task)

	cli.StartWorker()

	time.Sleep(10 * time.Second)

	cli.StopWorker()
}
