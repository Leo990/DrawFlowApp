package dao

import (
	"io/ioutil"
	"log"
)

func WriteProgram(program string, path string) {
	b := []byte(program)
	err := ioutil.WriteFile(path, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
