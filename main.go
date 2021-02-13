package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

type ProcessStatus struct {
	Proc string `json:"proc"`
	Ret  string `json:"ret"`
	cnt  int
}

func incCnt(psList *map[string]ProcessStatus, k string) {
	val, exists := (*psList)[k]

	if !exists {
		(*psList)[k] = ProcessStatus{Proc: k, Ret: "", cnt: 1}
		return
	}

	val.cnt++
	(*psList)[k] = val
}

func main() {

	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)

	port, _ := strconv.Atoi(result["port"].(string))
	_ = port

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		ps, _ := ps.Processes()

		var retList = make(map[string]ProcessStatus)

		for k, v := range result["processes"].(map[string]interface{}) {
			targetProc, _ := v.(map[string]interface{})["proc"].(string)
			for _, v2 := range ps {

				if v2.Executable() == "" {
					continue
				}

				if strings.Index(v2.Executable(), targetProc) == 0 {
					incCnt(&retList, k)
				}
			}
		}

		for k, v := range result["processes"].(map[string]interface{}) {
			// targetProc, _ := v.(map[string]interface{})["proc"].(string)
			retType, _ := v.(map[string]interface{})["ret"].(string)
			ret, _ := retList[k]
			if retType == "cnt" {
				ret.Ret = strconv.Itoa(retList[k].cnt)
			} else {
				if ret.cnt > 0 {
					ret.Ret = "on"
				} else {
					ret.Ret = "off"
				}
			}

			retList[k] = ret
		}

		jsonString, _ := json.Marshal(retList)

		w.Write([]byte(jsonString))
	})

	fmt.Printf("server listening on port %v\n", port)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
