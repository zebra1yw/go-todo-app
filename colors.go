package go_todo_app

import "fmt"

const (
	ColorDefault = "\x1b[39m"
	ColorRed     = "\x1b[91m"
	ColorGreen   = "\x1b[32m"
	ColorBlue    = "\x1b[94m"
	ColorGrey    = "\x1b[90m"
)

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func grey(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGrey, s, ColorDefault)
}
