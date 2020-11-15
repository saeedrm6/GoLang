package variables

import "fmt"

type floatsaeed float64

func Aliasing() {
	// type aliasName aliasTo
	var f floatsaeed = 1.24456
	fmt.Printf("f has value %v and type %T", f, f)


	// نکته : دو تایئ متفاوت رو نمیتوان در شرط ها بررسی کرد حتی اگر مشابه اش ساخته باشی
}
