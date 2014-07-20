package ks

/*
#include <windows.h>
#include <mmsystem.h>

#cgo LDFLAGS: -lwinmm
*/
import "C"

func init() {
	C.timeBeginPeriod(1)
}

func GetTimeMillis() uint64 {
	return uint64(C.timeGetTime())
}

func TimerShutdown() {
	C.timeEndPeriod(1)
}
