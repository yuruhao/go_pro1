package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("template/hello.html")
	if err != nil {
		return
	}

	u1 := User{
		Name: "test",
		Age:  18,
	}

	//渲染模板
	err = t.Execute(w, u1)
	if err != nil {
		fmt.Printf("render template error: %v", err)
		return
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, error: %v", err)
		return
	}
	//fmt.Printf("start init gin")
	//r := gin.Default()
	//r.Run("0.0.0.0:8080")
}
