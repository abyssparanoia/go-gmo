package parser

import "github.com/gorilla/schema"

// Encoder ... encoder
var Encoder = func() *schema.Encoder {
	return schema.NewEncoder()
}

// Decoder ... decoder
var Decoder = func() *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	return decoder
}
