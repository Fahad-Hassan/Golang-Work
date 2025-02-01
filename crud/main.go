package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userid"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer resp.Body.Close()
	// fmt.Println("resp", resp)
	// fmt.Println(err)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("response code", resp.StatusCode)
		return
	}
	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("error response body", err)
	// 	return
	// }
	// fmt.Println(string(data))

	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error decoding body", err)
		return
	}
	fmt.Println("todo", todo)
	fmt.Println("totle ", todo.Title)

	fmt.Println("completed ", todo.Completed)

}
func performPosRequest() {
	Todo := Todo{
		UserId:    23,
		Title:     "FH",
		Completed: true,
	}
	jsondata, err := json.Marshal(Todo)
	if err != nil {
		fmt.Println("error decoding body", err)
		return
	}
	//convert json to string
	jsonString := string(jsondata)
	//convert string data to reader
	jsonreader := strings.NewReader(jsonString)
	myurl := "https://jsonplaceholder.typicode.com/todos"
	resp, err := http.Post(myurl, "application/json", jsonreader)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response ", string(data))
	fmt.Println("resp status", resp.Status)
}
func performUpdateRequest() {
	Todo := Todo{
		UserId:    23,
		Title:     "FH",
		Completed: true,
	}
	jsondata, err := json.Marshal(Todo)
	if err != nil {
		fmt.Println("error decoding body", err)
		return
	}
	//convert json to string
	jsonString := string(jsondata)
	//convert string data to reader
	jsonreader := strings.NewReader(jsonString)
	myurl := "https://jsonplaceholder.typicode.com/todos/1"
	//create put Resquest

	req, err := http.NewRequest(http.MethodPut, myurl, jsonreader)
	if err != nil {
		fmt.Println("error decoding body", err)
		return
	}
	req.Header.Set("Content-type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error decoding body", err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response ", string(data))
	fmt.Println("resp status", res.Status)

}
func performDeleteRequest() {
	myurl := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("error delete request", err)
		return
	}

	req.Header.Set("Content-type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error decoding body", err)
		return
	}
	defer res.Body.Close()
	// data, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("Response ", string(data))
	fmt.Println("resp status", res.Status)

}
func main() {
	fmt.Println("learning crud")
	// performGetRequest()
	// performPosRequest()
	// performUpdateRequest()
	performDeleteRequest()

}
