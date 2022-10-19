package celery

import (
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
	"manager/send"
)

func Client() {
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
	cli, _ := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		1,
	)

	arr := []send.User{
		{TaskName: "send", Name: "Баирто", Surname: "Цыренов", Email: "bairto.c@gmail.com", Birthday: "15.05.1999", Link: "https://webhook.site/901da260-aed3-4251-80f2-40d29e100381"},
		{TaskName: "send", Name: "Иван", Surname: "Иванов", Email: "bairto.c@gmail.com", Birthday: "15.05.1999", Link: "https://webhook.site/901da260-aed3-4251-80f2-40d29e100381"},
	}

	for _, v := range arr {
		cli.Delay(
			v.TaskName,
			v.Name,
			v.Surname,
			v.Email,
			v.Birthday,
			v.Link,
		)
	}
}
