package main

import "log"

func main() {
	defer log.Println("[racgo] ...END")
	log.Println("[racgo] START...")
}
