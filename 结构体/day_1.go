package main

// 声明结构体,关键字;type
// 类型：更多的像是一个模板         对象：基于模板构建出来的东西叫对象
// 类型包含属性和方法
// 对象是类型实例化的结果，一个类型可以实例化出无数个对象
type Student struct { //结构体是值类型，不利用地址去转存
	sid    int //Student的成员对象
	t1, t2 int
	name   string
	age    int
	course []string
	tt     map[string]string
}

func main() {
	//在实际开发中，我们可以将一组类型不同的、但是用来描述同一件事物的变量放到结构体中。
	//学了结构体之后，我们就不需要定义多个变量了，将他们都放到结构体中即可
	//结构体起着类的作用

}
