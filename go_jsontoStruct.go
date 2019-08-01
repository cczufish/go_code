/*

{
    "resp": {
        "respCode": "000000",
        "respMsg": "成功",
        "app": {
            "appId": "d12abd3da59d47e6bf13893ec43730b8"
        }
    }
}

https://blog.csdn.net/zxy_666/article/details/80173288

*/

package main
import (
	"fmt"
	"encoding/json"

)
type appInfo struct {
	Appid string `json:"appId"`
}
type response struct {
	RespCode string  `json:"respCode"`
	RespMsg  string  `json:"respMsg"`
	AppInfo  appInfo `json:"app"`
}
type JsonResult struct {
	Resp response `json:"resp"`
}
func main() {
	jsonstr := `{"resp": {"respCode": "000000","respMsg": "成功","app": {"appId": "d12abd3da59d47e6bf13893ec43730b8"}}}`
	fmt.Println("jsonstr:", jsonstr)

	json := []byte(jsonstr)

	jsonresponse := JsonResult{}

	err:=json.Unmarshal(json,&jsonresponse)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("after parse", jsonresponse)
}


