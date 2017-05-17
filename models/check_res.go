package models

import (
	"time"
)

type Check_Res struct {
	Uuid     string
	Ip       string          `bson:"ip"`
	Hostname string          `bson:"hostname"`
	Udp_res  map[string]bool `bson:"udp_res"`
	Tcp_res  map[string]bool `bson:"tcp_res"`
	//Http_res map[string]bool `bson:"http_res"`
	Time time.Time
}
