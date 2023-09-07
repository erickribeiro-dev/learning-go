package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	// Tags describe some information about the fields
	// Can be used by validation library or framework
	name   string `required max:"100"`
	role   string
	salary float32
}

func structsTags() {
	t := reflect.TypeOf(Employee{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}
