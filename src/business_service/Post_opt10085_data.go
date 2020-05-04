package business_service

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Post_opt10085_data(c *gin.Context) {

	db, err := sqlx.Connect("postgres", Get_connection_string())
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	var sb strings.Builder
	sb.WriteString(`select 
			purchase_dt,
			acct_num,
			stock_cd,
			curr_amt,
			purchase_amt,
			tot_purchase_amt,
			possession_qty,
			today_sell_profit_loss,
			today_commission,
			today_tax,
			credit_gubun,
			loan_dt,
			payment_balance,
			sellable_qty,
			credit_amt,
			credit_interest,
			expiry_dt,
			valuation_profit_loss,
			earnings_rt,
			evaluated_amt,
			commission,
			selling_commission,
			buying_commission,
			selling_tax,
			will_profit_amt,
			not_commission_profit_loss,
			profit_loss_rt,
			order_status,
			crt_dtm,
			purchase_dt ||'-'||acct_num||'-'||stock_cd||'-'|| crt_dtm id
		from tb_opt10085 
		order by crt_dtm asc
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
