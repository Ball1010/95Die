package main

import (
	"fmt"
	"log"
	"github.com/Ball1010/95Die/Beginning"
	
)

func main(){
	log.SetPrefix("beginning: ")
	log.SetFlags(0)

	names:=[]string{"Diego","Di","Die","Ball"}
	messages , err:= beginning.Hellos(names)
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}