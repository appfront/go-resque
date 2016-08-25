package resque

import (
	"github.com/appfront/go-resque"
	"github.com/appfront/go-resque/driver"
	"gopkg.in/redis.v3"
)

func init() {
	resque.Register("redisv2", &drv{})
}

type drv struct {
	client *redis.Client
	driver.Enqueuer
}

func (d *drv) SetClient(client interface{}) {
	d.client = client.(*redis.Client)
}

func (d *drv) ListPush(queue string, jobJSON string) (int64, error) {
	return d.client.RPush(queue, jobJSON).Result()
}
