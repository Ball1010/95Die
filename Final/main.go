package main

import (
	"fmt"
	"github.com/Ball1010/95Die/Beginning"
	"log"
)

func main(){
	log.SetPrefix("beginning: ")
	log.SetFlags(0)

	msgf, err:= beginning.Hello("Diego")
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println(msgf)
}