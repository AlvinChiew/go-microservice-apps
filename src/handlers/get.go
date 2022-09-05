package handlers

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func GetHostName(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching host details...")
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(hostname)
}

func GetHostIP(w http.ResponseWriter, r *http.Request) net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP
}
