package CM_0100

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	bs "business_service"
)

type (
	// transformedTodo represents a formatted todo
	PgmData struct {
		PgmId    string `json:"pgm_id" binding:"required"`
		PgmNm    string `json:"pgm_nm" binding:"required"`
		PgmLink  string `json:"pgm_link" binding:"required"`
		Category string `json:"category" binding:"required"`
		Rmk      string `json:"rmk"`
	}
)

func Insert_pgm_mngr_data(c *gin.Context) {
	var json_tmp []PgmData //array로 감싸서 넘겨야한다. []
	if err := c.ShouldBindJSON(&json_tmp); err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := 0; i < len(json_tmp); i++ {
		tmp := json_tmp[i]
		fmt.Printf("tmp.PgmId %s\n", tmp.PgmId)
		fmt.Printf("tmp.PgmNm %s\n", tmp.PgmNm)
		fmt.Printf("tmp.PgmLink %s\n", tmp.PgmLink)
		fmt.Printf("tmp.Category %s\n", tmp.Category)
		fmt.Printf("tmp.Rmk %s\n", tmp.Rmk)
		insert_pgm_mngr_data(tmp.PgmId, tmp.PgmNm, tmp.PgmLink, tmp.Category, tmp.Rmk)

	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
func insert_pgm_mngr_data(pgm_id string, pgm_nm string, pgm_link string, category string, rmk string) {
	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`insert into tb_cm_pgm
		(
			pgm_id,
			pgm_nm,
			pgm_link,
			category,
			rmk,
			crt_dtm,
			updt_dtm
		)
		values
		(
			:pgm_id,
			:pgm_nm,
			:pgm_link,
			:category,
			:rmk,
			now(),
			now()
		)
		ON CONFLICT (pgm_id) 
		DO
		UPDATE
		SET  pgm_nm=EXCLUDED.pgm_nm, 
		pgm_link=EXCLUDED.pgm_link, 
		category=EXCLUDED.category, 
		rmk=EXCLUDED.rmk, 
		updt_dtm=now()
		`)
	p_m := map[string]interface{}{}
	p_m["pgm_id"] = pgm_id
	p_m["pgm_nm"] = pgm_nm
	p_m["pgm_link"] = pgm_link
	p_m["category"] = category
	p_m["rmk"] = rmk

	_, err = db.NamedExec(sb.String(), p_m)
	bs.CheckError(err)

}
