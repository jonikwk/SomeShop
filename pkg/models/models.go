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
	ID          int
	Title       string
	Price       string
	Color       string
	Description string
	Photo       string
}

type Order struct {
	Title    string
	Price    int
	Size     string
	Color    string
	Photo    string
	Quantity int
}
