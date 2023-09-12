package main

import (
	"encoding/json"
	"fmt"
)

// type course struct {
// 	Name      string
// 	Price     int
// 	Plateform string
// 	Password  string
// 	Tags      []string
// }

// creating alias...
// json me ab har baar es name se jaane jaenge or called as name change krna
type course struct {
	Name      string   `json:"coursename"`
	Price     int      `json:"price"`
	Plateform string   `json:"plateformname"`
	Password  string   `json:"-"`              // dont want this field to be reflected whoever is consuming by api
	Tags      []string `json:"tags,omitempty"` // omitempty simply means if this field is null dont through at all
}

func main() {
	fmt.Println("Json")
	// conversion of any type of data may be array slice int string whatever into json is called encoding json
	//EncodeJson()
	// conversion of json into any type of data may be array slice int string whatever is called decoding json
	DecodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"Java", 0, "youtube", "123abc", []string{"full stack", "dsa"}},
		{"JavaScript", 0, "youtube", "abc123", []string{"frontend", "coding"}},
		{"c++", 0, "youtube", "112233", nil},
	}
	// package this data as json data
	//myJson, err := json.MarshalIndent(myCourses, "", "\t")
	myJson, err := json.MarshalIndent(myCourses, "co", "\t") // (struct_name or interface, prefix, indent)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", myJson)
	//fmt.Println(myJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
	  "coursename": "JavaScript",
	  "price": 0,
	  "plateformname": "youtube",
	  "tags": ["frontend", "coding"]
	}
	`)

	var myCourses course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &myCourses)
		fmt.Printf("%#v\n", myCourses)
	} else {
		fmt.Println("JSON is not valid")
	}

	// in some cases where want to add data to key value
	var data map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &data)
	fmt.Printf("%#v\n", data)
	for k, v := range data {
		fmt.Printf("key is %v and value is %v and type is %T\n", k, v, v)
	}
}
