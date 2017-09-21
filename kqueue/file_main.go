package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
)

func main() {
	var filePath string

	flag.StringVar(&filePath, "path", "", "file path")
	flag.Parse()

	if len(filePath) == 0 {
		panic("Invalid filePath")
	}

	f, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	q, err := syscall.Kqueue()
	if err != nil {
		panic(err)
	}

	var ev = syscall.Kevent_t{
		Ident:  uint64(f.Fd()),
		Filter: syscall.EVFILT_VNODE,
		Flags:  syscall.EV_ADD | syscall.EV_CLEAR,
		Fflags: syscall.NOTE_WRITE,
	}

	_, err = syscall.Kevent(q, []syscall.Kevent_t{ev}, nil, nil)
	if err != nil {
		panic(err)
	}

	var evt = []syscall.Kevent_t{{}}
	var timeout = syscall.Timespec{Nsec: 0, Sec: 2}
	for {
		n, errn := syscall.Kevent(q, nil, evt, &timeout)
		if errn != nil {
			panic(err)
		} else if n == 0 {
			fmt.Printf("time out\n")
		} else if evt[0].Flags&syscall.EV_ERROR != 0 {
			fmt.Println(evt[0].Data)
		} else {
			fmt.Printf("something write:\n")
		}
	}
}
