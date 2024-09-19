package composite

import (
	"fmt"
	"strings"
)

// Component 是组合模式中的通用接口，表示可以统一管理的对象
type Component interface {
	Display(indent int)
}

// Employee 是 Leaf 结构，表示一个普通员工
type Employee struct {
	name string
}

// Display 用于展示员工的名称，表示叶节点
func (e *Employee) Display(indent int) {
	fmt.Printf("%s- %s (Employee)\n", generateIndent(indent), e.name)
}

// Department 是 Composite 结构，表示可以包含子部门和员工的部门
type Department struct {
	name      string
	subGroups []Component // 包含子部门和员工
}

// AddChild 添加子部门或员工
func (d *Department) AddChild(child Component) {
	d.subGroups = append(d.subGroups, child)
}

// Display 用于展示部门和其子组件，表示复合节点
func (d *Department) Display(indent int) {
	fmt.Printf("%s+ %s (Department)\n", generateIndent(indent), d.name)
	for _, child := range d.subGroups {
		child.Display(indent + 2) // 子组件增加缩进
	}
}

// 生成缩进字符串
func generateIndent(indent int) string {
	return strings.Repeat(" ", indent)
}
