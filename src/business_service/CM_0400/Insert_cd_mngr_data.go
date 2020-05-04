package CM_0400

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
	CdData struct {
		GrpCd string `json:"grp_cd" binding:"required"`
		Cd    string `json:"cd" binding:"required"`
		CdNm  string `json:"cd_nm" binding:"required"`
		UseYn string `json:"use_yn" binding:"required"`
		Ord   string `json:"ord" binding:"required"`
		Rmk   string `json:"rmk"`
	}
)

func Insert_cd_mngr_data(c *gin.Context) {
	var json_tmp []CdData //array로 감싸서 넘겨야한다. []
	if err := c.ShouldBindJSON(&json_tmp); err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := 0; i < len(json_tmp); i++ {
		tmp := json_tmp[i]
		fmt.Printf("tmp.GrpCd %s\n", tmp.GrpCd)
		fmt.Printf("tmp.Cd %s\n", tmp.Cd)
		fmt.Printf("tmp.CdNm %s\n", tmp.CdNm)
		fmt.Printf("tmp.UseYn %s\n", tmp.UseYn)
		fmt.Printf("tmp.Ord %s\n", tmp.Ord)
		fmt.Printf("tmp.Rmk %s\n", tmp.Rmk)
		insert_cd_mngr_data(tmp.GrpCd, tmp.Cd, tmp.CdNm, tmp.UseYn, tmp.Ord, tmp.Rmk)

	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
func insert_cd_mngr_data(grp_cd string, cd string, cd_nm string, use_yn string, ord string, rmk string) {
	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`insert into tb_cm_cd
		(
			grp_cd,
			cd,
			cd_nm,
			use_yn,
			ord,
			rmk,
			crt_dtm,
			updt_dtm
		)
		values
		(
			:grp_cd,
			:cd,
			:cd_nm,
			:use_yn,
			:ord,
			:rmk,			
			now(),
			now()
		)
		ON CONFLICT (grp_cd,cd) 
		DO
		UPDATE
		SET  cd_nm=EXCLUDED.cd_nm, 
		use_yn=EXCLUDED.use_yn, 
		ord=EXCLUDED.ord, 
		rmk=EXCLUDED.rmk,
		updt_dtm=now()
		`)
	p_m := map[string]interface{}{}
	p_m["grp_cd"] = grp_cd
	p_m["cd"] = cd
	p_m["cd_nm"] = cd_nm
	p_m["use_yn"] = use_yn
	p_m["ord"] = ord
	p_m["rmk"] = rmk

	_, err = db.NamedExec(sb.String(), p_m)
	bs.CheckError(err)
}
