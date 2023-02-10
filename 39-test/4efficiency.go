package main

import (
	"fmt"
	"strings"

	efficiency "github.com/AmitSuresh/learning-golang/39-test/4efficiency"
)

const str = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

func main() {
	slices := strings.Split(str, " ")

	for _, v := range slices {
		fmt.Println(v)
	}

	fmt.Println(efficiency.InefficientWay(slices))
	fmt.Println(efficiency.EfficientWay(slices))

}
