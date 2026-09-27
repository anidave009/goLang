package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{
		"name": "John Doe",
		"age": 30,
		"city": "New York",
		"height" : 5.9,
		"isStudent": false,
		"email": "john.doe@example.com"
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

	// fields := []reflect.StructField{
	// 	{
	// 		Name: "Name",
	// 		Type: reflect.TypeOf(""),
	// 	},
	// 	{
	// 		Name: "Age",
	// 		Type: reflect.TypeOf(float64(0)),
	// 	},
	// 	{
	// 		Name: "City",
	// 		Type: reflect.TypeOf(""),
	// 	},
	// 	{
	// 		Name: "Height",
	// 		Type: reflect.TypeOf(float64(0)),
	// 	},
	// 	{
	// 		Name: "IsStudent",
	// 		Type: reflect.TypeOf(false),
	// 	},
	// 	{
	// 		Name: "Email",
	// 		Type: reflect.TypeOf(""),
	// 	},
	// }

	// structType := reflect.StructOf(fields)

	// fmt.Println(structType)

	// structValue := reflect.New(structType).Elem()

	// // Set values
	// structValue.FieldByName("Name").SetString(data["name"].(string))
	// structValue.FieldByName("Age").SetFloat(data["age"].(float64))
	// structValue.FieldByName("City").SetString(data["city"].(string))
	// structValue.FieldByName("Height").SetFloat(data["height"].(float64))
	// structValue.FieldByName("IsStudent").SetBool(data["isStudent"].(bool))
	// structValue.FieldByName("Email").SetString(data["email"].(string))

	// fmt.Println(structValue)

	//for creating dynamic struct fields based on the JSON data, you can use the following approach:
	// for field, value := range data {
	// 	fieldName := strings.Map(func(r rune) rune {
	// 		if unicode.IsLetter(r) || unicode.IsDigit(r) {
	// 			return r
	// 		}
	// 		return -1
	// 	}, field)

	// 	fieldName = strings.ToUpper(fieldName[:1]) + fieldName[1:]

	// 	fields = append(fields, reflect.StructField{
	// 		Name: fieldName,
	// 		Type: reflect.TypeOf(value),
	// 	})
	// }

	// structType2 := reflect.StructOf(fields)

	// fmt.Println(structType2)

	// structValue := reflect.New(structType2).Elem()

	// fmt.Println(structValue)

	// for field, value := range data {
	// 	fieldName := strings.ToUpper(field[:1]) + field[1:]

	// 	fieldValue := structValue.FieldByName(fieldName)

	// 	fieldValue.Set(reflect.ValueOf(value))
	// }

	// fmt.Println(structValue)

	// structString := fmt.Sprintf("%+v", structValue.Interface())

	// fmt.Println(structString)
}
