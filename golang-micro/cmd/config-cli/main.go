package main

import (
	"flag"
	"log"
	"time"

	"github.com/Cloudera-Sz/golang-micro/clients/config/cli"
)

func main() {
	//os.Setenv("LOCALHOST", "192.168.1.21")
	configPath := flag.String("c", "dev", "config file path")
	dialTimeout := flag.Int("dt", 5, "dail timeout")
	requestTimeout := flag.Int("rt", 10, "request timeout")
	flag.Parse()
	err := cli.SyncConfig(
		time.Duration(*dialTimeout)*time.Second,
		time.Duration(*requestTimeout)*time.Second,
		*configPath,
	)
	log.Println("err", err)
}
