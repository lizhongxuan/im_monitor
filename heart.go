package main

import (
	"net/http"
	"fmt"
)


func HeartDo(name string, path string) {
	fmt.Println("HeartDo - path:",path)
	resp, err := http.Get(path)
	if err != nil || resp.StatusCode != 200  {
		monitorHeart.Lock()
		times,ok := monitorHeart.heartMap[name]
		if !ok {
			times = 0
		}
		times = times+1
		monitorHeart.heartMap[name] = times
		monitorHeart.Unlock()
		return
	}

	defer resp.Body.Close()
	fmt.Println(name, " HTTPDo success.")
}
