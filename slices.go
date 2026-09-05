package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
case "free":
		msgSlice := messages[0 : 1]
		return msgSlice , nil
	case "pro":
		msgSlice := messages[:]
		return msgSlice , nil
	default:
		var m [] string
		return m , errors.New("unsupported plan")
	} 
}
