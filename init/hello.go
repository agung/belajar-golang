package init

import "fmt"

func Hello(st string) string {
	fmt.Println("Open")
	defer func() {
		fmt.Println("Closed")
	}()

	local := st
	fmt.Println(local)

	return local
}
