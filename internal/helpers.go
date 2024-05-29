package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyPrintJSON(jsonStr string) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(jsonStr), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prettyJSON.String())
}
