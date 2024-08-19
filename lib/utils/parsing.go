package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ReqBodyParse(req *http.Request, data interface{}) {

    body, err := io.ReadAll(req.Body);
    if err != nil {
        return
    }
    defer req.Body.Close();

    err = json.Unmarshal(body, data);

    if err != nil {
        return
    }
}

func ParamsParse(req *http.Request, data interface{}) {
    params := mux.Vars(req);
    
    num, err1 := strconv.Atoi(params["id"]);
    
    if err1 != nil {
        return
    }
    
    formatted := map[string]interface{} {
        "id": num,
    }
    
    fmt.Println(formatted);
    
    formattedBytes, err := json.Marshal(formatted)
    if err != nil {
        return
    }
    
    fmt.Println(formattedBytes);

    err = json.Unmarshal(formattedBytes, data);

    if err != nil {
        return
    }
}