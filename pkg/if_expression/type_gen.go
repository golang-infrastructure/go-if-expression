//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
)

var packageName = "package if_expression"

func main() {
	f, err := os.Open("int64.go")
	if err != nil {
		panic(err)
	}
	filedata, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	w := new(bytes.Buffer)
	start_pos := strings.Index(string(filedata), packageName)
	w.WriteString(string(filedata)[start_pos : start_pos+len(packageName)])

	ts := []string{"Byte", "Complex64", "Complex128", "Float32", "Float64", "Int", "Int8", "Int16", "Int32", "Rune", "String", "Uint", "Uint8", "Uint16", "Uint32", "Uint64", "Uintptr"}

	for _, upper := range ts {
		lower := strings.ToLower(upper)
		data := string(filedata)

		data = data[start_pos+len(packageName):]

		data = strings.Replace(data, "int64", lower, -1)
		data = strings.Replace(data, "Int64", upper, -1)

		w.WriteString(data)
		w.WriteString("\r\n")
	}
	out, err := format.Source(w.Bytes())
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("types.go", out, 0660); err != nil {
		panic(err)
	}
}
