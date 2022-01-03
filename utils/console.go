package utils

import (
	"github.com/k0kubun/pp"
)

func Log(obj interface{}) {
	pp.Print(obj)
	pp.Print("\n")
}

func Logf(msg string, obj interface{}) {
	pp.Printf(msg+"\r\n%v", obj)
}
