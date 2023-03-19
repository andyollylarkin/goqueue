package router

import "io"

// CommandHandler handle broker command. Be careful with w. It may already be closed, handle Write error
type CommandHandler func(w io.Writer)
