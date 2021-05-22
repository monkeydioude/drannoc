package misc

import (
	"encoding/json"
	"fmt"
)

func PPrintln(v interface{}) {
	d, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(d))
}
