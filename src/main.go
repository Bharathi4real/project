package main

import (
	"fmt"

	"project.com/attendance/api"
	"project.com/attendance/config"
)

func main() {
	config.InitConfig()
	fmt.Printf("Configuration: %+v\n", config.Config)
	api.StartServer()
}
