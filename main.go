package main

import (
	"os"
	"encoding/json"
	"fmt"
)

type configuration struct {
	Enabled bool
	Path string
}

func main() {
	file,_ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Print("Error:",err)
	}
	fmt.Println(conf.Path)
}



