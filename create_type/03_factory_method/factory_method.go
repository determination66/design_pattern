package factorymethod

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

// SetA 设置 A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// SetB 设置 B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperator Operator 的实际加法实现
type AddOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o AddOperator) Result() int {
	return o.a + o.b
}

// AddOperatorFactory 是 AddOperator 的工厂类
type AddOperatorFactory struct{}

func (a *AddOperatorFactory) Create() Operator {
	return &AddOperator{
		OperatorBase: &OperatorBase{},
	}
}

var _ OperatorFactory = &AddOperatorFactory{}

// SubOperatorFactory 是 SubOperator 的工厂类
type SubOperatorFactory struct{}

func (SubOperatorFactory) Create() Operator {
	return &SubOperator{
		OperatorBase: &OperatorBase{},
	}
}

var _ OperatorFactory = &SubOperatorFactory{}

// SubOperator Operator 的实际减法实现
type SubOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o SubOperator) Result() int {
	return o.a - o.b
}
