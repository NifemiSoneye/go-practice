package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (authi authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authroization : Basic %s : %s" , authi.username , authi.password)
}