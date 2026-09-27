package model

// Book merepresentasikan data buku perpustakaan
type Book struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Code        string `json:"code" gorm:"unique;not null"`
	Title       string `json:"title" gorm:"not null"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	PublishYear int    `json:"publish_year"`
	Stock       int    `json:"stock"`
}

// Pakai nama tabel sendiri biar nggak bentrok sama tabel "books" punya Prisma
func (Book) TableName() string {
	return "buku_go"
}
