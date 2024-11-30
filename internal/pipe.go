// Package pipe is the main entry point of formatter
package pipe

import (
	"bufio"
	"fmt"

	"os"

	formatter "github.com/olbrichattila/logpretty/internal/formatters"
)

// New creates a new pipe, which listens on standard input and output formatted text
func New() Piper {
	return &pipe{}
}

// Piper implements Listen on standard input
type Piper interface {
	Listen()
}

type pipe struct {
}

func (*pipe) Listen() {
	err := formatter.Validate()
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(formatter.Run(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
