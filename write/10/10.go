package main

import "fmt"

func groupDec(arr []float32) map[int][]float32 {
	// для группировки будем использовать map, где ключем будет соответствующий шаг, а значением -
	//слайс, удовлетворяющий соответствующему шагу
	temp := make(map[int][]float32)
	var c int

	// в цикле преводим температурное колебание к типу int и обрабатываем для использования в качестве ключа.
	//само температурное колебание добавляем в слайс.
	for _, val := range arr {
		c = int(val) / 10 * 10
		temp[c] = append(temp[c], val)
	}
	return temp

}

func main() {
	// создаем слайс с температурными коллебаниями
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//выводим полученный результат в консоль
	fmt.Println(groupDec(arr))
}
