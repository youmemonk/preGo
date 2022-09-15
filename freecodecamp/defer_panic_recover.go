package main

import (
	"net/http"
)

func behaviours() {
	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// a, b := 1, 0
	// fmt.Println(a / b)
	// fmt.Println("start")
	// panic("Something went wrong")
	// fmt.Println("end")

	helperFunc()
}

func helperFunc() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}
}
