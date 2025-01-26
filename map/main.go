package main

import "fmt"

func main() {
	studentgrades := make(map[string]int)

	studentgrades["prince"] = 100
	studentgrades["Alice"] = 200
	studentgrades["Bob"] = 300
	studentgrades["charlie"] = 400

	fmt.Println("max", studentgrades["Bob"])
	studentgrades["Bob"] = 100
	fmt.Println("max", studentgrades["Bob"])
	// delete(studentgrades, "Bob")

	grades,exist :=studentgrades["prince"]
	fmt.Println(grades,exist)

	for i,j :=range studentgrades{
		fmt.Println(i,j)
	}
	person :=map[string]int{
		"a":33,
		"b":22,
		"c":22,
	}
	for i,j :=range person{
		fmt.Println(i,j)
	}
}
