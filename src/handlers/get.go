package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func GetHostDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching host details...")
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ip_addr := conn.LocalAddr().(*net.UDPAddr).IP

	fmt.Println(hostname, ip_addr)

	response := map[string]string{
		"hostname": hostname,
		"ip":       ip_addr.String(),
	}
	json.NewEncoder(w).Encode(response)
}
