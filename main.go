package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Line []View

type View map[string]interface{}

func LoadJson(path string, dist interface{}) (err error) {
	var content []byte
	if content, err = ioutil.ReadFile(path); err == nil {
		err = json.Unmarshal(content, dist)
	}
	return err
}

func main() {
	var (
		line1 Line
	)
	f, _ := os.Create("cover.log")
	defer f.Close()
	if err := LoadJson("./default_cover.json", &line1); err != nil {
		fmt.Println(1, err)
		return
	}

	for _, v1 := range line1 {
		var g1 map[string]interface{}
		if gs1, ok := v1["groups"].([]interface{}); ok {
			if len(gs1) >= 1 {
				g1 = gs1[0].(map[string]interface{})

			}
		}
		n1 := make([]interface{}, 0)
		if ns1, ok := g1["nodes"].([]interface{}); ok {
			n1 = ns1
		}
		f.WriteString(v1["view"].(string) + ":\n")
		printNodes(n1, f)
	}
}

func printNodes(n1 []interface{}, f *os.File) {
	for _, n := range n1 {
		f.WriteString("    " + n.(string) + "\n")
	}
}
