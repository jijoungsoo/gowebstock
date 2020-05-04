package CM_0100

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	bs "business_service"
)

func Get_pgm_mngr_data(c *gin.Context) {
	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`SELECT 
		pgm_id,
		pgm_nm,
		category,
		pgm_link,
		rmk	,
		to_char(crt_dtm, 'YYYY-MM-DD') crt_dtm,		
		to_char(updt_dtm, 'YYYY-MM-DD') updt_dtm		
	FROM tb_cm_pgm a
	order by pgm_id
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
