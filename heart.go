package main

import (
	"net/http"
	"time"
	"fmt"
)


func HeartDo(name string, path string) {
	fmt.Println("HeartDo - path:",path)
	resp, err := http.Get(path)

	if err != nil  {
		monitorHeart.Lock()
		times,ok := monitorHeart.heartMap[name]
		if !ok {
			times = 0
		}
		times = times+1
		monitorHeart.heartMap[name] = times
		monitorHeart.Unlock()


		tags := map[string]string{
			"v": "heart",
		}

		fields := map[string]interface{}{
			"times": times,
		}

		monitorCommitCh <- &influxdbNode{
			name: name,
			time: time.Now().Unix(),
			tags:tags,
			fields:fields,
		}

		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200  {
		monitorHeart.Lock()
		times,ok := monitorHeart.heartMap[name]
		if !ok {
			times = 0
		}
		times = times+1
		monitorHeart.heartMap[name] = times
		monitorHeart.Unlock()


		tags := map[string]string{
			"v": "heart",
		}

		fields := map[string]interface{}{
			"times": times,
		}

		monitorCommitCh <- &influxdbNode{
			name: name,
			time: time.Now().Unix(),
			tags:tags,
			fields:fields,
		}

		return
	}
	fmt.Println(name, " HTTPDo success.")
}
