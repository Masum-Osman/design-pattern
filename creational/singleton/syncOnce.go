package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleOnce struct{}

var singleInstanceOnce *single

func getInstanceFromOnce() *single {
	if singleInstanceOnce == nil {
		once.Do(
			func() {
				fmt.Println("Creating a single instance now.")
				singleInstanceOnce = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstanceOnce
}
