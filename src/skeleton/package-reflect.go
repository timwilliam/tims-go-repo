package main

import (
	"fmt"
	"reflect"
)

// reflect is used to see our code structure during code runtime
// very helpful when making a general library so that it is easy to use

type Sample struct {
	// with StructTag
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	sample := Sample{"Tim"}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	fmt.Println(IsValid(sample))

	sample.Name = ""
	fmt.Println(IsValid(sample))
}
