package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var y = 14
var z = "Superman, batsie"
var a = `Let us check out "Raw String Literal"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	type Dictionary map[string]interface{}
	data := []Dictionary{}
	dict1 := map[string]interface{}{"key": 1}
	dict2 := Dictionary{"key": 2}
	data = append(data, dict1, dict2)
	fmt.Printf("\n%v\n",dict1)
	fmt.Printf("\n%v",data)
	
	//type Dictionary map[string]interface{}

	newdata1 := []Dictionary{
		{
			"metrics": []Dictionary{
				{ "tag_name": "output_current", "id": 3 },
				{ "tag_name": "input_voltage", "id": 2 },
			},
			"port":       161,
			"timeout":    1,
			"sleep_time": 5,
		},
		{
			"metrics": []Dictionary{
				{ "tag_name": "destructor", "id": 10 },
			},
			"port":       161,
			"timeout":    1,
			"sleep_time": 4,
		},
	}
	
	buf, err := json.Marshal(newdata1)
	if err !=nil {
		panic(err)
	}

	err = ioutil.WriteFile("fileame.json", buf, 0644)
	if err !=nil {
		panic(err)
	}

	jsonString := string(buf)
	fmt.Println(jsonString)

	type Metric struct {
		TagName string `json:"tag_name"`
		ID      int    `json:"id"`
	}
	
	type Data struct {
		Port      int      `json:"port"`
		Timeout   int      `json:"timeout"`
		SleepTime int      `json:"sleep_time"`
		Metrics   []Metric `json:"metrics"`
	}
	
	newdata2 := []Data{
		{
			Port:      161,
			Timeout:   1,
			SleepTime: 5,
			Metrics: []Metric{
				{TagName: "output_current", ID: 3},
				{TagName: "input_voltage", ID: 2},
			},
		},
		{
			Port:      161,
			Timeout:   1,
			SleepTime: 4,
			Metrics: []Metric{
				{TagName: "destructor", ID: 10},
			},
		},
	}
	buf1, err1 := json.Marshal(newdata2)
	if err1 !=nil {
		panic(err1)
	}

	err1 = ioutil.WriteFile("fileame1.json", buf1, 0644)
	if err1 !=nil {
		panic(err1)
	}

	jsonString1 := string(buf1)
	fmt.Println(jsonString1)
	

}