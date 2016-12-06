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
import(
    "github.com/yeahno/golog"
    "os"
)
var Logger *golog.Logger
func init(){
    log_filepath := "test.log"
    if log_filepath == "" {
        Logger = golog.NewLogger(os.Stdout)
    } else {
        fd, err := os.OpenFile(log_file_path,os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
        if err != nil {
            panic(err)
        }
        Logger = golog.NewLogger(fd)
    }
}

func main(){
    
}
```
####使用方法

```
import(

)
```