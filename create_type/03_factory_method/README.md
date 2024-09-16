# 工厂方法

这里将OperatorBase 作为一个子类，先定义，之后再定义AddOperator、SubOperator 等子类。

工厂方法模式使用子类的方式延迟生成对象到子类中实现。
Go 中不存在继承 所以使用匿名组合来实现
