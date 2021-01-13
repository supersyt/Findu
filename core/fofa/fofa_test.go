package fofa

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestFofa_QueryAsJSON(t *testing.T) {

	email := os.Getenv("FOFA_EMAIL")
	key := os.Getenv("FOFA_KEY")
	clt := NewFofaClient([]byte(email), []byte(key))
	if clt == nil {
		fmt.Printf("create fofa client\n")
		return
	}
	//result, err := clt.QueryAsJSON(1, []byte(`body="小米"`))
	//if err != nil {
	//	fmt.Printf("%v\n", err.Error())
	//}
	//fmt.Printf("%s\n", result)
	arr, err := clt.QueryAsArray(1, []byte(`fofa.so`), []byte("title,port,protocol,domain,host"))
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	fmt.Printf("count: %d\n", len(arr))
	encodeArr, _ := json.Marshal(arr)
	fmt.Printf("\n%s\n", encodeArr)
}
