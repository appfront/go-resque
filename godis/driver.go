package resque

import (
	"github.com/appfront/go-resque"
	"github.com/appfront/go-resque/driver"
	"github.com/simonz05/godis/redis"
)

func init() {
	resque.Register("godis", &drv{})
}

type drv struct {
	client *redis.Client
	driver.Enqueuer
}

func (d *drv) SetClient(client interface{}) {
	d.client = client.(*redis.Client)
}

func (d *drv) ListPush(queue string, jobJSON string) (int64, error) {
	return d.client.Rpush(queue, jobJSON)
}
