package examples

import "fmt"

type Return func() string

func Returner(name string) Return {
	return func() string {
		return name
	}
}

type Text func()

func Maker() Text {
	return func() {
		fmt.Println("Hello world")
	}
}

func Execute(functions ...Text) {
	for _, fn := range functions {
		fn()
	}
}

type Handler func(string, string)

func Janko(name, suerName string) {
	fmt.Println(name, suerName)

	return
}

func HandlerMaker(handler Handler) Handler {
	return func(s1, s2 string) {
		fmt.Println("Import another code")

		handler(s1, s2)
	}
}
