package userCenter

import (
	"errors"
	"github.com/firerainos/firerain-web-go/core"
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);unique"`
	Description string
}

func AddGroup(name, description string) error {
	db, err := core.GetSqlConn()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Create(&Group{Name: name, Description: description}).Error
}

func GetGroup() ([]Group, error) {
	db, err := core.GetSqlConn()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var groups []Group
	db.Find(&groups)

	return groups, nil
}

func GetGroupByName(name string) (Group, error) {
	var group Group

	db, err := core.GetSqlConn()
	if err != nil {
		return group, err
	}
	defer db.Close()

	if db.Where("name = ?", name).First(&group).RecordNotFound() {
		return group, errors.New("group " + name + " not found")
	}

	return group, nil
}

func GetGroupByNames(names []string) ([]Group, error) {
	var groups []Group

	db, err := core.GetSqlConn()
	if err != nil {
		return groups, err
	}
	defer db.Close()

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
	db, err := core.GetSqlConn()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Where("name = ?", name).Delete(&Group{}).Error
}
