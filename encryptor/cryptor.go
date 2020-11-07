package encryptor

import (
	"encoding/base64"
	"encoding/json"
)

// Encode encodes the input in base64
// It can optionally zip the input before encoding
func Encode(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(b)
}

// It can optionally unzip the input after decoding
func Decode(in string, obj interface{}) {
	b, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		panic(err)
	}

	/*if compress {
		b = unzip(b)
	}*/

	err = json.Unmarshal(b, obj)
	if err != nil {
		panic(err)
	}
}
