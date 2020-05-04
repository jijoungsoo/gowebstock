package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	// transformedTodo represents a formatted todo
	SendOrder struct {
		RqName    string `json:"rq_name" binding:"required"`
		AcctNum   string `json:"acct_num" binding:"required"`
		OrderType string `json:"order_type" binding:"required"`
		StockCd   string `json:"stock_cd" binding:"required"`
		Qty       string `json:"qty" binding:"required"`
		/*
			※ 시장가, 최유리지정가, 최우선지정가, 시장가IOC, 최유리IOC, 시장가FOK, 최유리FOK,
			장전시간외, 장후시간외 주문시 주문가격을 입력하지 않습니다.
			0으로 보내야한다. WCF에서 빈값을 인식못해서 매칭이 안되서 실행이 안된다.
			금액이 문자열이면 "" 디폴트로 C#에서 해주면 되는데 INT형이라 그렇게 못했다.
			UI에서 0을 입력하도록 한다.
		*/
		Amt string `json:"amt"  binding:"required" `

		HogaGubun   string `json:"hoga_gubun" binding:"required"`
		OrgnOrderNo string `json:"orgn_order_no"`
	}
)

//{"acct_num":"1","order_type":"2","stock_cd":"3","qty":"4","amt":"5","hoga_gubun":"6","orgn_order_no":"2"}
func external_send_order(c *gin.Context) {
	var json_tmp SendOrder
	if err := c.ShouldBindJSON(&json_tmp); err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("json_tmp.acct_num %s\n", json_tmp.AcctNum)
	fmt.Printf("json_tmp.order_type %s\n", json_tmp.OrderType)
	fmt.Printf("json_tmp.stock_cd %s\n", json_tmp.StockCd)
	fmt.Printf("json_tmp.qty %s\n", json_tmp.Qty)
	fmt.Printf("json_tmp.amt %s\n", json_tmp.Amt)
	fmt.Printf("json_tmp.hoga_gubun %s\n", json_tmp.HogaGubun)
	fmt.Printf("json_tmp.orgn_order_no %s\n", json_tmp.OrgnOrderNo)
	/*
		url := "http://127.0.0.1:5000/kiwoomOpenApi/json/SendOrder"
		resp, _ := http.Post(url)
		robots, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s\n", robots)
	*/
	/*json 데이터 전달  http://golang.site/go/article/103-HTTP-POST-%ED%98%B8%EC%B6%9C  */

	url := "http://127.0.0.1:5000/kiwoomOpenApi/json/SendOrder"
	pbytes, _ := json.Marshal(json_tmp)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(url, "application/json", buff)
	fmt.Printf("%s\n", "aaaa1")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	resp.Body.Close()
	c.JSON(200, json_tmp)
}
