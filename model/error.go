package model

import "fmt"

type Error struct {
	Message string `json:"msg"`
	File    string `json:"file"`
	Line    uint   `json:"line"`
}

func (receiver *Error) Error() string {
	return fmt.Sprintf("error in %s on line %d: %s", receiver.File, receiver.Line, receiver.Message)
}
