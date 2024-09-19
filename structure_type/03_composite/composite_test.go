package composite

import "testing"

func TestComposite(t *testing.T) {
	// 创建员工
	e1 := &Employee{name: "Alice"}
	e2 := &Employee{name: "Bob"}
	e3 := &Employee{name: "Charlie"}

	// 创建部门
	department1 := &Department{name: "Engineering"}
	department2 := &Department{name: "HR"}

	// 将员工添加到部门中
	department1.AddChild(e1)
	department1.AddChild(e2)
	department2.AddChild(e3)

	// 创建公司大部门
	company := &Department{name: "Headquarters"}
	company.AddChild(department1)
	company.AddChild(department2)

	// 展示整个组织结构
	company.Display(0)
}
