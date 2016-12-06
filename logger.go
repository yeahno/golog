package golog

import (
	"fmt"
	"io"
	"log"
)

const (
	DEBUG   = 1 << iota //    1  1
	TRACE               //   10  2
	INFO                //  100  4
	WARNING             // 1000  8
	ERROR               //10000  16
)

/**
 * 新建一个log对象
 * @params w 输出目标文件句柄
 */
func NewLogger(w io.Writer) *Logger {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	logger := new(Logger)
	logger.log = log.New(w, "", flag)
	logger.permit = DEBUG | TRACE | INFO | WARNING | ERROR
	return logger
}

type Logger struct {
	log    *log.Logger
	permit int
}

/**
 * 记录debug日志
 * @param  class 业务或模块分类
 * @param  v 	 日志信息
 */
func (this *Logger) Debug(class string, v ...interface{}) {
	if this.permit&DEBUG == DEBUG {
		this.log.SetPrefix("DEBUG ")
		this.write(class, v)
	}
}

/**
 * 记录trace日志
 * @param  class 业务或模块分类
 * @param  v 	 日志信息
 */
func (this *Logger) Trace(class string, v ...interface{}) {
	if this.permit&TRACE == TRACE {
		this.log.SetPrefix("TRACE ")
		this.write(class, v)
	}
}

/**
 * 记录info日志
 * @param  class 业务或模块分类
 * @param  v 	 日志信息
 */
func (this *Logger) Info(class string, v ...interface{}) {
	if this.permit&INFO == INFO {
		this.log.SetPrefix("INFO ")
		this.write(class, v)
	}
}

/**
 * 记录warning日志
 * @param  class 业务或模块分类
 * @param  v 	 日志信息
 */
func (this *Logger) Warning(class string, v ...interface{}) {
	if this.permit&WARNING == WARNING {
		this.log.SetPrefix("WARNING ")
		this.write(class, v)
	}
}

/**
 * 记录error日志
 * @param  class 业务或模块分类
 * @param  v 	 日志信息
 */
func (this *Logger) Error(class string, v ...interface{}) {
	if this.permit&ERROR == ERROR {
		this.log.SetPrefix("ERROR ")
		this.write(class, v)
	}
}

/**
 * 写出日志
 * @param  class 业务或模块分类
 * @param  v 	 日志信息
 */
func (this *Logger) write(class string, v []interface{}) {
	if len(v) > 0 {
		this.log.Output(2, fmt.Sprint(class, " ", v))
	} else {
		this.log.Output(2, fmt.Sprint(class))
	}
}

/**
 * 设置日志格式
 * @param  flag 日志格式
 *      Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
 *	    Ltime                         // the time in the local time zone: 01:23:23
 *	    Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
 *	    Llongfile                     // full file name and line number: /a/b/c/d.go:23
 *	    Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
 *	    LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
 *	    LstdFlags     = Ldate | Ltime // initial values for the standard logger
 */
func (this *Logger) SetFlags(flag int) {
	this.log.SetFlags(flag)
}

/**
 * 获取日志格式
 */
func (this *Logger) GetFlags() int {
	return this.log.Flags()
}

/**
 * 设置输出目标
 * @param  w  输出目标文件句柄
 */
func (this *Logger) SetOutput(w io.Writer) {
	this.log.SetOutput(w)
}

/**
 * 允许记录类型
 * @param  permit  日志类型
 */
func (this *Logger) Enable(permit int) {
	this.permit = this.permit | permit
}

/**
 * 禁止记录类型
 * @param  permit  日志类型
 */
func (this *Logger) Disable(permit int) {
	this.permit = this.permit ^ permit
}
