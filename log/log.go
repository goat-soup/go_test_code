package log

import "fmt"

func Info(msg string) error {
	_, err := fmt.Printf("info : %s\n", msg)
	return err
}

func Error(msg string) error {
	_, err := fmt.Println("error: ", msg)
	return err
}
