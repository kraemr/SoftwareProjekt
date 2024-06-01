package moderator

import (
	"fmt"
)

/*
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
email TEXT,
password TEXT,
city TEXT,
username TEXT
*/

type Moderator struct {
	Id             int32
	Email          string
	Moderates_city string
	Username       string
}

var ErrNoModerator = fmt.Errorf("No Moderators Found")

func GetModeratorById(id int64) (Moderator, error) {
	_ = id
	return Moderator{}, ErrNoModerator
}

func GetModerators(city string) ([]Moderator, error) {
	_ = city
	return nil, ErrNoModerator
}

func UpdateModerator(moderator Moderator) error {
	_ = moderator
	return ErrNoModerator
}

func InsertModerator(moderator Moderator) ([]Moderator, error) {
	_ = moderator
	return nil, ErrNoModerator
}

func DeleteModerator(id int64) error {
	_ = id
	return ErrNoModerator
}
