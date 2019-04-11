package helper

import (
  "fmt"
  "github.com/takama/daemon"
  "os"
  "path/filepath"
)

/*
myservice install  安装服务，以后重启机器时会自己启动
myservice remove  删除服务
myservice start  启动服务，当用kill -9 或程序死掉后会自动重启
myservice stop  停止服务
myservice status  查看服务
./myservice   不以daemon启动
*/
func startDaemon() {
  os.Chdir(filepath.Dir(os.Args[0]))
  if len(os.Args) > 1 {
    proc := filepath.Base(os.Args[0])
    svc, err := daemon.New(proc, proc, proc+".service")
    if err == nil {
      var msg string
      switch os.Args[1] {
      case "install":
        msg, err = svc.Install()
      case "remove":
        msg, err = svc.Remove()
      case "start":
        msg, err = svc.Start()
      case "stop":
        msg, err = svc.Stop()
      case "status":
        msg, err = svc.Status()
      default:
        msg = "Usage: " + proc +
          " install | remove | start | stop | status"
      }
      fmt.Println(msg)
    }
    if err != nil {
      fmt.Println("Error:", err)
    }
    os.Exit(1)
  }
}
