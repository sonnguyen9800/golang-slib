package colors

// ANSI escape codes for colored output
const (
	green = "\033[32m"
	red   = "\033[31m"
	reset = "\033[0m"
	bold  = "\033[1m"
	cyan  = "\036[1m"
)

// Color functions for better usage
func Green(text string) string {
	return green + text + reset
}

func Red(text string) string {
	return red + text + reset
}
func Cyan(text string) string {
	return red + text + reset
}

func Bold(text string) string {
	return bold + text + reset
}
func Reset(text string) string {
	return reset + text
}
