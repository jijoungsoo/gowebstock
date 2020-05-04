package CM_0300

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	bs "business_service"
)

func Get_menu_mngr_data(c *gin.Context) {

	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`select 
	menu_cd
	,prnt_menu_cd
	,menu_kind
	,menu_nm
	,rmk
	,fst_ord
	,sed_ord
	,pgm_id
	,crt_dtm
	,updt_dtm 
	,menu_cd id
	from tb_cm_menu
	order by fst_ord,sed_ord
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
