package helper

import (
  "gopkg.in/mgo.v2"

  "github.com/spf13/viper"
)

func MgoChaos() {
  GMgoConnectionStringURI = viper.GetString("mongodb.addr")
  GMgoDatabaseName = viper.GetString("mongodb.database")
}

var gMgoSession *mgo.Session

//mgo.PrimaryPreferred
//5 mode：
//primary Perform all read operations on the master node
//primaryPreferred Priority on the main node to read, if the primary node is not available, and then from the slave operation
//secondary All read operations are performed on the slave node
//secondaryPreferred Priority to read from the slave node, if all slave nodes are unavailable, and then from the master node operation。
//nearest According to the network delay time, the nearest read operation, regardless of the node type。
//default is strong mode its named primary. eg. gMgoSession.SetMode(mgo.Strong)
//des:https://segmentfault.com/a/1190000000460489
var GMgoConnectionStringURI = ""
var GMgoDatabaseName = ""

func NewMgoSession() *mgo.Session {
  if gMgoSession == nil {
    var err error
    gMgoSession, err = mgo.Dial(GMgoConnectionStringURI)
    if err != nil {
      panic(err)
    }
    if err = gMgoSession.Ping(); err != nil {
      panic(err)
    }
    gMgoSession.SetMode(mgo.PrimaryPreferred, false)
  }
  return gMgoSession
}

func MgoExecute(colName string, q func(*mgo.Collection) error) error {
  s := NewMgoSession().Clone()
  defer s.Close()
  return q(s.DB(GMgoDatabaseName).C(colName))
}

func MgoExecuteBulk(colName string, q func(*mgo.Bulk) error) error {
  s := NewMgoSession().Clone()
  defer s.Close()
  return q(s.DB(GMgoDatabaseName).C(colName).Bulk())
}
