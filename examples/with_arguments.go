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
