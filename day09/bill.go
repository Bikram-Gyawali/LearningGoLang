package main


import "fmt"

type bill struct{
	name string
	items map[string]float64
	tip float64
}


// function for making new bill

func newBill(name string)bill{
	newBill:=bill{
		name:name,
		// items:map[string]float64{},
		items:map[string]float64{"tea":30,"coffee":50},
		tip:0,
	}
	return newBill
}


func (b bill) format() string{
	fs:="bill breakdown: \n"
	var total float64=0

	//list items
	for k,v := range b.items{
		fs+=fmt.Sprintf("%-25v ... $%v \n",k+",",v)
		total +=v
	}


	fs+= fmt.Sprintf("%-25v ... $%0.2f\n","tip: ",b.tip)


	// total
	fs+= fmt.Sprintf("%-25v ... $%0.2f\n","total: ",total)

	b.addItem()

	return fs
}



// update the bill
func (b *bill) updateTip(tip float64){
	b.tip=tip
}

// /add an itemto the bill	
func (b bill) addItem(name string, price float64){
	b.items[name]=price
}


// // output 
// {bikramg ko bill map[coffee:50 tea:30] 0}
// bill breakdown: 
// tea,                      ... $30 
// coffee,                   ... $50 
// total:                    ... $80.00