package main

import (
	"net"
	"time"
)

func Check(url string, port string) string {
	addr := url + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", addr, timeout)

	var status string

	if err != nil {
		status = "The website is DOWN"
	} else {
		status = "The website is UP"
	}

	return status + "\nFrom: " + conn.LocalAddr().String() + "\nTo: " + conn.RemoteAddr().String()
}
