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
		fmt.Fprintln(os.Stderr, "pretty: %v", e)
	}
	fmt.Println(dst.String())
}
