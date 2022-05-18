package main

import (
	"os"
)

// func main() {
// 	ctx := context.Background()
// 	res := resources.GetResources(ctx)

// 	rr := createRotator(res)
// 	initServer(ctx, rr, res)
// }

// func initServer(ctx context.Context, r *rotator.Rotator, res *resources.Resources) {
// 	s := server.NewServer(
// 		r,
// 		res,
// 	)
// 	log.Fatal(s.Run(ctx))
// }

// func createRotator(res *resources.Resources) *rotator.Rotator {
// 	return rotator.NewRotator(res)
// }

func main() {
	path := "~/golang/banner-rotation/migrations/createtables.up.sql"
	c, ioErr := os.ReadFile(path)
	_, err := *pgx.Conn.Exec(string(c))
}
