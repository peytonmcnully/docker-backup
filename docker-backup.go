package main

import (
	"flag"
	"log"
	"os"

	"github.com/discordianfish/docker-backup/backup"
)

const (
	defaultAddr  = "/var/run/docker.sock"
	defaultProto = "unix"
)

var (
	addr  = flag.String("addr", defaultAddr, "address to connect to")
	proto = flag.String("proto", defaultProto, "protocol to use (unix, tcp)")
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("Syntax: store|restore filename [container-id]")
	}

	action := flag.Arg(0)
	filename := flag.Arg(1)

	switch action {
	case "store":
		if flag.NArg() < 3 {
			log.Fatal("Error: `store` requires a container-id")
		}
		containerId := flag.Arg(2)
		log.Printf("Storing %s's volume container as %s", containerId, filename)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		b := backup.NewBackup(*addr, *proto, file)
		if err := b.Store(containerId); err != nil {
			log.Fatal(err)
		}
	case "restore":
		log.Printf("Restoring %s", filename)
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		b := backup.NewBackup(*addr, *proto, file)
		if err := b.Restore(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Invalid action %s", action)
	}
}