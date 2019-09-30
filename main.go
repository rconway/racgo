package main

import (
	"log"
	"os"
)

// ShowEnv get the named environment variable, logs it, and returns its value
func ShowEnv(varname string) (varval string) {
	varval = os.Getenv(varname)
	log.Printf("The value of env:%v is %v\n", varname, varval)
	return
}

func main() {
	defer log.Println("[racgo] ...END")
	log.Println("[racgo] START...")
}
