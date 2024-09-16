package prototype

import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	// 创建具体的原型对象
	proto1 := &ConcretePrototype1{field1: "Prototype1", field2: 42}
	proto2 := &ConcretePrototype2{data: []int{1, 2, 3, 4}}

	// 克隆对象,深拷贝
	clone1 := proto1.Clone()
	clone2 := proto2.Clone()

	// 输出原型对象和克隆对象
	fmt.Printf("Original proto1: %+v,addr:%p\n", proto1, &proto1)
	fmt.Printf("Cloned proto1: %+v,addr:%p\n", clone1, &clone1)

	fmt.Printf("Original proto2: %+v,addr:%p\n", proto2, &proto2)
	fmt.Printf("Cloned proto2: %+v,addr:%p\n", clone2, &clone2)
}
