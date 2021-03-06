package rotator

import (
	"fmt"
	"log"

	"github.com/dragun-igor/banner-rotation/internal/resources"
	"github.com/dragun-igor/banner-rotation/pkg/ucb1"
	"github.com/lib/pq"
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

func (r Rotator) AddBanner(description string) error {
	query := "INSERT INTO banners VALUES ($1);"
	row := r.res.DB.QueryRow(query, description)
	return row.Err()
}

func (r Rotator) AddSlot(description string) error {
	query := "INSERT INTO slots VALUES ($1);"
	row := r.res.DB.QueryRow(query, description)
	return row.Err()
}

func (r Rotator) AddUserGroup(description string) error {
	query := "INSERT INTO user_groups VALUES ($1);"
	row := r.res.DB.QueryRow(query, description)
	return row.Err()
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

type Stat struct {
	slotID      int
	usergroupID int
}

func (r Rotator) SelectBanner(slotID int, userGroup int) (int, error) {
	var bannerID int
	var banners []int
	var arms ucb1.Arms
	var arm ucb1.Arm
	query := "SELECT banner_id FROM rotation WHERE slot_id=$1;"
	rows, err := r.res.DB.Query(query, slotID)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&bannerID); err != nil {
			log.Fatal(err)
		}
		banners = append(banners, bannerID)
	}
	fmt.Println(banners)
	query = "SELECT * FROM stat WHERE banner_id=ANY($1);"
	rows, err = r.res.DB.Query(query, pq.Array(banners))
	stats := make(map[int]*Stat)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		stat := &Stat{}
		rows.Scan(&stat.slotID, &bannerID, &stat.usergroupID, &arm.Reward, &arm.Count)
		stats[bannerID] = stat
		arms = append(arms, arm)
	}
	fmt.Println(*stats[1])
	fmt.Println(arms)

	i := ucb1.UCB1(arms)
	fmt.Println(i, arms[i])

	return banners[i], nil
}
