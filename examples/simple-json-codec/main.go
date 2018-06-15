package main

import (
	"log"

	json "github.com/sniperkit/codecs/plugin/codec/json"
)

var jsonCodec json.JsonCodec

func main() {

	// make a big object
	obj := map[string]interface{}{}
	obj["name"] = "Mat"
	obj["age"] = 30
	obj["address"] = map[string]interface{}{
		"street":  "Pearl Street",
		"city":    "Boulder",
		"state":   "CO",
		"country": "USA",
	}
	obj["animals"] = map[string]interface{}{
		"favourite": []string{"Dog", "Cat"},
	}

	bytes, err := jsonCodec.Marshal(obj, nil)

	if err != nil {
		log.Printf("Failed to marshal simple JSON: %s", err)
	} else {
		log.Printf("%v", string(bytes))
	}

}
