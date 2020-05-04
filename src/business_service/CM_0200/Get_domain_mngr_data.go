package CM_0200

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	bs "business_service"
)

func Get_domain_mngr_data(c *gin.Context) {

	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder
	sb.WriteString(`SELECT 
		domain_cd,
		domain_nm,
		data_type,
		to_char(crt_dtm, 'YYYY-MM-DD') crt_dtm,
		to_char(updt_dtm, 'YYYY-MM-DD') updt_dtm,
		rmk	
	FROM tb_cm_domain a
	order by domain_nm,domain_cd
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
