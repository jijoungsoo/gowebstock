package business_service

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Post_theme_cd_data(c *gin.Context) {
	db, err := sqlx.Connect("postgres", Get_connection_string())
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	var sb strings.Builder

	sb.WriteString(`select 
	theme_cd,
	theme_nm,
	(select count(*) from tb_theme_stock b where b.theme_cd=a.theme_cd) stock_cnt,
	(select ARRAY_TO_STRING(ARRAY_AGG(substring(stock_cd,2) ORDER BY stock_cd),',')  from tb_theme_stock b where b.theme_cd=a.theme_cd) list_stock_cd
	from tb_theme a
	order by theme_cd asc	
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
