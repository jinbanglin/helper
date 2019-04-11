package helper

import (
  "github.com/spf13/viper"
  "fmt"
  "github.com/fsnotify/fsnotify"
)

func Chaos(fn string, f ...func()) {
  viper.SetConfigType("toml")
  viper.SetConfigFile(fn)
  if err := viper.ReadInConfig(); err != nil {
    panic(fmt.Errorf("fatal error config file: %s \n", err))
  }
  for _, v := range f {
    v()
  }
  viper.WatchConfig()
  viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Println("config file changed:", e.Name)
    for _, v := range f {
      v()
    }
  })
}
