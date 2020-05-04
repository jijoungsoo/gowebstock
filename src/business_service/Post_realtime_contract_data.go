package business_service

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Post_realtime_contract_data(c *gin.Context) {
	stock_cd := c.PostForm("stock_cd")
	stock_dt := c.PostForm("stock_dt")
	fmt.Printf("stock_cd data: %s\n", stock_cd)

	db, err := sqlx.Connect("postgres", Get_connection_string())
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	var sb strings.Builder
	sb.WriteString(`
	with vw00 as (
		select 
			stock_dt
			,stock_cd
			,curr_time
			,count(*) OVER(
					partition by stock_cd
					, stock_dt
					, ((cast(substring(curr_time,1,2) as integer)-9 )*60 + cast(substring(curr_time,3,2)  as integer) )/b.divide
			) cnt
			,FIRST_VALUE(curr_amt) OVER(
					partition by stock_cd
					, stock_dt
					, ((cast(substring(curr_time,1,2) as integer)-9 )*60 + cast(substring(curr_time,3,2)  as integer) )/b.divide
					order by curr_time asc
			) start_amt
			,last_valUE(curr_amt) OVER(
					partition by stock_cd
					, stock_dt
					, ((cast(substring(curr_time,1,2) as integer)-9 )*60 + cast(substring(curr_time,3,2)  as integer) )/b.divide
					order by curr_time asc		  
			) curr_amt
			,RANK() OVER(
					partition by stock_cd
					, stock_dt
					, ((cast(substring(curr_time,1,2) as integer)-9 )*60 + cast(substring(curr_time,3,2)  as integer) )/b.divide
					order by curr_time		  
			) rnk
			,((cast(substring(curr_time,1,2) as integer)-9 )*60 + cast(substring(curr_time,3,2)  as integer) )/b.divide  mm
			,max(abs(curr_amt)) over (partition by stock_cd
									, stock_dt
									, ((cast(substring(curr_time,1,2) as integer)-9 )*60 + cast(substring(curr_time,3,2)  as integer) )/b.divide 
									) as high_amt 
			,min(abs(curr_amt)) over (partition by stock_cd
									, stock_dt
									, ((cast(substring(curr_time,1,2) as integer)-9 )*60 + cast(substring(curr_time,3,2)  as integer) )/b.divide 
									) as low_amt 
			,(case when trade_qty>0 then trade_qty else 0 end)  offered_trade_qty
			,(case when trade_qty<0 then trade_qty else 0 end)  bid_trade_qty
			,b.divide
		from tb_realtime_contract  a
		cross join (select 5 as  divide ) b
		where 1=1`)
	//if stock_cd != "" {
	sb.WriteString(` 
		and a.stock_cd='005930'`)
	//}

	if stock_dt != "" {
		sb.WriteString(` 
		and a.stock_dt='20200403'`)
	}

	sb.WriteString(`
		)
		select 
			stock_dt
			,stock_cd
			,curr_amt
			,mm
			,curr_time
			,start_amt
			,high_amt
			,low_amt
			,offered_trade_qty
			,bid_trade_qty 
			,divide
			,cnt
			,stock_dt ||'-'|| stock_cd ||'-'|| curr_time  id 
		from vw00 
		where rnk=1
		order by mm asc
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
	//fmt.Printf("Marshalled data: %s\n", mm)
	//엄청 많은 건을 가지고 와서 1분 27초만에 조회를 했는데
	//출력도 되었는데
	//c.JSON에서 해결이 안된다.
	//혹은 여기서 해결이 되었는데 chrome에서 해결을 못한다.
	//메모리가 1,8G가 까지 치솟았는데 네트워크에서는 응답을 못받았다.
	//sql 로그 남기는 건
	//https://github.com/luna-duclos/instrumentedsql 이것을 해보자 dddata는 유료다.
	//api 툴에서 호출해보았다.
	// response로는 값이 안보이는데 hex로 보면 값이 전달됨을 확인했다.
	// 네트워크는 전달되었다고 보는게 맞는것 같다.
	// 그럼 그리드에서 확인이 안된다는건데.
	// console.log로 데이터가 왔는지 찍어보자.
	// console.log로도 데이터가 찍힘을 확인했다.
	// 그리드에서 못받춰주는건지. grid.on 걸었던 이벤트에서 찍은 로그가 문제인지 확인해보자.
	// 결국엔 그리드가 문제였다. 대용량 데이터를 처리못하는데 네이버 그리드에서는
	// 대용량 데이터도 처리하는 것처럼 설명해두었다.

	c.JSON(200, mm)
}
