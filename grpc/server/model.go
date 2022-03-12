package main

import (
	"ex5/api/db"
	"ex5/api/pkg/exam"
	"ex5/grpc/exam1"
	"fmt"
	"log"

	"xorm.io/xorm"
)

type UserInfo struct {
	Id                   string           
	UserId               string           
	PartnerId            string           
	AliasUserId          string           
	Apps                 map[string]int64 
	Phone                string           
	Created              int64            
	UpdateAt             int64            
}
// chuyen doi du lieu struct
func ConvertUser(ci *exam1.UserPartner) *UserInfo {
	return &UserInfo{
		Id: ci.Id,
		UserId: ci.UserId,
		PartnerId: ci.PartnerId,
		AliasUserId: ci.AliasUserId,
		Apps: ci.Apps,
		Phone: ci.Phone,
		Created: ci.Created,
		UpdateAt: ci.UpdateAt,
	}
}

func ConvertUser2(ci *UserInfo) *exam1.UserPartner {
	return &exam1.UserPartner{
		Id: ci.Id,
		UserId: ci.UserId,
		PartnerId: ci.PartnerId,
		AliasUserId: ci.AliasUserId,
		Apps: ci.Apps,
		Phone: ci.Phone,
		Created: ci.Created,
		UpdateAt: ci.UpdateAt,
	}
}

func ConvertUser1(ci *exam.UserPartner) *exam1.UserPartner {
	return &exam1.UserPartner{
		Id: ci.Id,
		UserId: ci.UserId,
		PartnerId: ci.PartnerId,
		AliasUserId: ci.AliasUserId,
		Apps: ci.Apps,
		Phone: ci.Phone,
		Created: ci.Created,
		UpdateAt: ci.UpdateAt,
	}
}

func conn() *xorm.Engine {
	engine, err := db.Connect()
	if err != nil {
		fmt.Println(err)
		return nil
	} 
	return engine
}

// them du lieu vao database
func (c *UserInfo) InsertUser() error {
	engine := conn()
	_,err := engine.Table("user_partner").Insert(&c)
	if err != nil {
		log.Printf("Insert user %+v err %v\n",c,err)
		return err
	}
	log.Printf("Insert %+v successfully\n", c)
	return nil
}

// doc du lieu trong database
func ReadUser(id string) (*exam.UserPartner,error) {
	engine := conn()
	tb := exam.UserPartner{}
	_, err := engine.Table("user_partner").Where("user_id = ?",id).Get(&tb)
	if err != nil {
		log.Println(err)
	}
	return &tb,nil
}

// sua du lieu trong database
func (c *UserInfo) UpdateUser(id string) error {
	engine := conn()
	_, err := engine.Table("user_partner").Where("user_id = ?",id).Update(c)
	if err!= nil {
		log.Printf("Update %+v err %v\n", c.UserId, err)
		return err
	}
	log.Printf("update user %+v\n", c.UserId)
	return nil
}

// lay danh sach trong database
func ListUser(Count int32) ([]exam1.UserPartner, error) {
	engine := conn()
	tb := []exam1.UserPartner{}
	err := engine.Table("user_partner").Cols("id","user_id","partner_id","alias_user_id","phone").Limit(int(Count),0).Find(&tb)
	if err != nil {
		log.Println(err)	
	}
	return tb,nil
}