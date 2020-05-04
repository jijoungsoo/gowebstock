package CM_0100

import (
	"fmt"
	"net/http"
	"strings"

	bs "business_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type (
	// transformedTodo represents a formatted todo
	DeletePgmData struct {
		PgmId string `json:"pgm_id" binding:"required"`
	}
)

func Delete_pgm_mngr_data(c *gin.Context) {
	var json_tmp []DeletePgmData //array로 감싸서 넘겨야한다. []
	if err := c.ShouldBindJSON(&json_tmp); err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(json_tmp) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paramter 0"})
		return
	}
	var arr_pgm_id []string
	arr_pgm_id = make([]string, len(json_tmp), 3)
	for i := 0; i < len(json_tmp); i++ {
		tmp := json_tmp[i]
		fmt.Printf("tmp.PgmId %s\n", tmp.PgmId)
		arr_pgm_id[i] = tmp.PgmId
	}

	arg := map[string]interface{}{
		"pgm_id": arr_pgm_id,
	}
	delete_pgm_mngr_data(arg)
	fmt.Printf("tmp.arg %s\n", arg)

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func delete_pgm_mngr_data(p_arg map[string]interface{}) {
	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	err = db.Ping()
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`delete from tb_cm_pgm where pgm_id IN (:pgm_id)`)
	//_, err = db.NamedExec(sb.String(), p_arg)
	query, args, err := sqlx.Named(sb.String(), p_arg)
	fmt.Printf("%s\n", query)
	query, args, err2 := sqlx.In(query, args...)
	bs.CheckError(err2)
	fmt.Printf("%s\n", query)
	query = db.Rebind(query)
	db.Query(query, args...)
	/*잘된다. */

}
