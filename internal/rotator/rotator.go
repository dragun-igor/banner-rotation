package rotator

import (
	"fmt"

	"github.com/dragun-igor/banner-rotation/internal/resources"
)

type Rotator struct {
	res *resources.Resources
}

func NewRotator(res *resources.Resources) *Rotator {
	return &Rotator{
		res: res,
	}
}

type ResultQuery struct {
	ID          int
	Description string
}

func (r Rotator) AddBanner(bannerID int, slotID int) error {
	var rq ResultQuery
	query := `
	SELECT *
	FROM banners
	WHERE id=$1;`
	res, _ := r.res.DB.Query(query, bannerID)
	for res.Next() {
		_ = res.Scan(&rq.ID, &rq.Description)
		fmt.Println(rq)
	}

	query = `
	INSERT INTO slots (id, description)
	VALUES (1, $1)
	RETURNING id;`
	_ = r.res.DB.QueryRow(query, rq.Description)

	return nil
}
