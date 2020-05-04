package main

import (
	"business_service"
	"business_service/CM_0100"
	"business_service/CM_0200"
	"business_service/CM_0300"
	"business_service/CM_0400"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	bs "business_service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var map_url map[string]string

func get_to(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	//views폴더가 자동으로 지정되는 듯하다. view로 폴더이름 했을땐 에러났다.
	//fmt.Println(fmt.Sprintf("%v", c.Request.URL))
	pgm_id := fmt.Sprintf("%v", c.Request.URL)[1:]
	//fmt.Println("aaa")
	fmt.Println(pgm_id)
	//fmt.Println("aaa")
	tmp := map_url[pgm_id]
	fmt.Println(tmp + ".html")
	c.Header("Content-Type", "text/html")
	//views폴더가 자동으로 지정되는 듯하다. view로 폴더이름 했을땐 에러났다.
	uuid1, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	fmt.Println(uuid1)
	c.HTML(http.StatusOK, tmp+".html", gin.H{"uuid": uuid1})
}

func main() {
	map_url = make(map[string]string)
	var pgm_map_info = get_route_map()
	var menu_json = CM_0300.Get_menu_data()
	for _, tmp_map_url := range pgm_map_info {
		//fmt.Println(key)
		//fmt.Println(tmp_map_url)
		fmt.Println(tmp_map_url["pgm_id"])
		fmt.Println(tmp_map_url["pgm_link"])
		map_url[tmp_map_url["pgm_id"].(string)] = tmp_map_url["pgm_link"].(string)
	}
	map_url["stock_order"] = "stock_order"
	map_url["stock_detail"] = "stock_detail"

	r := gin.Default()
	gin.ForceConsoleColor()
	//r.Static("/", "./views")
	r.Static("/src", "./views/src")

	r.HTMLRender = ginview.Default()
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//쿠키용
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/index", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")

		//z쿠키용
		session := sessions.Default(c)
		//var count int
		v := session.Get("id")
		//redirect
		//	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
		//c.Request.URL.Path = "/test2"
		//r.HandleContext(c)
		fmt.Println(fmt.Sprintf("%v", v))
		fmt.Println("aaa11")
		if v == nil {
			fmt.Println("bb")
			//redirect
			//c.Request.URL.Path = "/login"
			//r.HandleContext(c)
			//return
			/*계속 진행된다. 아래 리턴이 꼭 필요한듯 하다. 그런데 url 이 d왜 안 바귈까?
			이러면 자바에 server.forword나 transfer하고 같은건데.  앞에 도메인을 붙이면
			url 주소가 이동하는 리디렉션이 된다.
			도메인 주소 여부에 따라 선택하는 듯하다.
			나는 도메인 주소를 붙이겠다.
			페이지가 새로고침이 잘안되서 그냥 위처럼 했다.
			그런데 캐시가 남아서 잘 안되었다.
			크롬 301 삭제 를 해야한다.
			개발자 도구에서 disable cache도 된다.
			*/
			c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/login")
			c.Abort()
			return
		}
		fmt.Println("cc")
		/*
			session.Set("count", count)
			session.Save()
		*/

		fmt.Println(fmt.Sprintf("%v", c.Request.URL))
		//views폴더가 자동으로 지정되는 듯하다. view로 폴더이름 했을땐 에러났다.
		tmp := map_url[fmt.Sprintf("%v", c.Request.URL)]

		fmt.Println(tmp + ".html")
		c.Header("Content-Type", "text/html")
		//views폴더가 자동으로 지정되는 듯하다. view로 폴더이름 했을땐 에러났다.

		uuid1, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}
		fmt.Println(uuid1)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"pgm_map_info": pgm_map_info,
			"menu_json":    menu_json,
			"id":           v,
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		fmt.Println("ㄱㄱㄱㄱ")
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.POST("/post_logout", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")

		session := sessions.Default(c)

		session.Clear()
		session.Save() //clear한다음에 꼭 save해야 저장됨

		c.JSON(200, gin.H{
			"result": "ok",
		})
	})

	r.POST("/post_login", func(c *gin.Context) {
		//c.Header("Content-Type", "text/html")

		type (
			// transformedTodo represents a formatted todo
			Login struct {
				Id string `json:"id" binding:"required"`
				Pw string `json:"pw" binding:"required"`
			}
		)

		var json_tmp Login //array로 감싸서 넘겨야한다. []
		if err := c.ShouldBindJSON(&json_tmp); err != nil {
			fmt.Printf("%s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("json_tmp.Id %s\n", json_tmp.Id)
		fmt.Printf("json_tmp.Pw %s\n", json_tmp.Pw)

		session := sessions.Default(c)
		var ret string
		_ = ret

		if json_tmp.Id != "" {
			session.Set("id", json_tmp.Id)
			session.Set("pw", json_tmp.Pw)
			session.Save()
			ret = "ok"
		} else {
			ret = "nok"
		}

		c.JSON(200, gin.H{
			"result": ret,
		})
	})

	/*html*/
	/*
		r.GET("/index", get_index)
		r.GET("/optkwfid", get_optkwfid)\
	*/

	/*이 방법이 베스트 였는데 서버를 껐다 켜야하는 단점이 있다.*/
	for key, _ := range map_url {
		//fmt.Println(key, val)
		r.GET(key, get_to)
	}

	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위

	/*r.get은 반복적이니 이걸 for 문으로 돌리려고 했는데 그렇게 했더니 안먹는다.
	  이걸 만약에 프로그램으로 짠다면 인젝션으로 파일을 건드는 형식으로 짜야겠다.
	  다시 시도해보자.
	*/

	/*json 데이터 조회*/
	r.POST("/chejan_balance_data", bs.Post_chejan_balance_data)
	r.POST("/chejan_order_data", bs.Post_chejan_order_data)
	r.POST("/realtime_contract_data", bs.Post_realtime_contract_data)
	r.POST("/opt10085_data", bs.Post_opt10085_data)
	r.POST("/opt10001_data", bs.Post_opt10001_data)
	r.POST("/optkwfid_data", bs.Post_optkwfid_data)
	r.POST("/theme_cd_data", bs.Post_theme_cd_data)

	r.POST("/external_send_order", external_send_order)
	r.POST("/get_api_cm_cd", bs.Get_cm_cd)
	r.POST("/Get_stock_code_info", bs.Get_stock_code_info)

	r.POST("/domain_mngr_data", CM_0200.Get_domain_mngr_data)
	r.POST("/insert_domain_mngr_data", CM_0200.Insert_domain_mngr_data)

	r.POST("/pgm_mngr_data", CM_0100.Get_pgm_mngr_data)
	r.POST("/insert_pgm_mngr_data", CM_0100.Insert_pgm_mngr_data)
	r.POST("/delete_pgm_mngr_data", CM_0100.Delete_pgm_mngr_data)

	r.POST("/menu_mngr_data", CM_0300.Get_menu_mngr_data)
	r.POST("/insert_menu_mngr_data", CM_0300.Insert_menu_mngr_data)
	r.POST("/delete_menu_mngr_data", CM_0300.Delete_menu_mngr_data)

	r.POST("/cd_grp_mngr_data", CM_0400.Get_cd_grp_mngr_data)
	r.POST("/cd_mngr_data", CM_0400.Get_cd_mngr_data)
	r.POST("/insert_cd_grp_mngr_data", CM_0400.Insert_cd_grp_mngr_data)
	r.POST("/insert_cd_mngr_data", CM_0400.Insert_cd_mngr_data)

	r.Run(":8080")

	//아래처럼 두번 호출해도 안먹는다.
	//결국 데이터를 조회하는 프로그램은 별도로 빼고
	//http socket을 통해서 호출하도록 하고
	//router를 따로 관리하도록 하는게 맞는것 같다.
	//비즈엑터에 biz_id 처럼  id는 유일해야하고 !! 그렇게 구성하면 될 것 같다.
	//그러면 ui프레임 워크는  데이터 조회와 무관하게 동작할수 있다.
	// 다른 방법으로는 리플렉터를 이용해서
	//https://stackoverflow.com/questions/18017979/golang-pointer-to-function-from-string-functions-name
	//디비로 보관하고 이걸 라우터에 뿌려주는 방법
	//쉽지 않다. 일단 이대로 가자.

	//r.POST("/delete_pgm_mngr_data", bs.Delete_pgm_mngr_data)
	//r.Run(":8080")
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func get_route_map() []map[string]interface{} {
	db, err := sqlx.Connect("postgres", business_service.Get_connection_string())
	business_service.CheckError(err)
	//err = db.Ping()
	//business_service.CheckError(err)
	var sb strings.Builder
	sb.WriteString(`select pgm_id,pgm_link,pgm_nm from tb_cm_pgm`)
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
	return mm
}
