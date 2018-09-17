package main

import (
	"github.com/golang/glog"
	"github.com/influxdata/influxdb/client/v2"
	"time"

	"sync"
)

type Pbs struct {
	pbs   client.BatchPoints
	sync.RWMutex
}

var pbs_c *Pbs

func runInfluxdb(addr string, username string, password string) {

	for {
		if pbs_c != nil && pbs_c.pbs != nil{
			c, err := client.NewHTTPClient(client.HTTPConfig{
				Addr:     addr,
				Username: username,
				Password: password,
			})
			if err != nil {
				glog.Fatal(err)
			}
			pbs_c.Lock()
			// Write the batch
			if err := c.Write(pbs_c.pbs); err != nil {
				glog.Warning(err)
			}
			pbs_c.pbs = nil
			pbs_c.Unlock()
			// Close client resources
			if err := c.Close(); err != nil {
				glog.Warning(err)
			}


		} else {
			time.Sleep(5 * time.Second)
		}
		time.Sleep(1 * time.Second)
	}
}

func writeInfluxdb(measurement string, tags map[string]string, fields map[string]interface{}) {
	pbs_c.Lock()
	defer pbs_c.Unlock()

	if measurement == "" {
		glog.Warning("name = nil")
		return
	}

	t := time.Now()
	pt, err := client.NewPoint(measurement, tags, fields, t)
	if err != nil {
		glog.Warning(err)
		return
	}


	if pbs_c.pbs == nil {
		pbs2, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  "im_monitor",
			Precision: "s",
		})
		if err != nil {
			glog.Warning(err)
			return
		}
		pbs_c.pbs = pbs2
	}


	glog.Info("pbs :",pbs_c.pbs)
	glog.Info("pt:",pt)

	pbs_c.pbs.AddPoint(pt)
}
