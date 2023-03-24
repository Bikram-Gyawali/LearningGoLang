package main

import (
	"fmt"
	"strings"
	"sort"
)

func main(){
	samman:="nameste sabai jana lai"


	//string package
	// strings.Contains is like searching inside string
	// fmt.Println(strings.Contains(samman,"bz"))
	// fmt.Println(strings.ReplaceAll(samman,"sabai","ek"))
	// fmt.Println(strings.ToUpper(samman))
	// fmt.Println(strings.Index(samman,"lai"))
	fmt.Println(strings.Split(samman,"a"))


	ages:=[]int{21,32,1,32,43,545,6,154,15,62,643,56}

	sort.Ints(ages)

	fmt.Println(ages)


	index:=sort.SearchInts(ages,32)
	fmt.Println(index)




}