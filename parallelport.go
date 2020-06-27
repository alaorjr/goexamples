package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("LPT1", os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	f.Write([]byte{0x1B, '@'})
	f.Write([]byte{178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178, 178})
	f.Write([]byte{'\n', '\r'})
	f.Sync()
	f.Close()
}
