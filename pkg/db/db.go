package db

import (
	"github.com/shomali11/xredis"
	"reflect"
	"fmt"
	"github.com/koinkoin-io/koinkoin.backend/third_party"
	"github.com/koinkoin-io/koinkoin.backend/pkg/util"
)

var opts = &xredis.Options{
	Host: util.GetEnvOrDefault("db_host", "localhost"),
	Port: util.GetEnvOrDefaultInt("db_port", 6379),
}

var client = xredis.SetupClient(opts)

// GetDb provide redis client
func GetDb() *xredis.Client {
	if client == nil {
		client = xredis.SetupClient(opts)
	}
	return client
}

// SaveStructure save fields with tag json into db.
func SaveStructure(key string, va interface{}) {
	t := reflect.TypeOf(va)
	v := reflect.ValueOf(va)
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			client.HSet(key, t.Field(i).Tag.Get("json"), fmt.Sprint(v.Field(i)))
		}
	}
}

// StructFromKey retrieve fields from db and inject it into
// val. Fields need to be tagged with 'json'.
func StructFromKey(key string, val interface{}) {
	keys, err := client.HKeys(key)
	t := reflect.TypeOf(val).Elem()
	v := reflect.ValueOf(val)

	if err != nil {
		return
	}

	mapStruct := structToMapTags(t, v)
	for _, tag := range keys {

		toSet, _, _ := client.HGet(key, tag)
		field, ok := mapStruct[tag]

		if !ok {
			continue
		}
		third.SetFieldTo(field, toSet)
	}
}

// structToMapTags convert a struct to a map with JSON Key-> Value.
// Because just one read, no loop thought all fields and check if the tag match.
func structToMapTags(t reflect.Type, v reflect.Value) map[string]reflect.Value {
	m := make(map[string]reflect.Value)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Tag.Get("json")] = v.Elem().Field(i)
	}
	return m
}

// CloseDb close the connection with the client and redis
func CloseDb() {
	client.Close()
}
