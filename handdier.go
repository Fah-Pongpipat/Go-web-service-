package main

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// func main() {
// 	JsonData := `[{"id":101,"name": "John", "age":12}]`
// 	var person []map[string]interface{}
// 	err := json.Unmarshal([]byte(JsonData), &person)

// 	if err != nil {
// 		fmt.Print(err)
// 	} else {
// 		// แก้ไขให้แสดงข้อมูลที่ถูก Unmarshal
// 		fmt.Println(person)
// 	}
// }
