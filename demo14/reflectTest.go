package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

/**
Go 语言的作者在博客上总结了反射的三大定律。
1. 任何接口值 interface{} 都可以反射出反射对象，也就是 reflect.Value 和 reflect.Type，通过函数 reflect.ValueOf 和 reflect.TypeOf 获得。
2. 反射对象也可以还原为 interface{} 变量，也就是第 1 条定律的可逆性，通过 reflect.Value 结构体的 Interface 方法获得。
3. 要修改反射的对象，该值必须可设置，也就是可寻址，参考上节课修改变量的值那一节的内容理解。
小提示： 任何类型的变量都可以转换为空接口 interface{}，所以第 1 条定律中函数 reflect.ValueOf 和 reflect.TypeOf 的参数就是 interface{}，
       表示可以把任何类型的变量转换为反射对象。
       在第 2 条定律中，reflect.Value 结构体的 Interface 方法返回的值也是 interface{}，表示可以把反射对象还原为对应的类型变量。
*/
func main() {
	//1.reflect.Value 和 reflect.Type
	reflectValueAndType()

	//2.reflect.Value 和 int 类型互转
	reflectValueChange()

	//3.修改对应的值
	reflectModifyValue()

	//4.修改结构体的值(字段需要是共有的，即大写的)
	reflectModifyValueForStruct()

	//5.获取底层类型
	getUnderlyingType()

	//6.获取方法和字段（FieldByName，MethodByName获取指定字段，方法）
	getMethodAndField()

	//7.判断是否实现了接口
	implementsTest()

	//8.struct和json互相转换
	structConverseJson()

	//9.struct Tag的使用
	structTag()

	//自己实现struct转json
	structTransferToJson()
}

//自己实现struct转json
func structTransferToJson() {
	p := person1{Name: "飞雪无情", Age: 20}
	pv := reflect.ValueOf(p)
	pt := reflect.TypeOf(p)
	//自己实现的struct to json
	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")
	num := pt.NumField()
	for i := 0; i < num; i++ {
		jsonTag := pt.Field(i).Tag.Get("json") //获取json tag
		jsonBuilder.WriteString("\"" + jsonTag + "\"")
		jsonBuilder.WriteString(":")
		//获取字段的值
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", pv.Field(i)))
		if i < num-1 {
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String()) //打印json字符串
}

func structTag() {
	p := person1{Name: "飞雪无情", Age: 20}
	//struct to json
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}
	//json to struct
	respJSON := "{\"Name\":\"李四\",\"Age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(p)

	//遍历person字段中key为json、bson的tag
	pt := reflect.TypeOf(p)
	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n", sf.Name, sf.Tag.Get("json"))
		fmt.Printf("字段%s上,bson tag为%s\n", sf.Name, sf.Tag.Get("bson"))
	}
}

//8.struct和json互相转换
func structConverseJson() {
	p := person{Name: "飞雪无情", Age: 20}
	//struct to json
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}
	//json to struct
	respJSON := "{\"Name\":\"李四\",\"Age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(p)
}

//7.判断是否实现了接口
func implementsTest() {
	p := person{Name: "飞雪无情", Age: 20}
	pt := reflect.TypeOf(p)
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer：", pt.Implements(stringerType))
	fmt.Println("是否实现了io.Writer：", pt.Implements(writerType))
}

//6.获取方法和字段（FieldByName，MethodByName获取指定字段，方法）
func getMethodAndField() {
	p := person{Name: "飞雪无情", Age: 20}
	pt := reflect.TypeOf(p)
	//遍历person的字段
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("字段：", pt.Field(i).Name)
	}

	//遍历person的方法
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("方法：", pt.Method(i).Name)
	}
}

//5.获取底层类型
func getUnderlyingType() {
	p := person{Name: "飞雪无情", Age: 20}
	ppv := reflect.ValueOf(&p)
	fmt.Println(ppv.Kind())
	pv := reflect.ValueOf(p)
	fmt.Println(pv.Kind())
}

//4.修改结构体的值(字段需要是共有的，即大写的)
func reflectModifyValueForStruct() {
	p := person{Name: "飞雪无情", Age: 20}
	ppv := reflect.ValueOf(&p)
	ppv.Elem().Field(0).SetString("张三")
	fmt.Println(p)
}

//3.修改对应的值
func reflectModifyValue() {
	i := 3
	//ValueOf传入的仍然是值的拷贝，因此要换成指针
	ipv := reflect.ValueOf(&i)
	ipv.Elem().SetInt(4)
	fmt.Println(i)
}

//2.reflect.Value 和 int 类型互转
func reflectValueChange() {
	i := 3
	//int to reflect.Value
	iv := reflect.ValueOf(i)
	//reflect.Value to int
	i1 := iv.Interface().(int)
	fmt.Println(i1)
}

//1.reflect.Value 和 reflect.Type
func reflectValueAndType() (int, reflect.Value) {
	i := 3
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	fmt.Println(iv, it) //3 int
	return i, iv
}

type person struct {
	Name string
	Age  int
}

//实现 fmt.Stringer 接口
func (p person) String() string {
	return fmt.Sprintf("Name is %s,Age is %d", p.Name, p.Age)
}

//9.struct tag 的使用
type person1 struct {
	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age" bson:"b_name"`
}
