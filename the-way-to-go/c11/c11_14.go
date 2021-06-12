package c11

import "fmt"

type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

type Cars []*Car

// Process 函数用于 cars 遍历执行函数 f
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// 寻找所有符合条件（函数作为筛选器）的项
func (cs Cars) FindAll(f func(car *Car) bool) Cars {

	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})

	return cars
}

// 遍历并创建新数据（Any作为空接口已在 c11_9 中定义）
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})

	return result
}

// 根据入参返回不同的函数。用于产生特定的添加函数和 map 集
func MakeSortedAppender(manufactures []string) (func(car *Car), map[string]Cars) {
	// 为已排序的 cars 准备 maps
	sortedCars := make(map[string]Cars)
	for _, m := range manufactures {
		sortedCars[m] = make(Cars, 0)
	}
	sortedCars["Default"] = make(Cars, 0)

	// 准备 appender 函数
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}

	return appender, sortedCars
}

func C11_14() {
	// 数据初始化
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	// 简易 query
	allCars := Cars(Cars{ford, bmw, merc, bmw2})
	finderNewBMWs := func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	}
	allNewBMWs := allCars.FindAll(finderNewBMWs)
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)

	// 函数工厂
	manufactures := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufactures)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}
