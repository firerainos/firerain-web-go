package userCenter

import (
	"github.com/jinzhu/gorm"
	"github.com/firerainos/firerain-web-go/core"
	"os"
)

type User struct {
	gorm.Model
	Nickname string
	Username string `gorm:"type:varchar(100);unique"`
	Password string
	Email string `gorm:"type:varchar(100);unique"`
	Group []Group `gorm:"many2many:user_group"`
}

func AddUser(nickname,username,password,email string,group []string) error {
	g,err := GetGroupByNames(group)
	if err != nil {
		return err
	}

	db,err := core.GetSqlConn()
	if err != nil {
		return err
	}
	defer db.Close()


	user := User{Nickname:nickname,Username:username,Password:password,Email:email,Group:g}

	user.Password = EncryptionPassword(user.Username,user.Password,user.Email)

	err = db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUser() ([]User,error) {
	var users []User

	db,err := core.GetSqlConn()
	if err != nil {
		return users,err
	}
	defer db.Close()

	db.Preload("Group").Find(&users)

	return users, nil
}

func GetUserByName(name string) (User,error) {
	var user User

	db,err := core.GetSqlConn()
	if err != nil {
		return user,err
	}
	defer db.Close()

	if db.Where("username = ?",name).Preload("Group").First(&user).RecordNotFound() {
		return user, db.Error
	}

	return user, nil
}

func GetUserById(id int) (User,error) {
	var user User

	db,err := core.GetSqlConn()
	if err != nil {
		return user,err
	}
	defer db.Close()

	if db.Where("id = ?",id).Preload("Group").First(&user).RecordNotFound() {
		return user, db.Error
	}

	return user, nil
}

func (user User) Delete() error {
	db,err := core.GetSqlConn()
	if err != nil {
		return err
	}
	defer db.Close()

	db.Delete(user)
	os.Remove("./assets/avatar/"+user.Username)

	return nil
}

func (user User) AddGroup(group string) error {
	g,err := GetGroupByName(group)
	if err != nil {
		return err
	}

	user.Group = append(user.Group, g)

	db,err := core.GetSqlConn()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Save(user).Error
}

func (user User) Edit(nickName,password string) error {
	db,err := core.GetSqlConn()
	if err != nil {
		return err
	}
	defer db.Close()

	tmp := User{}

	if nickName != "" {
		tmp.Nickname = nickName
	}

	if password != "" {
		tmp.Password = EncryptionPassword(user.Username,password,user.Email)
	}

	return db.Model(&user).Update(tmp).Error
}

func (user User) DeleteGroup(group string) error {
	var groups []Group
	for _,g := range user.Group {
		if g.Name != group {
			groups = append(groups, g)
		}
	}

	user.Group = groups

	db,err := core.GetSqlConn()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Save(user).Error
}

func (user User) HasGroup(group string) bool {
	for _,g := range user.Group {
		if g.Name == group {
			return true
		}
	}
	return false
}