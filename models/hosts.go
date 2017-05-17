package models

import (
	u "code.google.com/p/go-uuid/uuid"
	"strings"
)

type Host struct {
	Uuid string
	// Http_list map[string]string `bson:"http_list"`
	Tcp_list []string `bson:"tcp_list"`
	Udp_list []string `bson:"udp_list"`
	Ip       string   `bson:"ip"`
	Hostname string   `bson:"hostname`
}

type Host_From_Web struct {
	Tcp_list string
	Udp_list string
	Ip       string
	Hostname string
}

/**
 * Gen uuid for new host
 * @param  {none} this *Host)        Gen_Id
 * @return {new host witch have uuid}
 */
func (this *Host) Gen_Id() {
	this.Uuid = u.New()
}

/**
 * Chang string of tcp\udp into list
 * @param  {host from web post} this *Host_From_Web) Change_To_Host() (host Host )
 * @return {Host}
 */
func (this *Host_From_Web) Change_To_Host() (host Host) {
	host.Hostname = this.Hostname
	host.Ip = this.Ip
	host.Tcp_list = strings.Split(this.Tcp_list, ",")
	host.Udp_list = strings.Split(this.Udp_list, ",")
	return
}
