package handler

import (
	"encoding/json"
	"ex5/api/db"
	"ex5/api/pkg/exam"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"xorm.io/xorm"
)

func conn() (engine *xorm.Engine) {
	engine, err := db.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	return engine
}

// ex2
func Ex2(w http.ResponseWriter, request *http.Request) {
	engine := conn()
	var users []exam.UserPartner
	test := exam.UserPartnerRequest{
		UserId: "user1",
		Phone: "1234",
		Limit: 3,
	}
	count, err := engine.Where("user_id = ?",test.UserId).Or("phone = ?",test.Phone).Limit(int(test.Limit),0).FindAndCount(&users)
	if err != nil {
		fmt.Fprintln(w,"Error :",err)
	}
	for i:= 0; i < int(count); i++ {
		fmt.Fprintf(w,"User: %v, phone: %v\n",users[i].UserId,users[i].Phone)
	}
}

// lay tat ca du lieu userparnter
func GetAllUser(w http.ResponseWriter, request *http.Request) {
	engine := conn()
	users, count := db.ListUser(engine, []exam.UserPartner{})
	for i := 0; i < int(count); i++ {
		fmt.Fprintf(w,"id: %v, User_id: %v, Phone: %v\n",users[i].Id, users[i].UserId, users[i].Phone)
	}

}

// lay du lieu theo UserId
func GetUserById(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	engine := conn()
	us, err := db.ReadUser(engine,id)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	if id == us.UserId {
		fmt.Fprintf(w,"User: %v, phone: %v",us.UserId, us.Phone)
		return
	}
	fmt.Fprintln(w,"User not found")
}

// tao user
func CreateUser(w http.ResponseWriter, request *http.Request) {
	engine := conn()
	reqBody,_ := ioutil.ReadAll(request.Body)
	user1 := exam.UserPartner{
		Id:          xid.New().String(),
		Created:     time.Now().UnixMilli(),
		UpdateAt:    time.Now().UnixMilli(),
	}
	json.Unmarshal(reqBody, &user1)
	err := db.InsertTable(engine,&user1)
	if err != nil {
		fmt.Fprint(w,"Error: ",err)
		return
	}
	fmt.Fprintf(w,"Them xg")
}

// sua user
func UpdateUser(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	engine := conn()
	us_id := params["id"]
	if us_id == "" {
		fmt.Fprint(w,"User not found")
		return
	}
	reqBody, _ :=ioutil.ReadAll(request.Body)
	user := exam.UserPartner{
		UpdateAt:    time.Now().UnixMilli(),
	}
	json.Unmarshal(reqBody,&user)
	err := db.UpdateUser(engine,us_id,user)
	if err != nil {
		fmt.Fprint(w,"Error: ",err)
		return
	}
	fmt.Fprintln(w,"Sua xg")
}

// xoa user
func DeleteUser(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	user_id := params["id"]
	engine := conn()
	err := db.DeleteUser(engine,user_id)
	if err != nil {
		fmt.Fprint(w,"Error:",err)
		return
	}
	fmt.Fprintf(w,"Xoa thanh cong")
}
