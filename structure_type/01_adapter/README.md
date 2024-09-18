# 适配器模式
https://refactoringguru.cn/design-patterns/adapter
转换一种接口，适配另一种接口,重点是解决不兼容问题

就像type类型手机的耳机，假设我们只有圆孔，那么我们就需要一个适配器Adapter，将圆孔转换为type孔。

Adaptee为被适配的接口

在 Adapter 中匿名组合 Adaptee 接口，所以 Adapter 类也拥有 SpecificRequest 实例方法，又因为 Go 语言中非入侵式接口特征，其实 Adapter 也适配 Adaptee 接口。