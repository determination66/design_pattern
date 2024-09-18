package adapter

import (
	"fmt"
	"testing"
)

var expect = "adaptee method"

func TestAdapter(t *testing.T) {
	// 创建一个 3.5mm 的耳机
	earphone := &Earphone{}
	// 使用适配器将 3.5mm 耳机连接到 Type-C/Lightning 手机
	adapter := NewAdapter(earphone)

	// 模拟连接
	fmt.Println(adapter.ConnectWithPhone())
}
