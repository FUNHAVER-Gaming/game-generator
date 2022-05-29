package service

import "fmt"

func logWithArgs(msg string, args ...any) {
	fmt.Println(fmt.Sprintf(msg, args...))
}
