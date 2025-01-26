package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("ddkddld")
	resp,err :=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err !=nil{
		fmt.Println("error",err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("resp",resp)
	// fmt.Println(err)
	data, err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		fmt.Println("error response body",err)
		return
	}
	fmt.Println(string(data))
}
