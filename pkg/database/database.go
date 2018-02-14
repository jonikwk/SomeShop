package database

import (
	"database/sql"
	"fmt"

	c "../configuration"
	"github.com/fatih/color"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

//GetConnectionString -> Функция получения строки подключения базы данных
func GetConnectionString(config *c.Configuration) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Connect.DBHost, config.Connect.Port, config.User.Login, config.User.Password, config.Connect.DBName)
}

//OpenDB -> Функция открытия подклбчения базы данных
func OpenDB(config *c.Configuration, psqlInfo string) *sql.DB {
	var db = new(sql.DB)
	db, err := sql.Open(config.Connect.DBType, psqlInfo)
	if err != nil {
		glog.Exit(err)
	}
	return db
}

func AddUser(db *sql.DB, id int64) {
	stmt, err := db.Prepare(`insert into tables.users (id) values ($1)`)
	if err != nil {
		glog.Exit()
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		glog.Exit()
	}
}

func IsUserInDatabase(chatID int64, db *sql.DB) bool {
	row := db.QueryRow(`select id from tables.users where id = $1`, chatID)
	var id string
	row.Scan(&id)
	color.Green(id)
	if id == "" {
		return false
	}
	return true
}

func GetRootSection(db *sql.DB) []string {
	rows, err := db.Query(`select title from tables.catalog where id_parent = 0`)
	if err != nil {
		glog.Exit(err)
	}

	sections := make([]string, 0)
	for rows.Next() {
		section := ""
		rows.Scan(&section)
		sections = append(sections, section)
	}
	return sections
}

func GetClothesSection(db *sql.DB) []string {
	rows, err := db.Query(`select title from tables.catalog where id_parent = 1`)
	if err != nil {
		glog.Exit(err)
	}

	sections := make([]string, 0)
	for rows.Next() {
		section := ""
		rows.Scan(&section)
		sections = append(sections, section)
	}
	return sections
}

func GetWomanClothes(db *sql.DB, current int) []string {
	rows, err := db.Query(`select title from tables.catalog where id_parent = 3 limit 5 offset $1`, current)
	if err != nil {
		glog.Exit(err)
	}

	sections := make([]string, 0)
	for rows.Next() {
		section := ""
		rows.Scan(&section)
		sections = append(sections, section)
	}
	return sections
}

func GetClothes(db *sql.DB, current int, id int) []string {
	rows, err := db.Query(`select title from tables.catalog where id_parent = $1 limit 5 offset $2`, id, current)
	if err != nil {
		glog.Exit(err)
	}

	sections := make([]string, 0)
	for rows.Next() {
		section := ""
		rows.Scan(&section)
		sections = append(sections, section)
	}
	return sections
}

func GetManClothes(db *sql.DB) []string {
	rows, err := db.Query(`select title from tables.catalog where id_parent = 4`)
	if err != nil {
		glog.Exit(err)
	}

	sections := make([]string, 0)
	for rows.Next() {
		section := ""
		rows.Scan(&section)
		sections = append(sections, section)
	}
	return sections
}

func GetCurrentItem(db *sql.DB, chatID int64) int {
	row := db.QueryRow(`select current_offset from tables.users where id = $1 `, chatID)
	var current int
	row.Scan(&current)
	return current
}

func SetCurrentItem(db *sql.DB, current int, chatID int64) {
	color.Red(fmt.Sprintln("Current: ", current))
	stmt, err := db.Prepare(`update tables.users set current_offset = $1  where id = $2`)
	if err != nil {
		glog.Exit()
	}
	_, err = stmt.Exec(current, chatID)
	if err != nil {
		glog.Exit()
	}
}

func GetCurrentParnetId(db *sql.DB, chatID int64) int {
	row := db.QueryRow(`select id_current from tables.users where id = $1 `, chatID)
	var current int
	row.Scan(&current)
	return current
}

func SetCurrentParnetId(db *sql.DB, chatID int64, id int) {
	stmt, err := db.Prepare(`update tables.users set id_current = $1 where id = $2`)
	if err != nil {
		glog.Exit()
	}
	_, err = stmt.Exec(id, chatID)
	if err != nil {
		glog.Exit()
	}
}

func SetCurrentItemByDefault(db *sql.DB, chatID int64) {
	stmt, err := db.Prepare(`update tables.users set current_offset = 0  where id = $1`)
	if err != nil {
		glog.Exit()
	}
	_, err = stmt.Exec(chatID)
	if err != nil {
		glog.Exit()
	}
}

//select count(id_parent) from tables.catalog where id_parent = 3;
/*func GetRecordsCount(db *sql.DB) int { //передаем айди
	row := db.QueryRow(`select count(id_parent) from tables.catalog where id_parent = 3`)
	var count int
	row.Scan(&count)
	return count
}*/

func GetRecordsCount(db *sql.DB, id int) int { //передаем айди
	row := db.QueryRow(`select count(id_parent) from tables.catalog where id_parent = $1`, id)
	var count int
	row.Scan(&count)
	return count
}

func GetCatalogId(db *sql.DB, title string) int { //передаем айди
	row := db.QueryRow(`select id from tables.catalog where title = $1`, title)
	var id int
	row.Scan(&id)
	return id
}

func GetSectionTitle(db *sql.DB, id int) string {
	row := db.QueryRow(`select title from tables.catalog where id = $1 `, id)
	var title string
	row.Scan(&title)
	return title
}
