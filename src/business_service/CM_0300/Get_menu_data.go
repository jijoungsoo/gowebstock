package CM_0300

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	bs "business_service"
	"encoding/json"
)

type MenuJson struct {
	PgmId    string    `json:"pgm_id"`
	MenuCd   string    `json:"menu_cd"`
	MenuNm   string    `json:"menu_nm"`
	FstOrd   string    `json:"fst_ord"`
	MenuKind string    `json:"menu_kind"`
	Child    []SedMenu `json:"child"`
}

type SedMenu struct {
	PgmId    string `json:"pgm_id"`
	MenuCd   string `json:"menu_cd"`
	MenuNm   string `json:"menu_nm"`
	SedOrd   string `json:"sed_ord"`
	MenuKind string `json:"menu_kind"`
}

func Get_menu_data() []MenuJson {

	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`
	with vw00 as (
		select 
		a.menu_nm,
		a.menu_cd,
		a.menu_kind,
		a.fst_ord,
		a.pgm_id,
		(	
				select  JSON_agg(json_build_object(
					'menu_nm',menu_nm,
					'menu_cd',menu_cd,
					'menu_kind',menu_kind,
					'sed_ord',sed_ord,
					'pgm_id',pgm_id
				) order by b.sed_ord)
				from tb_cm_menu b
				where  b.prnt_menu_cd=a.menu_cd		
		) t
		from tb_cm_menu  a
		where 
		a.prnt_menu_cd ='' 
		order by a.fst_ord
		)
		,vw01 as (
		select JSON_agg(json_build_object(
			'menu_nm',menu_nm,
			'menu_cd',menu_cd,
			'menu_kind',menu_kind,
			'fst_ord',fst_ord,
			'child',t
		) order by fst_ord
		) t from vw00
		)
		select jsonb_pretty(t::jsonb) menu_json  from vw01`)
	rows, err := db.Query(sb.String())
	var menu_json string
	for rows.Next() {
		err = rows.Scan(&menu_json)
	}
	bs.CheckError(err)

	var menu_map []MenuJson
	fmt.Printf("menu_json %s\n", menu_json)
	err = json.Unmarshal([]byte(menu_json), &menu_map)

	fmt.Printf("mem %s\n", menu_map)

	return menu_map

}

func Get_menu_data_json() string {

	db, err := sqlx.Connect("postgres", bs.Get_connection_string())
	bs.CheckError(err)
	var sb strings.Builder

	sb.WriteString(`
	with vw00 as (
		select 
		a.menu_nm,
		a.menu_cd,
		a.menu_kind,
		a.fst_ord,
		a.pgm_id,
		(	
				select  JSON_agg(json_build_object(
					'menu_nm',menu_nm,
					'menu_cd',menu_cd,
					'menu_kind',menu_kind,
					'sed_ord',sed_ord,
					'pgm_id',pgm_id
				) order by b.sed_ord)
				from tb_cm_menu b
				where  b.prnt_menu_cd=a.menu_cd		
		) t
		from tb_cm_menu  a
		where 
		a.prnt_menu_cd ='' 
		order by a.fst_ord
		)
		,vw01 as (
		select JSON_agg(json_build_object(
			'menu_nm',menu_nm,
			'menu_cd',menu_cd,
			'menu_kind',menu_kind,
			'fst_ord',fst_ord,
			'child',t
		) order by fst_ord
		) t from vw00
		)
		select jsonb_pretty(t::jsonb) menu_json  from vw01`)
	rows, err := db.Query(sb.String())
	var menu_json string
	for rows.Next() {
		err = rows.Scan(&menu_json)
	}
	bs.CheckError(err)

	return menu_json

}
