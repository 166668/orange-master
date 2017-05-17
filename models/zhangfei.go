package models

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	//"sync"
	"time"
)

type ZhangFei struct{}

func (z *ZhangFei) Check_Port(host Host) (status Check_Res) {
	status.Time = time.Now()
	status.Hostname = host.Hostname
	status.Ip = host.Ip
	status.Uuid = host.Uuid

	//初始化 res 的 map 字典
	status.Udp_res = make(map[string]bool)
	status.Tcp_res = make(map[string]bool)

	if len(host.Udp_list) < 1 || host.Udp_list[0] == "" {
		log.Printf("No udp ports need to check")
	} else {
		p := make([]byte, 2048)
		for _, port := range host.Udp_list {
			udp := host.Ip + ":" + port
			log.Printf(udp)
			conn, err := net.DialTimeout("udp", udp, time.Second*2)
			if err != nil {
				log.Println(err)
				status.Udp_res[port] = false
				break
			}
			fmt.Fprintf(conn, "Hi?")
			_, err = bufio.NewReader(conn).Read(p)
			if err == nil {
				status.Udp_res[port] = true
				fmt.Printf("%s\n", p)
			} else {
				status.Udp_res[port] = false
				fmt.Printf("Some error %v\n", err)
			}
			//log.Println(status.Udp_res)
			conn.Close()
		}
	}

	if len(host.Tcp_list) < 1 || host.Tcp_list[0] == "" {
		log.Printf("No Tcp ports need to check")
	} else {
		//log.Printf("Begin Tcp check!")
		for _, port := range host.Tcp_list {
			tcp := host.Ip + ":" + port
			log.Printf(tcp)
			conn, err := net.DialTimeout("tcp", tcp, time.Second*2)
			if err != nil {
				log.Println(err)
				status.Tcp_res[port] = false
			} else {
				status.Tcp_res[port] = true
				conn.Close()
			}
		}
	}
	log.Println(status)

	//wg.Wait()
	return
}

//检测 web，返回 Web_Status
func (z *ZhangFei) Check_Web_Status(web Web) (status Web_Status) {
	status.Name = web.Name
	status.Pattern = web.Pattern
	status.Url = web.Url
	status.Time = time.Now()

	body, err := http.Get(web.Url)
	if err != nil {
		log.Println(err)
		status.Web_status = false
		return
	}

	//使用 ioutil 读取得到的响应
	robots, err := ioutil.ReadAll(body.Body)
	//关闭资源
	body.Body.Close()
	//失败返回原因
	if err != nil {
		log.Fatal(err)
		status.Web_status = false
	}

	//调用 regexp 函数查找 checkword
	word, err := regexp.MatchString(web.Pattern, string(robots))
	if err != nil {
		log.Println(err)
		status.Web_status = false
	}

	if word {
		// log.Printf("The `%s`  find in `%s`", web.Pattern, web.Url)
		status.Web_status = true
	} else {
		status.Web_status = false
	}
	return
}
