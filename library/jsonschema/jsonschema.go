package main

import (
	"encoding/json"
	"fmt"

	"github.com/qri-io/jsonschema"
)

func main() {
	var schemaData = []byte(`{
      "title": "Person",
      "type": "object",
      "properties": {
          "firstName": {
              "type": "string"
          },
          "lastName": {
              "type": "string"
          },
          "age": {
              "description": "Age in years",
              "type": "integer",
              "minimum": 0
          },
          "friends": {
            "type" : "array",
            "items" : { "title" : "REFERENCE", "$ref" : "#" }
          }
      },
      "required": ["firstName", "lastName"]
    }`)

	rs := &jsonschema.RootSchema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	var valid = []byte(`{
    "firstName" : "George",
    "lastName" : "Michael"
    }`)

	if errors, _ := rs.ValidateBytes(valid); len(errors) > 0 {
		panic(errors)
	}

	var invalidPerson = []byte(`{
    "firstName" : "Prince"
    }`)
	if errors, _ := rs.ValidateBytes(invalidPerson); len(errors) > 0 {
		fmt.Println(errors[0].Error())
	}

	var invalidFriend = []byte(`{
    "firstName" : "Jay",
    "lastName" : "Z",
    "friends" : [{
      "firstName" : "Nas"
      }]
    }`)
	if errors, _ := rs.ValidateBytes(invalidFriend); len(errors) > 0 {
		fmt.Println(errors[0].Error())
	}
}
