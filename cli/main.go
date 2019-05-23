package main

import "log"

// main calls the root command. The actual processing happens when
// the executed command runs the associated handler.
func main() {
	if err := Execute(); err != nil {
		log.Println(err)
	}
}
