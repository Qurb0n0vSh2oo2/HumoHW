package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func Plus(w http.ResponseWriter, r *http.Request) {
	first := mux.Vars(r)["first"]
	znak := mux.Vars(r)["znak"]
	second := mux.Vars(r)["second"]

	firstInt, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}

	secondInt, err := strconv.Atoi(second)
	if err != nil {
		panic(err)
	}
	sum := firstInt + secondInt
	subtraction := firstInt - secondInt
	multiplication := firstInt * secondInt
	division := firstInt / secondInt

	if znak == "+" {
		response := fmt.Sprint(sum)
		fmt.Fprint(w, response)
	} else if znak == "-" {
		response := fmt.Sprint(subtraction)
		fmt.Fprint(w, response)
	} else if znak == "*" {
		response := fmt.Sprint(multiplication)
		fmt.Fprint(w, response)
	} else if znak == "div" {
		response := fmt.Sprint(division)
		fmt.Fprint(w, response)
	}

}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/{first}/{znak}/{second}", Plus)

	err := http.ListenAndServe("127.0.0.1:8000", mux)

	if err != nil {
		fmt.Println(err)
	}
}

//http://127.0.0.1:8000/2/+/3

// /hello?name=yakub&age=19

// /hello/{name}/{age}

// name:yakub age:19

//http://127.0.0.1:8000/bye?name=umed&age=19

//http://127.0.0.1:8000/%7B2%7D/%7B+%7D%7B3%7D
