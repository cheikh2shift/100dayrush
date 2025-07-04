package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath(".") // look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	text := viper.Get("text")
	sleepTime := viper.GetInt("sleepTime")

	fmt.Printf("Printing %s and waiting for %v seconds\n", text, sleepTime)

	time.Sleep(time.Duration(sleepTime) * time.Second)
}
