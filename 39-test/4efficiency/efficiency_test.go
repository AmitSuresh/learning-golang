package efficiency

import (
	"fmt"
	"strings"
	"testing"
)

func TestInefficientWay(t *testing.T) {
	str := "Johnny Johnny Yes Papa"
	slices := strings.Split(str, " ")
	str = InefficientWay(slices)
	if str != "Johnny Johnny Yes Papa" {
		t.Errorf("actual: %v, expected: %v", str, "Johnny Johnny Yes Papa")
	}
}

func TestEfficientWay(t *testing.T) {
	str := "Johnny Johnny Yes Papa"
	slices := strings.Split(str, " ")
	str = EfficientWay(slices)
	if str != "Johnny Johnny Yes Papa" {
		t.Errorf("actual: %v, expected: %v", str, "Johnny Johnny Yes Papa")
	}
}

func ExampleEfficientWay() {
	str := "Johnny Johnny Yes Papa"
	slices := strings.Split(str, " ")
	str = InefficientWay(slices)
	fmt.Println(str)
	//Output:
	//Johnny Johnny Yes Papa
}
func ExampleInefficientWay() {
	str := "Johnny Johnny Yes Papa"
	slices := strings.Split(str, " ")
	str = EfficientWay(slices)
	fmt.Println(str)
	//Output:
	//Johnny Johnny Yes Papa
}

const str = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

func BenchmarkInefficientway(b *testing.B) {
	slices := strings.Split(str, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InefficientWay(slices)
	}
}

func BenchmarkEfficientway(b *testing.B) {
	slices := strings.Split(str, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EfficientWay(slices)
	}
}
