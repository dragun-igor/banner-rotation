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
	query := "INSERT INTO slots_banners VALUES ($1, $2)"
	row := r.res.DB.QueryRow(query, slotID, bannerID)
	return row.Err()
}

func (r Rotator) RemoveBannerFromSlot(slotID int, bannerID int) error {
	query := "DELETE FROM slots_banners WHERE slot_id=$1 AND banner_id=$2"
	row := r.res.DB.QueryRow(query, slotID, bannerID)
	return row.Err()
}

func (r Rotator) Click(slotID int, bannerID int, userGroupID int) error {
	var id int
	query := "SELECT id FROM slots_banners WHERE slot_id=$1 AND banner_id=$2"
	err := r.res.DB.QueryRow(query, slotID, bannerID).Scan(&id)
	if err != nil {
		return err
	}
	query = "INSERT INTO"
	return nil
}
