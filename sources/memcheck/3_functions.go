// +build OMIT
package main

import (
	"fmt"
	"log"
	"os"
	"syscall" // HL
	"time"
)

// START_FUNC OMIT
func getMem() { // HL
	var s syscall.Sysinfo_t
	for {
		err := syscall.Sysinfo(&s)

		if err != nil {
			log.Fatal(err)
		}
		w := os.Stdout
		fmt.Fprintf(w, "RAM: %v / %v\nSWAP: %v / %v\n", s.Freeram, s.Totalram, s.Freeswap, s.Totalswap)
		time.Sleep(100 * time.Millisecond)
	}
} // HL

func main() {
	getMem() // HL
}

// END_FUNC OMIT
