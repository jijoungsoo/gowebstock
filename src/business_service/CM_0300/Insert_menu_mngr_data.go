package CM_0300

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
	MenuData struct {
		MenuCd     string `json:"menu_cd" binding:"required"`
		MenuNm     string `json:"menu_nm" binding:"required"`
		PrntMenuCd string `json:"prnt_menu_cd" binding:"required"`
		FstOrd     string `json:"fst_ord" binding:"required"`
		SedOrd     string `json:"sed_ord" binding:"required"`
		PgmId      string `json:"pgm_id" binding:"required"`
		MenuKind   string `json:"menu_kind" binding:"required"`
		Rmk        string `json:"rmk"`
	}
)

func Insert_menu_mngr_data(c *gin.Context) {
	var json_tmp []MenuData //array로 감싸서 넘겨야한다. []
	if err := c.ShouldBindJSON(&json_tmp); err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := 0; i < len(json_tmp); i++ {
		tmp := json_tmp[i]
		fmt.Printf("tmp.MenuCd %s\n", tmp.MenuCd)
		fmt.Printf("tmp.MenuNm %s\n", tmp.MenuNm)
		fmt.Printf("tmp.PrntMenuCd %s\n", tmp.PrntMenuCd)
		fmt.Printf("tmp.FstOrd %s\n", tmp.FstOrd)
		fmt.Printf("tmp.SedOrd %s\n", tmp.SedOrd)
		fmt.Printf("tmp.PgmId %s\n", tmp.PgmId)
		fmt.Printf("tmp.MenuKind %s\n", tmp.MenuKind)
		fmt.Printf("tmp.Rmk %s\n", tmp.Rmk)
		insert_menu_mngr_data(tmp.MenuCd, tmp.MenuNm, tmp.PrntMenuCd, tmp.FstOrd, tmp.SedOrd, tmp.PgmId, tmp.MenuKind, tmp.Rmk)

	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
func insert_menu_mngr_data(menu_cd string, menu_nm string, prnt_menu_cd string, fst_ord string, sed_ord string, pgm_id string, menu_kind string, rmk string) {
	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`insert into tb_cm_menu
		(
			menu_cd,
			menu_nm,
			prnt_menu_cd,
			fst_ord,
			sed_ord,
			pgm_id,
			menu_kind,
			rmk,
			crt_dtm,
			updt_dtm
		)
		values
		(
			:menu_cd,
			:menu_nm,
			:prnt_menu_cd,
			:fst_ord,
			:sed_ord,
			:pgm_id,
			:menu_kind,
			:rmk,
			now(),
			now()
		)
		ON CONFLICT (menu_cd) 
		DO
		UPDATE
		SET  menu_nm=EXCLUDED.menu_nm, 
		prnt_menu_cd=EXCLUDED.prnt_menu_cd, 
		fst_ord=EXCLUDED.fst_ord, 
		sed_ord=EXCLUDED.sed_ord, 
		pgm_id=EXCLUDED.pgm_id, 
		menu_kind=EXCLUDED.menu_kind, 
		rmk=EXCLUDED.rmk, 
		updt_dtm=now()
		`)
	p_m := map[string]interface{}{}
	p_m["menu_cd"] = menu_cd
	p_m["menu_nm"] = menu_nm
	p_m["prnt_menu_cd"] = prnt_menu_cd
	p_m["fst_ord"] = fst_ord
	p_m["sed_ord"] = sed_ord
	p_m["pgm_id"] = pgm_id
	p_m["menu_kind"] = menu_kind
	p_m["rmk"] = rmk

	_, err = db.NamedExec(sb.String(), p_m)
	bs.CheckError(err)
}
