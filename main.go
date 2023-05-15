package main

import (
	"fmt"
	"root/examples"
)

func MyFunc() {
	fmt.Println("ovaj kod je super")
}

func Branko() {
	fmt.Println("Branko")
}

type Reader func()

func Setting(fn ...Reader) Reader {
	return func() {
		fmt.Println("Janko je legenda", "=>")
		fmt.Println("in front of req")

		for _, fn := range fn {
			fn()
		}

		fmt.Println("After req")
	}
}

func PrintAll(name string, fn Reader) {
	fmt.Println(name)
	fn()
}

// 🐰😎//🌞
func main() {
	PrintAll("Janko", Setting(MyFunc, Branko, MyFunc))
	fmt.Println(examples.Returner("Janko the KING")())

	// http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Add("fap", "asd")
	// })

	// http.ListenAndServe(":5000", nil)

	examples.Execute(examples.Maker())
}
