package main

import (
	"fmt"
	"net/http"
	"root/examples"
)

func MyFunc() {
	fmt.Println("Ovaj kod je super")
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

type MiddlewareFunc func(http.HandlerFunc, string) http.HandlerFunc

func MyMiddleware(arg1 string, arg2 int) MiddlewareFunc {
	return func(next http.HandlerFunc, arg3 string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Implementacija middlewarea
			// Možete koristiti arg1, arg2, arg3 ovdje
			// ...

			// Izvršavanje sljedećeg middlewarea
			next(w, r)
		}
	}
}

type Map struct {
	mapa SuperMap
}

func New() *Map {
	return &Map{
		mapa: make(map[string]string),
	}
}

type SuperMap map[string]string

func (m *Map) Add(key, value string) SuperMap {
	m.mapa[key] = value
	return m.mapa
}

// 🐰😎//🌞
func main() {
	PrintAll("Janko", Setting(MyFunc, Branko, MyFunc))
	fmt.Println(examples.Returner("Janko the KING")())

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("fap", "asd")
	})

	// http.ListenAndServe(":5000", nil)

	// middlewareChain := MyMiddleware("arg1", 123)(
	// 	OtherMiddleware1,
	// 	OtherMiddleware2,
	// )

	examples.Execute(examples.Maker())

	fja := New()

	fmt.Println(fja.Add("sda", "sad"))

	examples.HandlerMaker(examples.Janko)("asd", "asd")
}
