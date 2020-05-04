package CM_0400

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	bs "business_service"
)

func Get_cd_grp_mngr_data(c *gin.Context) {

	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`select 
	 grp_cd
	,grp_cd_nm
	,use_yn
	,ord
	,rmk
	,crt_dtm
	,updt_dtm
	from tb_cm_cd_grp
	order by ord,grp_cd
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
