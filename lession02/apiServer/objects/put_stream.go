package objects

import (
	"fmt"
	"lib/objectStream"
	"storages/lession02/apiServer/heartbeat"
)

func putStream(object string) (*objectStream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("can not found any server")
	}
	return objectStream.NewPutStream(server, object), nil
}
