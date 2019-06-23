package flog

import "testing"

func TestFlog(t *testing.T) {
	DEBUG.Println(NET, "ss")
	INFO.Println(PNG, "sddd")
	WARN.Println(DEC, "sdecg")
	ERROR.Println(MID, "vde")
	FATAL.Println(STA, "qqqq")
}
