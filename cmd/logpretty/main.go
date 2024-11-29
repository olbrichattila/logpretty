// Package main reads from pipe and tries to recognize content and displays in a more readable way
package main

import pipe "github.com/olbrichattila/logpretty/internal"

func main() {
	piper := pipe.New()
	piper.Listen()
}
