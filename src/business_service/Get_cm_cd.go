package business_service

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Get_cm_cd(c *gin.Context) {
	//grp_cd := c.PostForm("grp_cd")
	type (
		// transformedTodo represents a formatted todo
		p_json struct {
			GrpCd string `json:"grp_cd" binding:"required"`
		}
	)

	var json_tmp p_json
	if err := c.ShouldBindJSON(&json_tmp); err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	grp_cd := json_tmp.GrpCd

	db, err := sqlx.Connect("postgres", Get_connection_string())
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	var sb strings.Builder
	sb.WriteString(`SELECT 
		grp_cd,
		cd, 
		cd_nm, 
		use_yn, 
		ord	
		FROM tb_cm_cd
			where 1=1`)
	if grp_cd != "" {
		sb.WriteString(" and grp_cd=:grp_cd")
	}
	sb.WriteString(` 
			order by ord`)
	//}

	//Go through rows
	//기본  rows, err := db.Queryx(sb.String())
	//파리미터 파인딩
	//p_m := map[string]interface{}{"city": "Johannesburg"}
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
