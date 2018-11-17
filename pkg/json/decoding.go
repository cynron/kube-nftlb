package json

import (
	"encoding/json"

	types "github.com/zevenet/kube-nftlb/pkg/types"
)

// DecodeJSON decodes any encoded JSONnftlb object and returns a JSON string;
// the JSON string being returned is NOT indented.
func DecodeJSON(encodedJSON types.JSONnftlb) string {
	decodedJSON, err := json.Marshal(encodedJSON)
	if err != nil {
		panic(err.Error())
	}
	return string(decodedJSON)
}

// DecodePrettyJSON decodes any encoded JSONnftlb object and returns a JSON string;
// the JSON string being returned is indented.
func DecodePrettyJSON(encodedJSON types.JSONnftlb) string {
	decodedJSON, err := json.MarshalIndent(encodedJSON, "", "        ")
	if err != nil {
		panic(err.Error())
	}
	return string(decodedJSON)
}
