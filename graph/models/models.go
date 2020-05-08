package models

type Author struct {
	ID        string `json:"id" gorm:"primary_key"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	BookID    int    `json:"bookID"`
}

type Book struct {
	ID       string    `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Author   []*Author `json:"author"`
}
