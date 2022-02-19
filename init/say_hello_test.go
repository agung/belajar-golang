package init

import (
	"testing"
)

func sayHello() string {
	hello := Hello("Agung Nugraha")
	return hello
}

func TestSayHello(t *testing.T) {
	sayHello()
}
