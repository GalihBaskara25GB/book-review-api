package structs

import "time"

type Book struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       int       `json:"price"`
	TotalPage   int       `json:"total_page"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CategoryId  int       `json:"category_id"`
}

type Category struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Review struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	BookId      int64     `json:"book_id"`
	Description string    `json:"description"`
	Stars       int64     `json:"stars"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookWithCategory struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ImageUrl     string    `json:"image_url"`
	ReleaseYear  int       `json:"release_year"`
	Price        int       `json:"price"`
	TotalPage    int       `json:"total_page"`
	Author       string    `json:"author"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CategoryId   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
}

type ReviewWithUserWithBookWithCategory struct {
	Id              int64     `json:"id"`
	UserId          int64     `json:"user_id"`
	BookId          int64     `json:"book_id"`
	Description     string    `json:"description"`
	Stars           int64     `json:"stars"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	UserUsername    string    `json:"user_username"`
	BookTitle       string    `json:"book_title"`
	BookDescription string    `json:"book_description"`
	BookImageUrl    string    `json:"book_image_url"`
	BookReleaseYear int       `json:"book_release_year"`
	BookPrice       int       `json:"book_price"`
	BookTotalPage   int       `json:"book_total_page"`
	BookAuthor      string    `json:"book_author"`
	CategoryId      int       `json:"category_id"`
	CategoryName    string    `json:"category_name"`
}
