package validate

// 以slice示例
func SliceValidate() {
	v := validate
	var err error
	slice1 := []string{"11", "1", "12", "22"}
	err = v.Var(slice1, "gte=3,dive,gte=1,number")
	outRes("slice1", &err)

	slice2 := [][]string{
		{"12345", "67890", "1234567890"},
		{"12345", "67890", "1234567890"},
		{"12345", "67890", "1234567890"},
	}
	err = v.Var(slice2, "gte=3,dive,gte=3,dive,gte=5,lte=10,number") // 需要dive两层
	outRes("slice2", &err)
}
