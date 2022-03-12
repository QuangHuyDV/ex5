package main

import (
	"context"
	"ex5/grpc/exam1"
	"io"
	"log"
	// "time"

	// "github.com/rs/xid"
	"google.golang.org/grpc"
)

//bài 5: Tài liệu grpc Tạo 1 service gen code. Tạo 1 grpc server với message UserPartner. Nhằm getlist, create, update Tạo 1 grpc client để thực hiện

func main() {
	cc, err := grpc.Dial("localhost:3003", grpc.WithInsecure())

	if err != nil {
		log.Fatalf(" err while dial %v", err)
	}
	defer cc.Close()

	client := exam1.NewUserServerClient(cc)

	log.Println("Client running...")

	// InsertUser(client, exam1.UserPartner{
	// 	Id: xid.New().String(),
	// 	UserId: "user6",
	// 	PartnerId: "a6",
	// 	AliasUserId: "a6",
	// 	Apps: map[string]int64{
	// 		"qqq" : 222,
	// 		"sss" : 333,
	// 	},
	// 	Phone: "12342334",
	// 	Created: time.Now().UnixMilli(),
	// 	UpdateAt: time.Now().UnixMilli(),
	// })

	// UpdateUser(client, "user5", exam1.UserPartner{
	// 	PartnerId: "sa",
	// 	AliasUserId: "sa",
	// 	Phone: "12345",
	// 	UpdateAt: time.Now().UnixMilli(),
	// })

	// ReadUser(client, "user2")

	ListUser(client, 10)
}

// Them userpartner (request)
func InsertUser(cli exam1.UserServerClient, user1 exam1.UserPartner) {
	req := &exam1.InsertRequest{
		User: &user1,
	}
	resp, err := cli.Insert(context.Background(), req)

	if err != nil {
		log.Printf("call insert err %v\n", err)
		return
	}

	log.Printf("insert response %+v\n", resp)
}

// Doc userpartner (request)
func ReadUser(cli exam1.UserServerClient, userId string) {
	req := &exam1.ReadRequest{
		UserId: userId,
	}
	resp, err := cli.Read(context.Background(), req)
	if err != nil {
		log.Printf("call read err %v\n", err)
		return
	}

	log.Printf("insert response %+v\n", resp)
}

// Sua userpartner (request)
func UpdateUser(cli exam1.UserServerClient, id string, user1 exam1.UserPartner) {
	req := &exam1.UpdateRequest{
		UserID: id,
		NewUser: &user1,
	}
	resp, err := cli.Update(context.Background(), req)

	if err != nil {
		log.Printf("call update err %v\n", err)
		return
	}

	log.Printf("insert response %+v\n", resp)
}

// Danh sach userparnter (request)
func ListUser(cli exam1.UserServerClient, num int32) {
	req := &exam1.ListRequest{
		Count: num,
	}
	resp, err := cli.List(context.Background(), req)

	if err != nil {
		log.Printf("call list err %v\n", err)
		return
	}
	for {
		user, err := resp.Recv()
		if err == io.EOF {
			log.Println("server finish streaming")
			return
		}

		if err != nil {
			log.Fatalf("List error %v", err)
		}

		log.Printf("User: User_id: %v, Phone: %v\n", user.Users.UserId, user.Users.Phone)
	}
}
