package db

import "github.com/shomali11/xredis"

var opts = &xredis.Options{
	Host: "localhost",
	Port: 6379,
}

var client = xredis.SetupClient(opts)

// GetDb provide redis client
func GetDb() *xredis.Client {
	if client == nil {
		client = xredis.SetupClient(opts)
	}
	return client
}

// CloseDb close the connection with the client and redis
func CloseDb() {
	client.Close()
}
