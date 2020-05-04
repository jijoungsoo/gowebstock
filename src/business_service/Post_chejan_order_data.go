package business_service

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Post_chejan_order_data(c *gin.Context) {
	stock_cd := c.PostForm("stock_cd")
	fmt.Printf("stock_cd data: %s\n", stock_cd)
	db, err := sqlx.Connect("postgres", Get_connection_string())
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	var sb strings.Builder
	sb.WriteString(`
			select 
			contract_time
				,acct_num
				,order_num
				,stock_cd
				,order_business_classification
				,order_status
				,order_qty
				,order_amt
				,not_contract_qty
				,contract_tot_amt
				,ongn_order_num
				,order_gubun
				,trade_gubun

				,order_type
				,contract_num
				,contract_amt
				,contract_qty
				,curr_amt
				,offered_amt
				,bid_amt
				,contract_amt_unit
				,contract_amt_qty
				,today_commission
				,today_tax
				,screen_num
				,terminal_num
				,credit_gubun
				,loan_dt
				,crt_dtm
				,acct_num ||'-'|| order_num ||'-'|| contract_num ||'-'|| contract_time ||'-'|| crt_dtm  id 
			from tb_chejan_order
			order by crt_dtm desc
		`)
	//Go through rows
	rows, err := db.Queryx(sb.String())
	//fmt.Printf("Marshalled data: %s\n", mm)
	mm := []map[string]interface{}{}
	for rows.Next() {
		m := map[string]interface{}{}
		err := rows.MapScan(m)
		mm = append(mm, m)
		if err != nil {
			log.Fatal(err)
		}
	}
	c.JSON(200, mm)
}
