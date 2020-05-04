package business_service

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Post_opt10001_data(c *gin.Context) {
	db, err := sqlx.Connect("postgres", Get_connection_string())
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	var sb strings.Builder

	sb.WriteString(`SELECT 
		b.stock_cd,
		b.stock_nm,
		b.stock_cnt,
		b.last_price,
		a.face_amt, 
		a.capital_amt, 
		a.credit_rt, 
		CAST(TRUNC(CAST(a.credit_rt as numeric),2) as double precision) credit_rt,
		a.settlement_mm, 
		a.year_high_amt, 
		a.year_low_amt, 
		a.total_mrkt_amt, 
		a.total_mrkt_amt_rt, 
		CAST(TRUNC(CAST(a.total_mrkt_amt_rt as numeric),2) as double precision) total_mrkt_amt_rt,
		CAST(TRUNC(CAST(a.foreigner_exhaustion_rt as numeric),2) as double precision) foreigner_exhaustion_rt,
		a.substitute_amt, 
		CAST(TRUNC(CAST(a.per as numeric),2) as double precision) per,
		CAST(TRUNC(CAST(a.eps as numeric),2) as double precision) eps,
		CAST(TRUNC(CAST(a.roe as numeric),2) as double precision) roe,
		CAST(TRUNC(CAST(a.pbr as numeric),2) as double precision) pbr,
		CAST(TRUNC(CAST(a.ev as numeric),2) as double precision) ev,
		CAST(TRUNC(CAST(a.bps as numeric),2) as double precision) bps,
		a.sales, 
		a.business_profits, 
		a.d250_high_amt, 
		a.d250_low_amt, 
		a.start_amt, 
		a.high_amt, 
		a.low_amt, 
		a.upper_amt_lmt, 
		a.lower_amt_lmt, 
		a.yesterday_amt, 
		a.expectation_contract_amt, 
		a.expectation_contract_qty, 
		to_char(to_date(a.d250_high_dt,'YYYYMMDD'), 'YYYY-MM-DD' ) d250_high_dt,
		CAST(TRUNC(CAST(a.d250_high_rt as numeric),2) as double precision) d250_high_rt,
		CAST(TRUNC(CAST(a.d250_low_rt as numeric),2) as double precision) d250_low_rt,
		a.curr_amt, 
		a.contrast_symbol, 
		a.contrast_yesterday, 
		CAST(TRUNC(CAST(a.fluctuation_rt as numeric),2) as double precision) fluctuation_rt,
		a.trade_qty, 
		CAST(TRUNC(CAST(a.trade_contrast as numeric),2) as double precision) trade_contrast,
		to_char(a.crt_dtm, 'YYYY-MM-DD') crt_dtm,
		to_char(to_date(a.d250_low_dt,'YYYYMMDD'), 'YYYY-MM-DD' ) d250_low_dt,
		a.face_amt_unit,
		b.stock_cd id
	FROM tb_opt10001 a
	right outer join tb_stock b 
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
	c.JSON(200, mm)
}
