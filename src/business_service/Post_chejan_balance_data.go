package business_service

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Post_chejan_balance_data(c *gin.Context) {
	stock_cd := c.PostForm("stock_cd")
	fmt.Printf("stock_cd data: %s\n", stock_cd)
	db, err := sqlx.Connect("postgres", Get_connection_string())
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	var sb strings.Builder
	sb.WriteString(`
		select 
			 acct_num
			,stock_cd
			,curr_amt
			,possession_qty
			,purchase_amt
			,tot_purchase_amt
			,order_possible_qty
			,today_net_buy_qty
			,order_type
			,today_sell_profit_loss
			,deposit

			,offered_amt
			,bid_amt
			,yesterday_amt
			,profit_loss_rt
			,credit_amt
			,credit_interest
			,expiry_dt
			,today_profit_loss_amt
			,today_profit_loss_rt
			,credit_today_profit_loss_amt
			,credit_today_profit_loss_rt
			,loan_qty
			,loan_dt
			,credit_gubun
			,crt_dtm
			,acct_num ||'-'|| stock_cd ||'-'|| crt_dtm  id 
		from tb_chejan_balance
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
