package app

func area(length, width float64)float64{
	area := length * width
	return area
}

func Sumarea(length, width, height float64)float64{
	if length < 5{
		sumarea := (length * width * height)/2
		return sumarea
	}else{
		sumarea2 := area(8,8) * height
		return sumarea2
	}
	//sumarea := (length * width * height)/2
	//sumarea2 := area(8,8) * height
}