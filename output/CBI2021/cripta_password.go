package main

import "fmt"

func main() {
	var c rune
	for fmt.Scanf("%c", c); c != '\n'; fmt.Scanf("%c", c) {
		fmt.Print(password(c))
	}
}

func password(c rune) string {
	switch c {
	case 'a':
		return "@"
	case 'e':
		return "3"
	case 'i':
		return "1"
	case 'o':
		return "0"
	case 's':
		return "$"
	case 'E':
		return "€"
	case 'L':
		return "£"
	}
	return string(c)
}
