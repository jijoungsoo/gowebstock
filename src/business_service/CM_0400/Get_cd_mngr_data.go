package CM_0400

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	bs "business_service"
)

type (
	// transformedTodo represents a formatted todo
	paramData struct {
		GrpCd string `json:"grp_cd" binding:"required"`
	}
)

func Get_cd_mngr_data(c *gin.Context) {

	var json_tmp paramData //array로 감싸서 넘겨야한다. []
	if err := c.ShouldBindJSON(&json_tmp); err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var grp_cd string

	fmt.Printf("grp_cd: %s\n", json_tmp)
	grp_cd = json_tmp.GrpCd

	fmt.Printf("grp_cd: %s\n", grp_cd)

	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`select 
	 grp_cd
	,cd
	,cd_nm
	,use_yn
	,ord
	,rmk
	,crt_dtm
	,updt_dtm
	from tb_cm_cd	
	where 1=1
		`)
	if grp_cd != "" {
		sb.WriteString(` 
		and 	grp_cd=:grp_cd`)
	}
	sb.WriteString(`
	order by ord,grp_cd,cd
		`)
	p_m := map[string]interface{}{}
	if grp_cd != "" {
		p_m["grp_cd"] = grp_cd
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
