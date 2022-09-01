package utils

import (
	"net"
	"time"
)

func InArray[T []E, E comparable](slice T, data E) bool {
	for _, e := range slice {
		if e == data {
			return true
		}
	}
	return false
}

func AvailablePort(host string, port string) string {
	ports := []string{port, "8080", "8888", "9090", "9999"}
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			return port
		}
		if conn != nil {
			defer conn.Close()
			continue
		}
	}
	return port
}
