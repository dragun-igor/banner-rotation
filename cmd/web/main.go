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
	err = rot.RemoveBannerFromSlot(1, 2)
	fmt.Println(err)
	err = rot.Showed(1, 1, 1)
	fmt.Println(err)
	err = rot.Clicked(1, 1, 1)
	fmt.Println(err)
	rot.SelectBanner(1, 1)
}
