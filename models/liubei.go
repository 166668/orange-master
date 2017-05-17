package models

import (
	"log"
	//"log"
	//"gopkg.in/mgo.v2"
	//"log"
	//"time"
	//"runtime"
	"sync"
)

type LiuBei struct {
	//Task     chan Host,100
	//Res_Task chan Check_Res
	ZhangFei
	ZhuGeliang
}

func (this *LiuBei) Do_Port_Check() {
	var wg sync.WaitGroup
	wg.Add(3)

	host_chan := make(chan Host, 100)
	res_chan := make(chan Check_Res, 100)

	go func() {
		wg.Done()
		for _, host := range this.Query_Host_For_Api() {
			//log.Println(host)
			host_chan <- host
		}
		close(host_chan)
	}()

	go func() {
		wg.Done()
		for {
			host, ok := <-host_chan
			if !ok {
				close(res_chan)
				break
			}
			log.Println(host)
			res_chan <- this.Check_Port(host)
		}
	}()

	go func() {
		wg.Done()
		for {
			res, ok := <-res_chan
			if !ok {
				break
			}
			log.Println(res)
			this.Save_Res_Check(res)
		}
	}()

	wg.Wait()
}

func (this *LiuBei) Do_Web_Check() {
	var wg sync.WaitGroup
	wg.Add(2)
	web_chan := make(chan Web, 100)
	res_chan := make(chan Web_Status, 100)

	go func() {
		for _, web := range this.Query_Web_From_Db() {
			web_chan <- web
		}
	}()

	go func() {
		for {
			web, ok := <-web_chan
			if !ok {
				close(res_chan)
				break
			}
			res_chan <- this.Check_Web_Status(web)
		}
	}()

	go func() {
		for {
			res, ok := <-res_chan
			if !ok {
				break
			}
			this.Save_Webstatus_To_Db(res)
		}
	}()
}
