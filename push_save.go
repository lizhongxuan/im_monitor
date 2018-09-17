package main

import (
	"net/http"
	"github.com/golang/glog"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

func Push_saveDo(path string) {

	fmt.Println("Push_saveDo - path:",path)
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

	push_all,ok := m["push_all"].(float64)
	if ok {
		tags := map[string]string{
			"v": "push_all",
		}

		fields := map[string]interface{}{
			"number": push_all,
		}

		writeInfluxdb("push",tags,fields)
	}

	push_success,ok := m["push_success"].(float64)
	if ok {
		tags := map[string]string{
			"v": "push_success",
		}

		fields := map[string]interface{}{
			"number": push_success,
		}

		writeInfluxdb("push",tags,fields)
	}

	push_fail_badtoken,ok := m["push_fail_badtoken"].(float64)
	if ok {
		tags := map[string]string{
			"v": "push_fail_badtoken",
		}

		fields := map[string]interface{}{
			"number": push_fail_badtoken,
		}

		writeInfluxdb("push",tags,fields)
	}

	push_fail_other,ok := m["push_fail_other"].(float64)
	if ok {
		tags := map[string]string{
			"v": "push_fail_other",
		}

		fields := map[string]interface{}{
			"number": push_fail_other,
		}

		writeInfluxdb("push",tags,fields)
	}

	save_all,ok := m["save_all"].(float64)
	if ok {
		tags := map[string]string{
			"v": "save_all",
		}

		fields := map[string]interface{}{
			"number": save_all,
		}

		writeInfluxdb("save",tags,fields)
	}

	save_success,ok := m["save_success"].(float64)
	if ok {
		tags := map[string]string{
			"v": "save_success",
		}

		fields := map[string]interface{}{
			"number": save_success,
		}

		writeInfluxdb("save",tags,fields)
	}

	save_fail,ok := m["save_fail"].(float64)
	if ok {
		tags := map[string]string{
			"v": "save_fail",
		}

		fields := map[string]interface{}{
			"number": save_fail,
		}

		writeInfluxdb("save",tags,fields)
	}

	handle_goroutine,ok := m["handle_goroutine"].(float64)
	if ok {
		tags := map[string]string{
			"v": "handle_goroutine",
		}

		fields := map[string]interface{}{
			"number": handle_goroutine,
		}

		writeInfluxdb("handle",tags,fields)
	}


}

