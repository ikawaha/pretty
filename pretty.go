package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	var j bytes.Buffer
	io.Copy(&j, os.Stdin)
	var dst bytes.Buffer
	if e := json.Indent(&dst, j.Bytes(), "", "    "); e != nil {
		fmt.Fprintf(os.Stderr, "pretty: %v\n", e)
	}
	fmt.Println(dst.String())
}
