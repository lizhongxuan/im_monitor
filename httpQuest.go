package main

import (
	"fmt"
	"time"
	"github.com/spf13/viper"
)

//var monitorCommitCh chan *influxdbNode

//type influxdbNode struct {
//	name string
//	tags map[string]string
//	fields map[string]interface{}
//	time int64
//}

func monitorQuest() {
	go accumulater()
	for {
		select {
		case <-time.After(time.Second * 10):
			for k, v := range moni {
				fmt.Println("k:", k)
				fmt.Println("v:", v)
				go HeartDo(k, v)
			}
			go Push_saveDo(viper.GetString("im_handle"))
		}
	}
}


func accumulater() {
	for {
		select {
		case <-time.After(time.Second * 5):
			for k, times := range monitorHeart.heartMap {

				tags := map[string]string{
					"v": "heart",
				}

				fields := map[string]interface{}{
					"times": times,
				}
				writeInfluxdb(k,tags, fields)
			}
		}
	}

}
