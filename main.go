package main

import (
	"fmt"
	gjson "github.com/coderx31/coderx-gjson"
)

const val = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	//gjson.AddModifier("testing", func(json, arg string) string {
	//	if arg == "test" {
	//		fmt.Println("inside modifier")
	//		return toJsonString("value from modifier")
	//	}
	//	return json
	//})
	//const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	//value := gjson.Get(json, "name.last")
	//fmt.Println(value)
	//
	//modifierVal := gjson.Get(json, "name.first|@testing:test")
	//fmt.Println(modifierVal)
	//
	//gjson.DeleteModifier("testing")
	//
	//fmt.Println("after modifier delete")
	//valAfterModDel := gjson.Get(json, "name.first|@testing:test")
	//fmt.Println(valAfterModDel)

	//

	//gjson.AddModifier("[1]testing", func(json, arg string) string {
	//	if arg == "test" {
	//		fmt.Println("inside [1]testing modifier")
	//		return toJsonString("value from [1]testing modifier")
	//	}
	//	return json
	//})
	//
	//gjson.AddModifier("[2]testing", func(json, arg string) string {
	//	if arg == "test" {
	//		fmt.Println("inside [2]testing modifier")
	//		return toJsonString("value from [2]testing modifier")
	//	}
	//	return json
	//})
	//
	//value := gjson.Get(val, "name.last")
	//fmt.Println(value)
	//
	//modifierVal := gjson.Get(val, "name.first|@[1]testing:test")
	//fmt.Println(modifierVal)
	//
	//modifierVal2 := gjson.Get(val, "name.first|@[2]testing:test")
	//fmt.Println(modifierVal2)
	//
	//gjson.DeleteModifier("[2]testing")
	//
	//fmt.Println("after modifier delete")
	//valAfterModDel := gjson.Get(val, "name.first|@[1]testing:test")
	//fmt.Println(valAfterModDel)
	//
	//valAfterModDel2 := gjson.Get(val, "name.first|@[2]testing:test")
	//fmt.Println(valAfterModDel2)

	//	str := `{
	//  "name": "John Doe",
	//  "age": 30,
	//  "email": "johndoe@example.com",
	//  "city": "New York",
	//  "isEmployed": true
	//}
	//`

	//m := gjson.GetManyIntoMap(str, "name", "age", "email", "city", "isEmployed")
	//
	//fmt.Println(m["name"].String())
	//fmt.Println(m["age"].String())
	//fmt.Println(m["email"].String())
	//fmt.Println(m["city"].String())
	//fmt.Println(m["isEmployed"].String())

	//result := GetFromMap(str, "name", "age", "email", "city", "isEmployed")
	//fmt.Println(result)
	//
	//result2 := GetStr(str, "name", "age", "email", "city", "isEmployed")
	//fmt.Println(result2)
}

func toJsonString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func GetFromMap(str string, paths ...string) map[string]string {
	result := make(map[string]string)

	m := gjson.GetManyIntoMap(str, paths...)

	for _, path := range paths {
		result[path] = m[path].String()
	}
	return result
}

//func GetFromMap(str string, paths ...string) map[string]string {
//	result := make(map[string]string)
//
//	for _, path := range paths {
//		response := gjson.Get(str, path).String()
//		result[path] = response
//	}
//
//	return result
//}

func GetOne(str string, path string) string {
	res := gjson.Get(str, path)
	return res.String()
}
