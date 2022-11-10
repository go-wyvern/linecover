package main

import (
	"fmt"
	"testing"
)

func TestFixView(t *testing.T) {
	v := "默认-黑龙江-东北-联通-中国-亚洲"
	fmt.Println(fixView(v))
}
