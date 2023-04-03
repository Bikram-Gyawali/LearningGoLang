package main

import( 
	"fmt"
	"fmt"
	"os"
	"strings"
)

func createBill() bill{

	reader:=bufio.NewReader(os.Stdin)
	name,_ := reader.ReadString("\n")
	name  string.TrimSpace(name)
	b:= newBill(name)
	fmt.Println("Created the bill  ",b.name)

	return b
}


func main(){
	myBill:=createBill()

	fmt.Println(myBill)
}