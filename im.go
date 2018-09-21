package main

import (
	"fmt"
	"net/http"
	"github.com/golang/glog"
	"io/ioutil"
	"encoding/json"
)

func IMDo(path string) {

	fmt.Println("IMDo - path:",path)
	resp, err := http.Get(path)
	if err != nil  {
		glog.Warning(err)
		return
	}
	defer resp.Body.Close()

	b,err := ioutil.ReadAll(resp.Body)
	var m map[string]interface{}
	json.Unmarshal(b,&m)
	fmt.Println("m:",m)

	goroutine_count,ok := m["goroutine_count"].(float64)
	if ok {
		tags := map[string]string{
			"v": "goroutine_count",
		}

		fields := map[string]interface{}{
			"number": goroutine_count,
		}

		writeInfluxdb("im",tags,fields)
	}

	connection_count,ok := m["connection_count"].(float64)
	if ok {
		tags := map[string]string{
			"v": "connection_count",
		}

		fields := map[string]interface{}{
			"number": connection_count,
		}

		writeInfluxdb("im",tags,fields)
	}

	client_count,ok := m["client_count"].(float64)
	if ok {
		tags := map[string]string{
			"v": "client_count",
		}

		fields := map[string]interface{}{
			"number": client_count,
		}

		writeInfluxdb("im",tags,fields)
	}

	in_message_count,ok := m["in_message_count"].(float64)
	if ok {
		tags := map[string]string{
			"v": "in_message_count",
		}

		fields := map[string]interface{}{
			"number": in_message_count,
		}

		writeInfluxdb("im",tags,fields)
	}

	out_message_count,ok := m["out_message_count"].(float64)
	if ok {
		tags := map[string]string{
			"v": "out_message_count",
		}

		fields := map[string]interface{}{
			"number": out_message_count,
		}

		writeInfluxdb("im",tags,fields)
	}



}
