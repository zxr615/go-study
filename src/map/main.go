package main

func main()  {

	person := make(map[string]int)
	person["age"] = 18
	person["stature"] = 170
	person["vision"] = 5

	vision, exist := person["vision"]

	if exist {
		println("存在 vision = ", vision)
	} else {
		println("不存在 vision ")
	}

	println("age = ", person["age"])
	println("stature = ", person["stature"])


}