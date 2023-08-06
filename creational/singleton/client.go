package singleton

import "fmt"

func Singleton() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
