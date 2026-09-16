package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func main() {
	jsonData := `{
"name" : "Aniket",
"surname" : "Dave",
"age" : 25,
"city" : "Mumbai",
"isHandsome" : true,
"height" : 5.9
}`

	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for field, value := range data {
		fmt.Printf("Field: %s, Type: %T, Value: %v\n", field, value, value)
	}

	// var fields []reflect.StructField
	// fields = []reflect.StructField{
	// 	{
	// 		Name: "Name",
	// 		Type: reflect.TypeOf(""),
	// 	},
	// 	{
	// 		Name: "Surname",
	// 		Type: reflect.TypeOf(""),
	// 	},
	// 	{
	// 		Name: "Age",
	// 		Type: reflect.TypeOf(0),
	// 	},
	// 	{
	// 		Name: "City",
	// 		Type: reflect.TypeOf(""),
	// 	},
	// 	{
	// 		Name: "IsHandsome",
	// 		Type: reflect.TypeOf(false),
	// 	},
	// 	{
	// 		Name: "Height",
	// 		Type: reflect.TypeOf(0.0),
	// 	},
	// }

	//creating a dynamic struct type just like we create objects .this is instance
	// structType := reflect.StructOf(fields)
	// fmt.Println(structType)

	// // this is just a struct type, we can create a new instance of it using reflect.New
	// structValue := reflect.New(structType).Elem()

	// structValue.FieldByName("Name").SetString(data["name"].(string))
	// structValue.FieldByName("Surname").SetString(data["surname"].(string))
	// structValue.FieldByName("Age").SetInt(int64(data["age"].(float64)))
	// structValue.FieldByName("City").SetString(data["city"].(string))
	// structValue.FieldByName("IsHandsome").SetBool(data["isHandsome"].(bool))
	// structValue.FieldByName("Height").SetFloat(data["height"].(float64))

	var fields []reflect.StructField
	for field, value := range data {
		fieldName := strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				return r
			}
			return -1
		}, field)

		fieldName = strings.ToUpper(fieldName[:1]) + fieldName[1:]

		fields = append(fields, reflect.StructField{
			Name: fieldName,
			Type: reflect.TypeOf(value),
		})
	}

	structType2 := reflect.StructOf(fields)

	fmt.Println(structType2)

	structValue := reflect.New(structType2).Elem()

	fmt.Println(structValue)

	for field, value := range data {
		fieldName := strings.ToUpper(field[:1]) + field[1:]

		fieldValue := structValue.FieldByName(fieldName)

		fieldValue.Set(reflect.ValueOf(value))
	}

	fmt.Println(structValue)

	structString := fmt.Sprintf("%+v", structValue.Interface())

	fmt.Println(structString)
}
