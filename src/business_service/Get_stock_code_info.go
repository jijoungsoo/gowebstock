package business_service

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Get_stock_code_info(c *gin.Context) {
	//grp_cd := c.PostForm("grp_cd")
	type (
		// transformedTodo represents a formatted todo
		p_json struct {
			StockCd string `json:"stock_cd" binding:"required"`
		}
	)

	var json_tmp p_json
	if err := c.ShouldBindJSON(&json_tmp); err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stock_cd := json_tmp.StockCd

	db, err := sqlx.Connect("postgres", Get_connection_string())
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	var sb strings.Builder
	sb.WriteString(`SELECT 
		stock_cd,
		face_amt, 
		capital_amt, 
		stock_cnt, 
		credit_rt, 
		CAST(TRUNC(CAST(credit_rt as numeric),2) as double precision) credit_rt,
		settlement_mm, 
		year_high_amt, 
		year_low_amt, 
		total_mrkt_amt, 
		total_mrkt_amt_rt, 
		CAST(TRUNC(CAST(total_mrkt_amt_rt as numeric),2) as double precision) total_mrkt_amt_rt,
		CAST(TRUNC(CAST(foreigner_exhaustion_rt as numeric),2) as double precision) foreigner_exhaustion_rt,
		substitute_amt, 
		CAST(TRUNC(CAST(per as numeric),2) as double precision) per,
		CAST(TRUNC(CAST(eps as numeric),2) as double precision) eps,
		CAST(TRUNC(CAST(roe as numeric),2) as double precision) roe,
		CAST(TRUNC(CAST(pbr as numeric),2) as double precision) pbr,
		CAST(TRUNC(CAST(ev as numeric),2) as double precision) ev,
		CAST(TRUNC(CAST(bps as numeric),2) as double precision) bps,
		sales, 
		business_profits, 
		d250_high_amt, 
		d250_low_amt, 
		start_amt, 
		high_amt, 
		low_amt, 
		upper_amt_lmt, 
		lower_amt_lmt, 
		yesterday_curr_amt, 
		expectation_contract_amt, 
		expectation_contract_qty, 
		to_char(to_date(d250_high_dt,'YYYYMMDD'), 'YYYY-MM-DD' ) d250_high_dt,
		CAST(TRUNC(CAST(d250_high_rt as numeric),2) as double precision) d250_high_rt,
		CAST(TRUNC(CAST(d250_low_rt as numeric),2) as double precision) d250_low_rt,
		curr_amt, 
		contrast_symbol, 
		contrast_yesterday, 
		CAST(TRUNC(CAST(fluctuation_rt as numeric),2) as double precision) fluctuation_rt,
		trade_qty, 
		CAST(TRUNC(CAST(trade_contrast as numeric),2) as double precision) trade_contrast,
		to_char(crt_dtm, 'YYYY-MM-DD') crt_dtm,
		to_char(to_date(d250_low_dt,'YYYYMMDD'), 'YYYY-MM-DD' ) d250_low_dt,
		face_amt_unit,
		stock_cd id
		FROM public.tb_opt10001
		where stock_cd=:stock_cd
		order by stock_cd`)
	p_m := map[string]interface{}{}
	if stock_cd != "" {
		p_m["stock_cd"] = stock_cd
	}
	rows, err := db.NamedQuery(sb.String(), p_m)

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
