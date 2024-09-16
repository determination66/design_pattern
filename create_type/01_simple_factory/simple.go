package simplefactory

import (
	"fmt"
)

// API is interface
type API interface {
	Say(name string) string
}

// hiAPI is one of API implement
type hiAPI struct{}

// Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// helloAPI is another API implement
type helloAPI struct{}

// Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

// // NewAPI version1：简单写法，工厂函数,保证可读性
// func NewAPI(t string) (API, error) {
// 	switch t {
// 	case "hi":
// 		return &hiAPI{}, nil
// 	case "hello":
// 		return &helloAPI{}, nil
// 	default:
// 		return nil, fmt.Errorf("unknown API type: %s", t)
// 	}
// }

// 以上是一个简单的工厂模式，利用多态，实现接口
// 后续若是实例过多，那么可以继续解耦合，将newAPI与创建逻辑分离
var apiMap = map[string]func() API{}

func RegisterAPI(name string, factory func() API) {
	apiMap[name] = factory
}

func NewAPI(t string) (API, error) {
	if f, ok := apiMap[t]; ok {
		return f(), nil
	}
	return nil, fmt.Errorf("unknown API type: %s", t)
}

func init() {
	RegisterAPI("hi", func() API { return &hiAPI{} })
	RegisterAPI("hello", func() API { return &helloAPI{} })
}
