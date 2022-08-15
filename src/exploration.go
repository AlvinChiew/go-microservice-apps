package main

import (
	"fmt"
	"math"
	"errors"
)

func main()	{
	var name1 string = "Alvin"
	name2 := "JS"  // "Go can infer type"
	if name2 == "" {
		fmt.Println("Hello", name1, math.Pi)
	} else if name2 == "JS"{
		fmt.Println("Hello", name1, "and", name2)
	} else {
		fmt.Println("Hello", name1, "and", name2, math.Pi)
	}

	var a [5]int  // int null value is 0
	a[2] = 7
	b := [5]int{1, 2, 3, 4, 5}  // shorthand
	c := []int{1, 2, 3}  // slice = abstraction for array without declaring size
	c = append(c, 4)
	fmt.Println(a, b, c)

	kvp := make(map[string]int)
	kvp["alvin"] = 1
	kvp["kelvin"] = 2	
	kvp["meh"] = 3
	delete(kvp, "alvin")
	fmt.Println(kvp, kvp["alvin"])

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	arr := []string{"a", "b", "c"}
	for idx, val := range arr {
		fmt.Println(idx, val)
	}

	kvp2 := make(map[string]int)
	kvp2["a"] = 1
	kvp2["b"] = 2
	for key, val := range kvp2 {
		fmt.Println(key, val)
	}

	fmt.Println(sum(1, 2))
	result, err := sqrt(-16)
	fmt.Println(result, err)

	p := person{name: "John", age: 16}
	fmt.Println(p)
	fmt.Println(p.age)
	fmt.Println(p.get_next_year_age())
	p.add_age_attr(12)	
	fmt.Println(p.age)
	fmt.Println(p.get_next_year_age())
	p.age = p.age + 1	
	fmt.Println(p.age)

	j := 6
	increment_pointer(&j)
	fmt.Println(j)  //pointer

}

func sum(x int, y int) int {
	return x + y 
}

// func can return more than a type
// also, go doesn't have `Exception`
func sqrt(x float64) (float64, error) {  
	if x < 0 {
		return 0, errors.New("Error: Undefined")
	}
	return math.Sqrt(x), nil
}

func increment_pointer(i *int){
	*i++  // `*` to dereference pointer
}

type person struct {
	name string 
	age int 
}

func (p person) get_next_year_age() int {
	return p.age + 1
}

func (p *person) add_age_attr(year int) {
	p.age = p.age + year
}

// inheritence: https://golangdocs.com/inheritance-in-golang
// polymorphism (shared method among classes): https://golangdocs.com/polymorphism-in-golang
