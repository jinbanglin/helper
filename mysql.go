package helper

import (
  "github.com/go-xorm/xorm"
  "github.com/spf13/viper"
)

var engine *xorm.Engine

func MYSQLInstance() *xorm.Engine {
  if engine == nil {
    var err error
    engine, err = xorm.NewEngine("mysql", viper.GetString("mysql.addr"))
    if err != nil {
      panic(err)
      return nil
    }
    engine.ShowSQL(viper.GetBool("mysql.show_log"))
    return engine
  }
  return engine
}
