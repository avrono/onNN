package models

import (
	"gopkg.in/mgo.v2"
	"crypto/tls"
	"time"
	"net"
	"github.com/onNN/types"
	"errors"
)


var tlsConfig = &tls.Config{}

var dialInfo = &mgo.DialInfo{
	Addrs:    []string{	},
	Timeout:  60 * time.Second,
	Database: "",
	Username: "",
	Password: "",
}


type Datastore interface {
	InsertNeuron(*types.Neuron) error
	SomeOtherMethod() error
	BuildNN(s string) (string, error)
//	AllBooks() ([]*Book, error)
}

type DB struct {
	*mgo.Session
}

/*
Create a DB Object and return
 */
func NewDB(t string) (*DB, error) {

	switch {
		case t == "MGO_DIALWITHINFO":
			dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
				conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
				return conn, err
			}
			session, err := mgo.DialWithInfo(dialInfo)
			if err != nil {
				return nil, err //TODO: handle error - no panic
			}

			//defer session.Close()
			return &DB{session}, nil
	case t== "MGO_DIAL" :
		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)

		return &DB{session}, nil
	default:
		error_msg := "no database and/or approach supplied"
		return nil, errors.New(error_msg)
	}
}