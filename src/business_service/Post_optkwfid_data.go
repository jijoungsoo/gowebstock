package business_service

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Post_optkwfid_data(c *gin.Context) {
	db, err := sqlx.Connect("postgres", Get_connection_string())
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	var sb strings.Builder

	sb.WriteString(`select 
	a.stock_cd,
	b.stock_nm,
	a.yesterday_amt,
	a.curr_amt,
	a.contrast_yesterday,
	a.contrast_yesterday_symbol,
	CAST(TRUNC(CAST(a.fluctuation_rt as numeric),2) as double precision) fluctuation_rt,
	a.trade_qty,
	a.trade_amt,
	a.contract_qty,
	CAST(TRUNC(CAST(a.contract_strength as numeric),2) as double precision) contract_strength,
	CAST(TRUNC(CAST(a.yesterday_contrast_trade_rt as numeric),2) as double precision) yesterday_contrast_trade_rt,
	a.offered_amt,
	a.bid_amt,
	a.offered_amt_one,
	a.offered_amt_two,
	a.offered_amt_three,
	a.offered_amt_four,
	a.offered_amt_five,
	a.bid_amt_one,
	a.bid_amt_two,
	a.bid_amt_three,
	a.bid_amt_four,
	a.bid_amt_five,
	a.upper_amt_lmt,
	a.lower_amt_lmt,
	a.start_amt,
	a.high_amt,
	a.low_amt,
	a.clsg_amt,
	a.contract_time,
	a.expectation_contract_amt,
	a.expectation_contract_qty,
	a.capital_amt,
	a.face_amt,
	a.total_mrkt_amt,
	a.stock_cnt,
	a.hoga_time,
	to_char(to_date(a.stock_dt,'YYYYMMDD'), 'YYYY-MM-DD' ) stock_dt,	
	a.fst_offered_balance,
	a.fst_bid_balance,
	a.fst_offered_qty,
	a.fst_bid_qty,
	a.tot_offered_balance,
	a.tot_bid_balance,
	a.tot_offered_qty,
	a.tot_bid_qty,
	a.parity_rt,
	a.gearing,
	a.break_even_point,
	a.elw_strike_amt,
	a.conversion_rt,
	to_char(to_date(a.elw_expiry_dt,'YYYYMMDD'), 'YYYY-MM-DD' ) elw_expiry_dt,
	a.open_interest,
	a.contrast_open_interest,
	a.theorist_amt,
	a.implied_volatility,
	a.delta,
	a.gamma,
	a.theta,
	a.vega,
	a.lo,
	to_char(a.crt_dtm, 'YYYY-MM-DD' ) crt_dtm,	
	a.stock_cd id
	from tb_optkwfid a 
	left outer join tb_stock b 
	on a.stock_cd = b.stock_cd
	order by a.stock_cd
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
	//c.JSON(200, mm)
	c.PureJSON(200, mm)
}
