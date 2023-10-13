package jq

import (
	"encoding/json"
	"log"

	"github.com/itchyny/gojq"
)

func ParseJQ(jsonData []byte, query string) ([]byte, error) {

	// Parse the JQ query
	parsedQuery, err := gojq.Parse(query)
	if err != nil {
		log.Fatalf("JQ query parsing error: %v", err)
		return nil, err
	}

	var input interface{}
	err = json.Unmarshal(jsonData, &input)
	if err != nil {
		log.Fatalf("JSON parsing error: %v", err)
		return nil, err
	}

	iter := parsedQuery.Run(input) // or query.RunWithContext

	result := []any{}
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		log.Printf("ITEM %v", v)
		result = append(result, v)
	}

	// result, _ := iter.Next()
	// if result == nil {
	// 	log.Fatalf("JQ query execution error")
	// 	return nil, err
	// }

	// if len(result) == 1 {
	// 	log.Printf("ONE ROW = %T", result[0])
	// 	resultJSON, err := gojq.Marshal(result[0])
	// 	if err != nil {
	// 		log.Fatalf("JSON conversion error: %v", err)
	// 		return nil, err
	// 	}
	// 	return resultJSON, nil

	// } else {
	resultJSON, err := gojq.Marshal(result)
	if err != nil {
		log.Fatalf("JSON conversion error: %v", err)
		return nil, err
	}
	return resultJSON, nil

	// }
	// Convert the result back to a JSON byte array

}
