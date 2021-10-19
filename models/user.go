package models

import (
	"fmt"

	"github.com/spf13/cast"

	"github.com/beego/beego/v2/client/orm"
)

//create database orm_test;
//create table user ( id integer primary key, name varchar(255), profile_id integer );
type User struct {
	Id        int    `orm:"column(id);"`
	FirstName string `orm:"column(firstName);size(32)"`
	LastName  string `orm:"column(lastName);size(32)"`
	Phone     string `orm:"column(phone);size(10)"`
	Email     string `orm:"column(email);size(48)"`
	Company   string `orm:"column(company);size(128)"`
}

func (u *User) TableName() string {
	return "fission_user"
}

func AddUser(user *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err != nil {
		return 0, err
	}
	return id, nil

}

func GetAllUsers() ([]User, error) {
	o := orm.NewOrm()
	var v []orm.Params
	users := make([]User, 0)
	_, err := o.Raw("SELECT * FROM fission_user").Values(&v)
	if err != nil {
		fmt.Println("Error in getting all the users", err)
		return users, err
	}
	for _, row := range v {
		user := User{}
		user.Id = cast.ToInt(row["id"])
		user.FirstName = cast.ToString(row["firstName"])
		user.LastName = cast.ToString(row["lastName"])
		user.Phone = cast.ToString(row["phone"])
		user.Email = cast.ToString(row["email"])
		user.Company = cast.ToString(row["company"])
		users = append(users, user)
	}
	return users, nil

}

func UpdateUser(user *User) error {
	o := orm.NewOrm()
	id, err := o.Update(user)
	if err != nil {
		fmt.Println("Error in updating the user", err)
		return err
	}
	fmt.Println("custom", id)
	return nil
}
