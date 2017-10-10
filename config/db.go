package config

import (
	"gopkg.in/mgo.v2"
	"crypto/tls"
	"time"
)

var session *mgo.Session

var tlsConfig = &tls.Config{}

var dialInfo = &mgo.DialInfo{
	Addrs:    []string{	""},
	Timeout:  60 * time.Second,
	Database: "",
	Username: "",
	Password: "",
}
