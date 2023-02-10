package speak

import (
	"fmt"
	"testing"
)

func TestSpeak(t *testing.T) {
	str := Hello("Johnny")
	if str != "Greetings..Johnny" {
		t.Errorf("got: %v, expected: %v", str, "Greetings..Johnny")
	}
}

func ExampleHello_first() {
	fmt.Println(Hello("Johnny"))
	//Output:
	//Greetings..Johnny
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Johnny")
	}
}
