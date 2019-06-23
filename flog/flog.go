package flog

import (
	"log"
	"os"
)

//loggers
var (
	DEBUG *log.Logger
	INFO  *log.Logger
	WARN  *log.Logger
	ERROR *log.Logger
	FATAL *log.Logger
	Flog  *log.Logger
)

const (
	//NET ..
	NET = "[net] "
	//PNG ..
	PNG = "[pinger] "
	//CLI ..
	CLI = "[client] "
	//DEC ..
	DEC = "[decode] "
	//MES ..
	MES = "[message] "
	//STR ..
	STR = "[store] "
	//MID ..
	MID = "[msgids] "
	//TST ..
	TST = "[test] "
	//STA ..
	STA = "[state] "
	//GateWay ..
	GateWay = "[GateWay] "
	//Server ..
	Server = "[Server] "
)

//The default output for all the loggers is set to ioutil.Discard
func init() {
	DEBUG = log.New(os.Stdout, "\x1b[36m[DEBU]\x1b[0m ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	INFO = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	WARN = log.New(os.Stdout, "\x1b[33m[WARN]\x1b[0m ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	ERROR = log.New(os.Stdout, "\x1b[31m[ERRO]\x1b[0m ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	FATAL = log.New(os.Stdout, "\x1b[44;11m[FATA]\x1b[0m ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

// func init() {
// 	DEBUG = log.New(os.Stdout, "[DEBU] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
// 	INFO = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
// 	WARN = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
// 	ERROR = log.New(os.Stdout, "[ERRO] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
// 	FATAL = log.New(os.Stdout, "[FATA] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
// }
