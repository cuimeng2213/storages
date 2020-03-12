package objects

import (
	"fmt"
	"lib/objectstream"
	"storages/lession03/apiServer/heartbeat"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	fmt.Printf("%s find server=[%s] \n", "putStream", server)
	if server == "" {
		return nil, fmt.Errorf("can not found any server")
	}
	return objectstream.NewPutStream(server, object), nil
}
