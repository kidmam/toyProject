package singleton

import (
	"sync"
)

type singleton struct{}

var instance *singleton

var once sync.Once

func GetInstance() *singleton {
	//     if instance == nil {
	//         instance = &singleton{}
	//     }
	// The above code is not thread safe
	// We can use `Once.Do` function of "sync" package to make it thread safe.
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
