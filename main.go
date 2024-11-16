package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type cources struct {
		Name     string `json:"coursename"`
		Price    int
		Platform string `json:"website"`
		Password string `json:"-"`
		//Tags []string

	}

	fmt.Println("convert into json data")

	Playlist := []cources{
		{"golang", 499, "onlinecoding", "abcs678"},
		{"react", 199, "onlinecoding", "axyzghs678"},
		{"python", 99, "onlinecoding", "klmn678"},
	}

	Jsondata, err := json.MarshalIndent(Playlist, "", "\t")

	if err != nil {
		panic(err)
	}

	//byte type returned by jsondat
	fmt.Println("data encoded is: ", Jsondata)

	//nomal data to in format means string conversion of jsondat
	fmt.Println("data encoded is  :", string(Jsondata))

}
