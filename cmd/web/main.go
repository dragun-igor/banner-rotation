package main

import (
	"context"
	"fmt"

	"github.com/dragun-igor/banner-rotation/internal/resources"
	"github.com/dragun-igor/banner-rotation/internal/rotator"
)

func main() {
	res := resources.GetResources(context.Background())
	rot := rotator.NewRotator(res)
	err := rot.AddBannerToSlot(1, 2)
	fmt.Println(err)
}
