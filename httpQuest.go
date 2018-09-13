package main

import (
	"fmt"
	"time"
	"github.com/spf13/viper"
)

var monitorCommitCh chan *influxdbNode

type influxdbNode struct {
	name string
	tags map[string]string
	fields map[string]interface{}
	time int64
}


func monitorQuest() {
	monitorCommitCh = make(chan *influxdbNode,200)
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
	for node := range monitorCommitCh {
		fmt.Println("accumulater:", node.name, "  time:", node.time)
		writeInfluxdb(node.name, node.tags, node.fields)

	}
}
