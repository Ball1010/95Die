package beginning

import "fmt"
import "errors"

func Hello (name string)(string, error) {
	if name ==""{
		return "",  errors.New("empty name")
	}
	msg:=fmt.Sprintf("Hello %v. Welcome to Go.\n---Ball[:", name)
	return msg, nil

}