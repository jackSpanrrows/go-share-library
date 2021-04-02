package errhanding

import "fmt"

func TryRecover() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Nothing to recover. Please type uncomment errors below.")
			return
		}
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do : %v", r))
		}
	}()
}
