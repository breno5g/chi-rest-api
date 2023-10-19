package main

import (
	"fmt"

	"github.com/breno5g/chi-rest-api/configs"
)

func main() {
	config, _ := configs.InitConfig(".")
	fmt.Println(config)
}
