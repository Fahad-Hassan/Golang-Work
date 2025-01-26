package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learning api url")
	myurl := "https://localhost:8079/web?debug=1keyword=33"
	fmt.Printf("%T", myurl)

	parsedurl, err := url.Parse(myurl)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T", parsedurl)
	fmt.Println("scheme", parsedurl.Scheme)
	fmt.Println("host", parsedurl.Host)
	fmt.Println("path", parsedurl.Path)
	fmt.Println("Rawquery", parsedurl.RawQuery)

	parsedurl.Path="/newpath"
	parsedurl.RawQuery = "username=ffff"
	newurl :=parsedurl.String()
	fmt.Println(newurl)
}
