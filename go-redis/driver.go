package resque

import (
	"github.com/appfront/go-resque"
	"github.com/appfront/go-resque/driver"
	"github.com/fiorix/go-redis/redis"
)

func init() {
	resque.Register("redis-go", &drv{})
}

type drv struct {
	client *redis.Client
	driver.Enqueuer
}

func (d *drv) SetClient(client interface{}) {
	d.client = client.(*redis.Client)
}

func (d *drv) ListPush(queue string, jobJSON string) (int64, error) {
	listLength, err := d.client.RPush(queue, jobJSON)
	if err != nil {
		return -1, err
	}

	return int64(listLength), err
}
