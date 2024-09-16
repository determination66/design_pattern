package prototype

// Cloneable 是原型接口，定义了 Clone 方法
type Cloneable interface {
	Clone() Cloneable
}

// ConcretePrototype1 是第一个具体的原型
type ConcretePrototype1 struct {
	field1 string
	field2 int
}

// Clone 实现 Cloneable 接口，克隆对象
func (p *ConcretePrototype1) Clone() Cloneable {
	clone := *p // 创建一个当前对象的副本
	return &clone
}

// ConcretePrototype2 是第二个具体的原型
type ConcretePrototype2 struct {
	data []int
}

// Clone 实现 Cloneable 接口，克隆对象
func (p *ConcretePrototype2) Clone() Cloneable {
	clone := *p
	clone.data = make([]int, len(p.data)) // 深拷贝切片
	copy(clone.data, p.data)
	return &clone
}
