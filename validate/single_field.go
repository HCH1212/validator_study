package validate

import "time"

// bool
// 数字
// 字符串
// 切片
// 集合
// 时间

func SingleFieldValidate() {
	v := validate
	var err error

	var b bool
	err = v.Var(b, "boolean")
	outRes("boolean", &err)

	var i = "100"
	err = v.Var(i, "numeric") // number不能校验 数字字符串 中的小数点
	outRes("numeric", &err)

	var str = "helloworld"
	err = v.Var(str, "alpha") // 仅字母
	outRes("alpha", &err)

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	err = v.Var(slice, "max=15,min=0")
	outRes("slice", &err)

	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	err = v.Var(m, "max=15,min=0")
	outRes("map", &err)

	var timeStr = time.Now().Format("2006-01-02 15:04:05")
	err = v.Var(timeStr, "datetime=2006-01-02 15:04:05")
	outRes("datetime", &err)

	s1 := "abc"
	s2 := "abc"
	err = v.VarWithValue(s1, s2, "eqfield")
	outRes("eqfield", &err)
}
