#GO_LOG

This is a log library based on golang log.

### Features

- light weight
- easy use
- level support

### Quick Start

##### Installation

```
go get github.com/yeahno/golog
```

##### Usage

```
package main
import (
    "fmt"
    "github.com/yeahno/golog"
    "log"
    "os"
)

func main() {
    //log print to terminal
    Logger := golog.NewLogger(os.Stdout)

    Logger.Debug("app begin run")                                          //DEBUG 2016/12/06 18:17:54.593984 logger.go:42: app begin run
    Logger.Info("running", "extra msg")                                    //INFO 2016/12/06 18:17:54.594487 logger.go:66: running [extra msg]
    Logger.Warning("main.main", "log a slice", []int{1, 2, 3})             //WARNING 2016/12/06 18:17:54.594531 logger.go:78: main.main [log a slice [1 2 3]]
    Logger.Info("main.main", "running", map[string]string{"key": "value"}) //INFO 2016/12/06 18:17:54.594555 logger.go:66: main.main [running map[key:value]]
    err := fmt.Errorf("file not exist")
    Logger.Error("stoped", "error:", err) //ERROR 2016/12/06 18:17:54.594570 logger.go:90: stoped [error: file not exist]

    Logger.Disable(golog.DEBUG | golog.WARNING)
    Logger.Debug("this message will not log")   //
    Logger.Warning("this message will not log") //

    Logger.Enable(golog.WARNING)
    Logger.Debug("this message will not log") //
    Logger.Warning("this message will log")   //WARNING 2016/12/06 18:17:54.594578 logger.go:78: this message will log

    Logger.SetFlags(log.Ltime)
    Logger.Info("log format has changed") //INFO 18:17:54 log format has changed

    //change log to file
    log_filepath := "test.log"
    fd, err := os.OpenFile(log_filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        panic(err)
    }
    Logger.SetOutput(fd)
    Logger.Info("log write into file")

}

```