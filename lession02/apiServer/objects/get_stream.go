package objects

import (
	"fmt"
	"io"
	"lib/objectstream"
	"storages/lession02/apiServer/locate"
)

//获取读文件句柄
func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("locate %s fail", object)
	}

	return objectstream.NewGetStream(server, object)

}
