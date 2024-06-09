package main

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/malinoOS/malino/libmalino"
)

var Version string = "undefined"

func main() {
	defer libmalino.ResetTerminalMode()
	libmalino.ClearScreen()
	fmt.Printf("doomOS v%v - malino example\n", Version)

	// mount /proc
	if err := libmalino.MountProcFS(); err != nil {
		doomPanic(err, "mounting /proc")
	}

	// mount /dev
	if err := os.Mkdir("/dev", 0777); err != nil {

	}
	if err := syscall.Mount("udev", "/dev", "devtmpfs", syscall.MS_NOSUID, ""); err != nil {
		panic(err)
	}

	// start fbdoom
	if err := libmalino.SpawnProcessStdioFiles("/bin/fbdoom", "/", []string{}, true, true, "-iwad", "DOOM.WAD"); err != nil {
		doomPanic(err, "running DOOM")
	}
	libmalino.ShutdownComputer()
}

func doomPanic(err error, where string) {
	fmt.Println("\n--- doomOS \033[91mPANIC!\033[39m ---")
	fmt.Println(err.Error())
	fmt.Println("This happened while " + where)
	fmt.Println("\nThe system is halted.")
	for {
		time.Sleep(time.Hour)
	}
}
