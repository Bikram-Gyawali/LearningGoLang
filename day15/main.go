package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursname"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("json create")

	// EncodeJson()

	DecodeJson()
}

func EncodeJson() {
	jsCourses := []course{
		{"javascript", 299, "online.com", "bikramfdashah", []string{"web=dev", "js", "api"}},
		{"react", 249, "online.com", "bikramfdashah", []string{"web=dev", "js", "api", "react"}},
		{"nodejs", 799, "online.com", "bikramfdashah", []string{"fullstack", "js", "api", "nodejs"}},
		{"golang", 599, "online.com", "bikramfdashah", []string{}},
	}

	//package the data as json data

	finalaJson, err := json.MarshalIndent(jsCourses, "   ", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalaJson)

}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursname": "javascript",
		"price": 299,
		"website": "online.com",
		"tags": [
				"web=dev",
				"js",
				"api"
		]
}

	`)

	var courseOnline course 

	checkValid:=json.Valid(jsonData)

	if checkValid {
		fmt.Println("Json format is valid")
		json.Unmarshal(jsonData,&courseOnline)
		fmt.Printf("%#v\n",courseOnline)

	}else{
		fmt.Println("'Json is not valid format")
	}


	// some cases where youjust want to add data to keye value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonData,&myOnlineData);

	// fmt.Printf("%#v\n",myOnlineData)


	for k , v := range myOnlineData{
		fmt.Printf("key %v and value is %v and type is : %T\n",k,v,v)
	}
}
