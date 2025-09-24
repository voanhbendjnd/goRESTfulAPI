package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
  // log.Println("Loading server");
  // fmt.Println("Loading server 2");

  http.HandleFunc("/demo", demoHandler)

  err := http.ListenAndServe(":8080", nil);
  if (err != nil){
    log.Fatal("Server starter error"); // = return 0
  }
}
func demoHandler(response http.ResponseWriter,request *http.Request){
  log.Printf("%+v", request);
  if(request.Method != http.MethodGet){
    http.Error(response, "Chỉ áp dụng với phương thức GET", http.StatusMethodNotAllowed);
    return;
  }
  res := map[string]string{
    "message":"Đây là phương thức GET",
  }
  response.Header().Set("Content-Type", "applicaton/json")
  // data, err := json.Marshal(res);
  // if(err != nil){
  //   http.Error(response, "Lỗi mã hóa JSON", http.StatusInternalServerError)
  //   return;
  // }
  // response.Write(data)
  json.NewEncoder(response).Encode(res);
}