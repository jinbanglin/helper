package helper

import (
  "github.com/go-xorm/xorm"
  "github.com/spf13/viper"
)

var MysqlEngine *xorm.Engine

func MysqlChaos() *xorm.Engine {
  if MysqlEngine == nil {
    var err error
    MysqlEngine, err = xorm.NewEngine("mysql", viper.GetString("mysql.addr"))
    if err != nil {
      panic(err)
      return nil
    }
    MysqlEngine.ShowSQL(viper.GetBool("mysql.show_log"))
    return MysqlEngine
  }
  return MysqlEngine
}
