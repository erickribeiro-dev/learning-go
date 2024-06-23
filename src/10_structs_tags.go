package main

import (
	"fmt"
	"reflect"
)

type album struct {
	// Tags describe some information about the fields
	// Can be used by validation library or framework
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func structs_tags() {
	// The reflect package can be used to access the Tag values
	t := reflect.TypeOf(album{})
	field, _ := t.FieldByName("ID")
	fmt.Println(field.Tag) // json:"id"
}
