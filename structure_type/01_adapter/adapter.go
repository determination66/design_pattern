package adapter

import "fmt"

// Target 是手机支持的接口 (Type-C 或 Lightning)
type Target interface {
	ConnectWithPhone() string
}

// Adaptee 是传统的耳机接口 (3.5mm 圆孔)
type Adaptee interface {
	ConnectWithHeadphoneJack() string
}

// Earphone 是 Adaptee 的实现，即 3.5mm 圆孔耳机
type Earphone struct{}

// ConnectWithHeadphoneJack 实现 Adaptee 接口
func (e *Earphone) ConnectWithHeadphoneJack() string {
	return "Connected with 3.5mm headphone jack"
}

// Adapter 是适配器类，用于将 3.5mm 圆孔耳机转换为 Type-C/Lightning
type Adapter struct {
	earphone Adaptee
}

// NewAdapter 创建 Adapter 实例
func NewAdapter(earphone Adaptee) Target {
	return &Adapter{
		earphone: earphone,
	}
}

// ConnectWithPhone 实现 Target 接口，将耳机转换为 Type-C/Lightning
func (a *Adapter) ConnectWithPhone() string {
	// 使用适配器将耳机接口转换为 Type-C/Lightning
	return "Adapter converts: " + a.earphone.ConnectWithHeadphoneJack() + " to Type-C/Lightning"
}

func main() {
	// 创建一个 3.5mm 的耳机
	earphone := &Earphone{}

	// 使用适配器将 3.5mm 耳机连接到 Type-C/Lightning 手机
	adapter := NewAdapter(earphone)

	// 模拟连接
	fmt.Println(adapter.ConnectWithPhone())
}
