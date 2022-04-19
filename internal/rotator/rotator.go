package rotator

import (
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

func (r Rotator) AddBannerToSlot(slotID int, bannerID int) error {
	query := "INSERT INTO rotation VALUES ($1, $2);"
	row := r.res.DB.QueryRow(query, slotID, bannerID)
	return row.Err()
}

func (r Rotator) RemoveBannerFromSlot(slotID int, bannerID int) error {
	query := "DELETE FROM rotation WHERE slot_id=$1 AND banner_id=$2;"
	row := r.res.DB.QueryRow(query, slotID, bannerID)
	return row.Err()
}

func (r Rotator) Showed(slotID int, bannerID int, userGroupID int) error {
	var id int
	query := "INSERT INTO stat (slot_id, banner_id, user_group_id) VALUES ($1, $2, $3);"
	_ = r.res.DB.QueryRow(query, slotID, bannerID).Scan(&id)
	query = "UPDATE stat SET show = show + 1 WHERE slot_id=$1 AND banner_id=$2 AND user_group_id=$3;"
	row := r.res.DB.QueryRow(query, slotID, bannerID, userGroupID)
	return row.Err()
}

func (r Rotator) Clicked(slotID int, bannerID int, userGroupID int) error {
	var id int
	query := "INSERT INTO stat (slot_id, banner_id, user_group_id) VALUES ($1, $2, $3);"
	_ = r.res.DB.QueryRow(query, slotID, bannerID).Scan(&id)
	query = "UPDATE stat SET click = click + 1 WHERE slot_id=$1 AND banner_id=$2 AND user_group_id=$3;"
	row := r.res.DB.QueryRow(query, slotID, bannerID, userGroupID)
	return row.Err()
}
