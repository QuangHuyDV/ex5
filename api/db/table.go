package db

import (
	"ex5/api/pkg/exam"
	"log"

	"xorm.io/xorm"
)

//create
func CreateTable(engine *xorm.Engine, tb interface{}) error {
	_, err := engine.IsTableExist(tb)
	if err != nil {
		log.Println(err)
		return err
	}
	err = engine.Sync2(tb)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//insert
func InsertTable(engine *xorm.Engine, data interface{}) error {
	_, err := engine.Insert(data)
	if err != nil {
		return err
	}
	return nil
}

//read 
func ReadUser(engine *xorm.Engine, id string) (*exam.UserPartner, error) {
	tb := exam.UserPartner{}
	_, err := engine.Where("user_id = ?",id).Get(&tb)
	if err != nil {
		log.Println(err)
	}
	return &tb,nil
}

//list 
func ListUser(engine *xorm.Engine, tb []exam.UserPartner) ([]exam.UserPartner, int64) {
	count,_ := engine.Count(exam.UserPartner{})
	err := engine.Limit(int(count),0).Find(&tb)
	if err != nil {
		log.Println(err)	
	}
	return tb ,count
}

//delete
func DeleteUser(engine *xorm.Engine, id string) error {
	_, err := engine.Where("user_id = ?", id).Delete(exam.UserPartner{})
	if err != nil {
		return err
	}
	return nil
}

func Delete(engine *xorm.Engine) error {
	id := ""
	_, err := engine.Where("user_id = ?", id).Delete(exam.UserPartner{})
	if err != nil {
		return err
	}
	return nil
}

//update
func UpdateUser(engine *xorm.Engine, id string, data interface{}) error {
	_, err := engine.Where("user_id = ?",id).Update(data)
	if err != nil {
		return err
	}
	return nil
}