package server

import (
	cfg "mgoblin/users_grpc/internal/users/server/config"
	"reflect"
	"testing"
)

func TestServer(t *testing.T) {
	server := New(cfg.New())

	want := make([]string, 0)
	want = append(want, "v1.UserService")
	info := server.GetServiceInfo()
	keys := make([]string, 0, len(info))
	for k := range info {
		keys = append(keys, k)
	}

	eq := reflect.DeepEqual(keys, want)
	if !eq {
		t.Errorf("invalid registered services.actual %+v, want %+v", keys, want)
	}
}
