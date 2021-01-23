package main

import (
	"encoding/json"
	"fmt"
)

package main

import (
"encoding/json"
"fmt"
)

type Tagged struct {
	FieldA int    `json:"field_a"`
	FieldB string `json:"field_b"`
	fieldC string `json:"field_c"`
}

func main() {
	someJsonContent := `
	{
        "field_a": 15,
        "field_b": "Эта строка будет записана в поле",
		"field_c": "Эта строка не будет записана"
	}
	`

	taggedStruct := Tagged{}
	if err := json.Unmarshal([]byte(someJsonContent), &taggedStruct); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", taggedStruct)
}

