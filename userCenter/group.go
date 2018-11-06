package userCenter

import (
	"errors"
	"github.com/firerainos/firerain-web-go/database"
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);unique"`
	Description string
}

func AddGroup(name, description string) error {
	db := database.Instance()

	return db.Create(&Group{Name: name, Description: description}).Error
}

func GetGroup() ([]Group, error) {
	db := database.Instance()

	var groups []Group
	db.Find(&groups)

	return groups, nil
}

func GetGroupByName(name string) (Group, error) {
	var group Group

	db := database.Instance()

	if db.Where("name = ?", name).First(&group).RecordNotFound() {
		return group, errors.New("group " + name + " not found")
	}

	return group, nil
}

func GetGroupByNames(names []string) ([]Group, error) {
	var groups []Group

	db := database.Instance()

	for _, name := range names {
		group := Group{}
		if db.Where("name = ?", name).First(&group).RecordNotFound() {
			return groups, errors.New("group " + name + " not found")
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func DeleteGroup(name string) error {
	db := database.Instance()

	return db.Where("name = ?", name).Delete(&Group{}).Error
}
