package heartbeat

import (
	//     "fmt"
	"lib/rabbitmq"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	dataServers = make(map[string]time.Time)
	mutex       sync.Mutex
)

func ListenHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()

	q.Bind("apiServers")

	c := q.Consume()
	go removeExpiredDataServer()

	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		//         fmt.Printf("find a server %s \n", dataServer)
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	//map的key是服务器名称
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	//     fmt.Printf("GetDataServers %v \n", ds)
	return ds
}
