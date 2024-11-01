package util

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func BuildJson(s interface{}, filename string) {

	file, err := os.Create(fmt.Sprintf("%s.json", filename))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(&s); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}

func ReadJson(s interface{}, filename string) {

	file, err := os.Open(fmt.Sprintf("%s.json", filename))

	if err != nil {
		fmt.Println("build JSON Error : ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
	}
}
