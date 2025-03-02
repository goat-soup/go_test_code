package example

import (
	"fmt"
	"time"
)

func SwitchFunc() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is a weekend")
	default:
		fmt.Println("Today is a workday")
	}

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("boolean")
		default:
			fmt.Println("unknown")
		}
	}

	whoAmI(false)
}
