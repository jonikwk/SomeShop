package database

import (
	"database/sql"
	"fmt"

	c "../configuration"
	"../models"
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
	stmt, err := db.Prepare(`update tables.users set id_current = $1, current_offset = 0 where id = $2`)
	if err != nil {
		glog.Exit()
	}
	_, err = stmt.Exec(id, chatID)
	if err != nil {
		glog.Exit()
	}
}

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
	color.Red(fmt.Sprintln("ID: ", id))
	return id
}

func GetSectionTitle(db *sql.DB, id int) string {
	row := db.QueryRow(`select title from tables.catalog where id = $1 `, id)
	var title string
	row.Scan(&title)
	return title
}

func GetParentID(db *sql.DB, id int) int {
	row := db.QueryRow(`select id_parent from tables.catalog where id = $1 `, id)
	var parentID int
	row.Scan(&parentID)
	return parentID
}

func GetCatalogIDSameSections(db *sql.DB, chatID int64, section string) int {
	row := db.QueryRow(`select tables.catalog.id from tables.catalog 
		inner join tables.users on tables.catalog.id_parent=tables.users.id_current 
		where title = $1 and tables.users.id = $2`, section, chatID)
	var id int
	row.Scan(&id)
	color.Red(fmt.Sprintln("ID: ", id))
	return id
}

func GetItems(db *sql.DB, id int, offset int) []*models.Description {
	color.Green(fmt.Sprintln("ID IN GET ITEMS: ", id))
	color.Green(fmt.Sprintln("OFFSET BLYAD: ", offset))
	rows, err := db.Query(`select title, price, color, description, photo from tables.products
		 where id_category = $1 limit 5 offset $2`, id, offset)
	if err != nil {
		glog.Exit(err)
	}
	items := make([]*models.Description, 0)
	for rows.Next() {
		item := new(models.Description)
		rows.Scan(&item.Title, &item.Price, &item.Color, &item.Description, &item.Photo)
		color.Red(fmt.Sprintln("ITEM: ", item.Title))
		items = append(items, item)
	}
	return items
}

func GetItemsCount(db *sql.DB, id int) int { //передаем айди
	row := db.QueryRow(`select count(id) from tables.products where id_category = $1`, id)
	var count int
	row.Scan(&count)
	return count
}

func IsUserContainPhoneNumber(db *sql.DB, chatID int64) bool {
	row := db.QueryRow(`select phone from tables.users where id = $1`, chatID)
	var phone string
	row.Scan(&phone)
	color.Red("PHONE: ", phone)
	if phone == "none" || phone == "" {
		return false
	}
	return true
}

func SetUserPhoneNumber(db *sql.DB, chatID int64, phone string) {
	stmt, err := db.Prepare(`update tables.users set phone = $1 where id = $2`)
	if err != nil {
		glog.Exit()
	}
	_, err = stmt.Exec(phone, chatID)
	if err != nil {
		glog.Exit()
	}
}

func IsRegistrationCompleted(db *sql.DB, chatID int64) bool {
	row := db.QueryRow(`select registration_completed from tables.users where id = $1`, chatID)
	var registration bool
	row.Scan(&registration)
	color.Red(fmt.Sprintln(registration))
	return registration
}

func SetUserInformationByDefault(db *sql.DB, chatID int64) {
	stmt, err := db.Prepare(`update tables.users set phone = default, address = default where id = $1`)
	if err != nil {
		glog.Exit()
	}
	_, err = stmt.Exec(chatID)
	if err != nil {
		glog.Exit()
	}
}

func CompleteRegistration(db *sql.DB, chatID int64) {
	stmt, err := db.Prepare(`update tables.users set registration_completed = true  where id = $1`)
	if err != nil {
		glog.Exit()
	}
	_, err = stmt.Exec(chatID)
	if err != nil {
		glog.Exit()
	}
}

func AddAddress(db *sql.DB, chatID int64, address string) {
	stmt, err := db.Prepare(`update tables.users set address = $1  where id = $2`)
	if err != nil {
		glog.Exit()
	}
	_, err = stmt.Exec(address, chatID)
	if err != nil {
		glog.Exit()
	}
}

func GetAddress(db *sql.DB, chatID int64) string {
	row := db.QueryRow(`select address from tables.users where id = $1`, chatID)
	var address string
	row.Scan(&address)
	return address
}

//Пересмотреть
/*func IsGettingAddressTrue(db *sql.DB, chatID int64) bool {
	row := db.QueryRow(`select getting_address from tables.users where id = $1`, chatID)
	var gettingAddress string
	row.Scan(&gettingAddress)
	if gettingAddress == "" {
		return false
	}
	return true
}

//Пересмотреть
func IsGettingAddressCompleted(chatID int64, db *sql.DB) bool {
	row := db.QueryRow(`select getting_address from tables.users where id = $1`, chatID)
	var gettingAddress bool
	row.Scan(&gettingAddress)
	color.Green(fmt.Sprintln(gettingAddress))
	return gettingAddress
}
*/
