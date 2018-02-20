package models

//Users -> Пердсавляет структуру пользователя
type Users struct {
	ID       int64
	Username string
	Phone    string
	FullName string
	Address  string
}

type Description struct {
	Title       string
	Price       string
	Color       string
	Description string
	Photo       string
}
