package main

import (
	"manager/celery"
)

func main() {
	celery.Client()
	celery.Worker()
}

