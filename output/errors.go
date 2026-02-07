// Package output отвечает за вывод ошибок
package output

import (
	"github.com/fatih/color"
)

// PrintErrorSwitch could be useful in simple situatuins
func PrintErrorSwitch(value any) {
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Exit code: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестная ошибка:", t)
	}
}

func sum[T int | float64 | float32 | int16 | int32 | string](a, b T) T {
	return a + b
}

// PrintErrorIf could be useful in difficult situations
// func PrintErrorIf(value any) {
// 	if intValue, ok := value.(int); ok {
// 		color.Red("Exit code: %d", intValue)
// 		return
// 	}
// 	if strValue, ok := value.(string); ok {
// 		color.Red(strValue)
// 		return
// 	}
// 	if errorValue, ok := value.(error); ok {
// 		color.Red(errorValue.Error())
// 		return
// 	}
// 	color.Red("Неизвестная ошибка:", value)
// }
