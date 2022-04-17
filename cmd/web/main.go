package main

import (
	"context"
	"fmt"

	"github.com/dragun-igor/banner-rotation/internal/resources"
	"github.com/dragun-igor/banner-rotation/internal/rotator"
)

func main() {
	fmt.Println("res")
	res := resources.GetResources(context.Background())
	fmt.Println("rot")
	rot := rotator.NewRotator(res)
	fmt.Println("err")
	err := rot.AddBanner(1, 2)
	fmt.Println(err)
}
