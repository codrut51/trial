package hello

import "fmt"

func hello() string {
	return "hello"
}

func hello_name(x string) string {
	a := fmt.Sprintf("Hello %s", x)
	return a
}

func hello_user(x string) string {
	a := fmt.Sprintf("Hello %s", x)
	return a
}

func stop() {
	fmt.Println("Stop")
}
