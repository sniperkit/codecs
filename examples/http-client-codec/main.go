package main

import (
	"fmt"
	"log"

	// external
	pp "github.com/sniperkit/pp"

	// internal - plugin - service (with default list of codecs: json, jsonp, csv, yaml, xml and msgpack)
	codecs_svc "github.com/sniperkit/codecs/plugin/service"
)

func main() {
	fmt.Println("codecs - advanced example")

	// instanciate a new codec service
	codecService := codecs_svc.NewWebCodecService()

	// get the content type (probably from the request)
	var contentType string = "application/json"

	// get the codec
	codec, codecErr := codecService.GetCodec(contentType)
	if codecErr != nil {
		log.Fatalln("codecErr: ", codecErr)
		// handle errors - specifially ErrorContentTypeNotSupported
	}
	pp.Println("codec=", codec)

	/*
		[]bytes to object
	*/
	// get the raw data
	dataBytes := []byte(`{"somedata": true}`)

	// use the codec to unmarshal the dataBytes
	var obj interface{}
	unmarshalErr := codecService.UnmarshalWithCodec(codec, dataBytes, &obj)
	if unmarshalErr != nil {
		// handle this error
		log.Fatalln("unmarshalErr: ", unmarshalErr)
	}
	pp.Println("obj=", obj)

	// obj will now be an object built from the dataBytes

	// get some details about the kind of response we're going to make
	// (probably from the request headers again)
	var accept string = "application/json"
	var extension string = ".json"
	var hasCallback bool = false

	// get the codec to respond with (it could be different)
	responseCodec, responseCodecErr := codecService.GetCodecForResponding(accept, extension, hasCallback)
	if responseCodecErr != nil {
		log.Fatalln("responseCodecErr: ", responseCodecErr)
		// handle this error
	}
	pp.Println("responseCodec=", responseCodec)

	/*
		object to []bytes
	*/
	// get the data object
	dataObject := map[string]interface{}{"name": "Mat"}

	// use the codec service to marshal to bytes
	bytes, marshalErr := codecService.MarshalWithCodec(responseCodec, dataObject, nil)
	if marshalErr != nil {
		log.Fatalln("marshalErr: ", marshalErr)
		// handle marshalErr
	}

	// bytes would now be a representation of the data object
	pp.Println("bytes.String()=", string(bytes))

}
