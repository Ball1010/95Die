package beginning

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello (name string)(string, error) {
	if name ==""{
		return "",  errors.New("empty name")
	}
	msg:=fmt.Sprintf("%v Welcome to Go, %v.\n---Ball[:", randomFormat(), name)
	return msg, nil
}
func randomFormat() string{
	formats:=[]string{
		"Hi!",
		"Great to see you!",
		"Good eye mate!",
	}
	return formats[rand.Intn(len(formats))]
}