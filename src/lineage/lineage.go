package main

import (
	"lineage/utils"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("start")

	utils.MS_init_listener()
	//	TODO: concurrent event
	//	go listen_keyboard()

	for {
	}
}
