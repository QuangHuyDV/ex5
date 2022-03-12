package main

import (
	"ex5/api/pkg/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "Err: ", http.StatusNotFound)
        return
    }
    fmt.Fprintf(w, "Pong!")
}

func main() {

	r := mux.NewRouter()
	fmt.Println("Starting server at port 3003...")
	
	// bài 2: Viết một message UserPartnerRequest nhằm tạo 1 query xorm. Bao gồm lấy userpartner theo user_id, phone, với limit là số lượng row lớn nhất được quét ra. Với id được genere ngẫu nhiên với xid
	r.HandleFunc("/ex2",handler.Ex2)

	// bài 3: Sử dụng kiến thức đã đọc tao 1 server net/http hoặc mux sử dụng port 3001 mở 1 url. Viết 1 route / trả về pong`
	r.HandleFunc("/", helloHandler)

	// bài 4: kết hợp với kiến thức trên viết 1 server sử dụng grpc generate 2 message ở bài 1, kết hợp server bài 2 và kiến thức từ bài trước xorm. Viết 1 route: POST /user-partner tạo mới 1 partner, GET /user-partner lấy danh sách partner, DELETE /user-partner/{id} theo id, GET /user-partner/{id} lấy theo 1 id cụ thể.
	r.HandleFunc("/user-partner", handler.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/user-partner/{id}", handler.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/user-partner/create", handler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/user-partner/update/{id}", handler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user-partner/delete/{id}", handler.DeleteUser).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":3003", r))

}