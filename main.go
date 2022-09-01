package main

import (
	"flag"
)

func main() {
	fPort := flag.String("p", "8899", "Application's Port")
	fAppType := flag.String("t", "server", "Application's Type one of : [server|client]")
	flag.Parse()

	port := *fPort
	appType := *fAppType

	switch appType {
	case "server":
		runServer(port)
	case "client":
		runClient()
	default:
		flag.PrintDefaults()
	}
}
