package heartbeat

import (
	"fmt"
	"math/rand"
)

func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	fmt.Printf("get %d servers %v \n", n, ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
