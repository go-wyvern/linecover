package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// var views = []string{"上海电信", "上海移动", "上海联通", "云南电信", "云南移动", "云南联通", "内蒙古电信", "内蒙古移动", "内蒙古联通", "北京电信", "北京移动", "北京联通", "吉林电信", "吉林移动", "吉林联通", "四川电信", "四川移动", "四川联通", "天津电信", "天津移动", "天津联通", "宁夏电信", "宁夏移动", "宁夏联通", "安徽电信", "安徽移动", "安徽联通", "山东电信", "山东移动", "山东联通", "山西电信", "山西移动", "山西联通", "广东电信", "广东移动", "广东联通", "广西电信", "广西移动", "广西联通", "新疆电信", "新疆移动", "新疆联通", "江苏电信", "江苏移动", "江苏联通", "江西电信", "江西移动", "江西联通", "河北电信", "河北移动", "河北联通", "河南电信", "河南移动", "河南联通", "浙江电信", "浙江移动", "浙江联通", "海南电信", "海南移动", "海南联通", "湖北电信", "湖北移动", "湖北联通", "湖南电信", "湖南移动", "湖南联通", "甘肃电信", "甘肃移动", "甘肃联通", "福建电信", "福建移动", "福建联通", "西藏电信", "西藏移动", "西藏联通", "贵州电信", "贵州移动", "贵州联通", "辽宁电信", "辽宁移动", "辽宁联通", "重庆电信", "重庆移动", "重庆联通", "陕西电信", "陕西移动", "陕西联通", "青海电信", "青海移动", "青海联通", "黑龙江电信", "黑龙江移动", "黑龙江联通"}

var views = []string{"默认-上海-华东-联通-中国-亚洲"}

type Line []View

type View struct {
	ViewInfo string  `json:"viewInfo"`
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
		line1     Line
		lineCover = make(map[string][]string)
	)
	f, _ := os.Create("cover1.log")
	defer f.Close()
	if err := LoadJson("./quality_cover.json", &line1); err != nil {
		fmt.Println(1, err)
		return
	}

	for _, v1 := range line1 {
		if len(v1.Groups) > 0 {
			lineCover[v1.ViewInfo] = v1.Groups[0].Nodes
		} else {
			lineCover[v1.ViewInfo] = make([]string, 0)
		}
	}

	for _, v := range views {
		f.WriteString(v + ":")
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
