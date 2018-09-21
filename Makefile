all:im_monitor

im_monitor:im_monitor.go heart.go httpQuest.go influxdb_data.go push_save.go im.go
	go build im_monitor.go heart.go httpQuest.go influxdb_data.go push_save.go im.go
