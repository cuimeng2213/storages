package objects

import (
	"fmt"
	"lib/objectstream"
	"storages/lession02/apiServer/heartbeat"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("can not found any server")
	}
	return objectstream.NewPutStream(server, object), nil
}
