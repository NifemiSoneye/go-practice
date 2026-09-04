package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == "free" {
		msgSlice := messages[0 : 1]
		return msgSlice , nil
	} else if plan == "pro" {
		msgSlice := messages[0:]
		return msgSlice , nil
	} else {
		var m [] string
		return m , errors.New("unsupported plan")
	}
}
