package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// GO offers biilt-in support for JSON encoding and decoding, including to and from built-in and custom data types.
type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	// You can customize the encoded json key name. For example here, it doesn't have to be
	// named "Page", but can be "ppppage" instead 
	Page int `json:"ppppage"`
	Fruits []string `json:"fffffruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	// fmt.Println(bolB)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(1.23)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _:= json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

	res1D := &response1 {
		Page: 	1,
		Fruits: []string{"peach", "apple", "pear"}}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	
    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

	// Now let's look at decoding data
 	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	/*
	In order to use the values in the decoded map, we’ll need to convert them to their appropriate type. 
	For example here we convert the value in num to the expected float64 type
	*/

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	/*
	We can also decode JSON into custom data types. This has the advantages of adding 
	additional type-safety to our programs and eliminating the need for 
	type assertions when accessing the decoded data.
	*/
	str := `{"ppppage": 1, "fffffruits": ["apple", "peach"]}` 
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	/*
	In the examples above we always used bytes and strings as intermediates between the 
	data and JSON representation on standard out. We can also stream JSON 
	encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	*/
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}