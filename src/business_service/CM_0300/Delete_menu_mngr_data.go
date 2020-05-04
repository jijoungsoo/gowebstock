package CM_0300

import (
	"fmt"
	"net/http"
	"strings"

	bs "business_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Delete_menu_mngr_data(c *gin.Context) {

	type (
		// transformedTodo represents a formatted todo
		MenuData struct {
			MenuCd string `json:"menu_cd" binding:"required"`
		}
	)
	var json_tmp MenuData //array로 감싸서 넘겨야한다. []
	if err := c.ShouldBindJSON(&json_tmp); err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json_tmp == (MenuData{}) { /*빈값 체크라고 어디서 보았다. */
		/*https://play.golang.org/p/RXcE06chxE*/
		/*https://stackoverflow.com/questions/20240179/nil-detection-in-go*/
		c.JSON(http.StatusBadRequest, gin.H{"error": "paramter 0"})
		return
	}
	var menu_cd = json_tmp.MenuCd

	delete_menu_mngr_data(menu_cd)
	fmt.Printf("tmp.arg %s\n", menu_cd)

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func delete_menu_mngr_data(menu_cd string) {
	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	err = db.Ping()
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`delete from tb_cm_menu where menu_cd =:menu_cd`)
	_, err = db.NamedExec(sb.String(), menu_cd)
	bs.CheckError(err)
}
