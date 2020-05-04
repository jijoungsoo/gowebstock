package CM_0200

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
	DomainData struct {
		DomainCd string `json:"domain_cd" binding:"required"`
		DomainNm string `json:"domain_nm" binding:"required"`
		DataType string `json:"data_type" binding:"required"`
		Rmk      string `json:"rmk"`
	}
)

func Insert_domain_mngr_data(c *gin.Context) {
	var json_tmp []DomainData //array로 감싸서 넘겨야한다. []
	if err := c.ShouldBindJSON(&json_tmp); err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := 0; i < len(json_tmp); i++ {
		tmp := json_tmp[i]
		fmt.Printf("tmp.DomainCd %s\n", tmp.DomainCd)
		fmt.Printf("tmp.DomainNm %s\n", tmp.DomainNm)
		fmt.Printf("tmp.DataType %s\n", tmp.DataType)
		fmt.Printf("tmp.Rmk %s\n", tmp.Rmk)
		insert_domain_mngr_data(tmp.DomainCd, tmp.DomainNm, tmp.DataType, tmp.Rmk)

	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
func insert_domain_mngr_data(domain_cd string, domain_nm string, data_type string, rmk string) {
	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`insert into tb_cm_domain
		(
			domain_cd,
			domain_nm,
			data_type,
			rmk,
			crt_dtm,
			updt_dtm
		)
		values
		(
			:domain_cd,
			:domain_nm,
			:data_type,
			:rmk,
			now(),
			now()
		)
		ON CONFLICT (domain_nm) 
		DO
		UPDATE
		SET  domain_cd=EXCLUDED.domain_cd, 
		data_type=EXCLUDED.data_type, 
		rmk=EXCLUDED.rmk, 
		updt_dtm=now()
		`)
	p_m := map[string]interface{}{}
	p_m["domain_cd"] = domain_cd
	p_m["domain_nm"] = domain_nm
	p_m["data_type"] = data_type
	p_m["rmk"] = rmk

	_, err = db.NamedExec(sb.String(), p_m)
	bs.CheckError(err)

}
