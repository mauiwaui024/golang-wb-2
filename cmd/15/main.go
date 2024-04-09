// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

var justString string

func someFunc() {
	// v := createHugeString(1 << 10)
	v := createHugeString(32) //делаем пример попроще и смотрим по памяти че не так
	justString = string(v[:10])
	runtime.KeepAlive(v)

}

func main() {
	someFunc()
	justStringBytes := []byte(justString)
	for i, element := range justStringBytes {

		fmt.Printf("justString[%d] - Address: %p, Size: %d bytes, Value: %d\n", i, &justStringBytes[i], unsafe.Sizeof(element), element)

	}
}

func createHugeString(size int) string {
	hugeString := make([]byte, size)
	for i, element := range hugeString {
		if i < 10 {
			fmt.Printf("hugeString[%d] - Address: %p, Size: %d bytes, Value: %d\n", i, &hugeString[i], unsafe.Sizeof(element), element)
		}

	}
	return string(hugeString)
}

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = string(v[:100])
// }

// func createHugeString(size int) []byte {
// 	// Здесь реализация функции createHugeString
// 	// Возвращается большая строка
// }

// func main() {
// 	someFunc()
// }
