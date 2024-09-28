package main

import (
	"flag"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "local", "Specify enviroment. Default is local")
	flag.Parse()
}
