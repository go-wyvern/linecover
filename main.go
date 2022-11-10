package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var views = []string{"默认-上海-华东-电信-中国-亚洲", "默认-上海-华东-移动-中国-亚洲", "默认-上海-华东-联通-中国-亚洲", "默认-云南-西南-电信-中国-亚洲", "默认-云南-西南-移动-中国-亚洲", "默认-云南-西南-联通-中国-亚洲", "默认-内蒙古-华北-电信-中国-亚洲", "默认-内蒙古-华北-移动-中国-亚洲", "默认-内蒙古-华北-联通-中国-亚洲", "默认-北京-华北-电信-中国-亚洲", "默认-北京-华北-移动-中国-亚洲", "默认-北京-华北-联通-中国-亚洲", "默认-吉林-东北-电信-中国-亚洲", "默认-吉林-东北-移动-中国-亚洲", "默认-吉林-东北-联通-中国-亚洲", "默认-四川-西南-电信-中国-亚洲", "默认-四川-西南-移动-中国-亚洲", "默认-四川-西南-联通-中国-亚洲", "默认-天津-华北-电信-中国-亚洲", "默认-天津-华北-移动-中国-亚洲", "默认-天津-华北-联通-中国-亚洲", "默认-宁夏-西北-电信-中国-亚洲", "默认-宁夏-西北-移动-中国-亚洲", "默认-宁夏-西北-联通-中国-亚洲", "默认-安徽-华东-电信-中国-亚洲", "默认-安徽-华东-移动-中国-亚洲", "默认-安徽-华东-联通-中国-亚洲", "默认-山东-华东-电信-中国-亚洲", "默认-山东-华东-移动-中国-亚洲", "默认-山东-华东-联通-中国-亚洲", "默认-山西-华北-电信-中国-亚洲", "默认-山西-华北-移动-中国-亚洲", "默认-山西-华北-联通-中国-亚洲", "默认-广东-华南-电信-中国-亚洲", "默认-广东-华南-移动-中国-亚洲", "默认-广东-华南-联通-中国-亚洲", "默认-广西-华南-电信-中国-亚洲", "默认-广西-华南-移动-中国-亚洲", "默认-广西-华南-联通-中国-亚洲", "默认-新疆-西北-电信-中国-亚洲", "默认-新疆-西北-移动-中国-亚洲", "默认-新疆-西北-联通-中国-亚洲", "默认-江苏-华东-电信-中国-亚洲", "默认-江苏-华东-移动-中国-亚洲", "默认-江苏-华东-联通-中国-亚洲", "默认-江西-华东-电信-中国-亚洲", "默认-江西-华东-移动-中国-亚洲", "默认-江西-华东-联通-中国-亚洲", "默认-河北-华北-电信-中国-亚洲", "默认-河北-华北-移动-中国-亚洲", "默认-河北-华北-联通-中国-亚洲", "默认-河南-华中-电信-中国-亚洲", "默认-河南-华中-移动-中国-亚洲", "默认-河南-华中-联通-中国-亚洲", "默认-浙江-华东-电信-中国-亚洲", "默认-浙江-华东-移动-中国-亚洲", "默认-浙江-华东-联通-中国-亚洲", "默认-海南-华南-电信-中国-亚洲", "默认-海南-华南-移动-中国-亚洲", "默认-海南-华南-联通-中国-亚洲", "默认-湖北-华中-电信-中国-亚洲", "默认-湖北-华中-移动-中国-亚洲", "默认-湖北-华中-联通-中国-亚洲", "默认-湖南-华中-电信-中国-亚洲", "默认-湖南-华中-移动-中国-亚洲", "默认-湖南-华中-联通-中国-亚洲", "默认-甘肃-西北-电信-中国-亚洲", "默认-甘肃-西北-移动-中国-亚洲", "默认-甘肃-西北-联通-中国-亚洲", "默认-福建-华东-电信-中国-亚洲", "默认-福建-华东-移动-中国-亚洲", "默认-福建-华东-联通-中国-亚洲", "默认-西藏-西南-电信-中国-亚洲", "默认-西藏-西南-移动-中国-亚洲", "默认-西藏-西南-联通-中国-亚洲", "默认-贵州-西南-电信-中国-亚洲", "默认-贵州-西南-移动-中国-亚洲", "默认-贵州-西南-联通-中国-亚洲", "默认-辽宁-东北-电信-中国-亚洲", "默认-辽宁-东北-移动-中国-亚洲", "默认-辽宁-东北-联通-中国-亚洲", "默认-重庆-西南-电信-中国-亚洲", "默认-重庆-西南-移动-中国-亚洲", "默认-重庆-西南-联通-中国-亚洲", "默认-陕西-西北-电信-中国-亚洲", "默认-陕西-西北-移动-中国-亚洲", "默认-陕西-西北-联通-中国-亚洲", "默认-青海-西北-电信-中国-亚洲", "默认-青海-西北-移动-中国-亚洲", "默认-青海-西北-联通-中国-亚洲", "默认-黑龙江-东北-电信-中国-亚洲", "默认-黑龙江-东北-移动-中国-亚洲", "默认-黑龙江-东北-联通-中国-亚洲"}

type LineCover struct {
	Views []View `json:"line_covers"`
}

type View struct {
	ViewInfo string  `json:"view_info"`
	Groups   []Group `json:"groups"`
}

type Group struct {
	Nodes []string `json:"nodes"`
}

func LoadJson(path string, dist interface{}) (err error) {
	var content []byte
	if content, err = ioutil.ReadFile(path); err == nil {
		err = json.Unmarshal(content, dist)
	}
	return err
}

func main() {
	var (
		line1     LineCover
		lineCover = make(map[string][]string)
	)
	f, _ := os.Create("cover.log")
	defer f.Close()
	if err := LoadJson("./quality_cover.json", &line1); err != nil {
		fmt.Println(1, err)
		return
	}

	for _, v1 := range line1.Views {
		if len(v1.Groups) > 0 {
			lineCover[v1.ViewInfo] = v1.Groups[0].Nodes
		} else {
			lineCover[v1.ViewInfo] = make([]string, 0)
		}
	}

	for _, v := range views {
		f.WriteString(v + ":" + "\n")
		printNodes(findViewNodes(v, lineCover), f)
	}
}

func findViewNodes(v string, lineCover map[string][]string) []string {
	if n, ok := lineCover[v]; ok {
		return n
	}
	v = fixView(v)
	return findViewNodes(v, lineCover)
}

func fixView(v string) string {
	vv := strings.Split(v, "-")
	var v2 []string
	var first bool
	for _, t := range vv {
		if t == "默认" {
			v2 = append(v2, t)
		} else {
			if !first {
				v2 = append(v2, "默认")
				first = true
			} else {
				v2 = append(v2, t)
			}
		}
	}
	return strings.Join(v2, "-")
}

func printNodes(n1 []string, f *os.File) {
	for _, n := range n1 {
		f.WriteString("    " + n + "\n")
	}
}
