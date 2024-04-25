# Golang based Book API
A RESTful API example for review book, where an author can post their book and see what people think of their book. Meanwhile a user can make a review for book that has been published.

## Installation & Run
Before running this app, you should set the env variable
```bash
# copy the template for .env is from .env.example, then you can change your env variable in .env
cp .env.example .env
```

## API

#### /categories
* `GET` : Get all categories
* `POST` : Create a new category

#### /categories/:id
* `GET` : Get a category
* `PUT` : Update a category
* `DELETE` : Delete a category

##### example req body for categories
* `POST` or `PUT` : 
```json
{
  "name" : "novel",
}
```

#### /books
* `GET` : Get all books
* `POST` : Create a new book

#### /books/:id
* `GET` : Get a book
* `PUT` : Update a book
* `DELETE` : Delete a book

##### example req body for books
* `POST` or `PUT` : 
```json
{
  "title" : "Cantik Itu Luka",
  "description": "Award winning book by Eka Kurniawan",
  "image_url": "https://www.indonesiana.id/images/all/2016/08/03/700-cil2015.jpg",
  "release_year": 2021,
  "price": 200000,
  "total_page": 1000,
  "author": "Eka Kurniawan",
  "category_id": 6
}
```

#### /users
* `GET` : Get all users
* `POST` : Create a new user

#### /users/:id
* `GET` : Get a user
* `PUT` : Update a user
* `DELETE` : Delete a user

##### example req body for users
* `POST` or `PUT` : 
```json
{
  "username" : "cella",
  "password": "12345",
  "role": "reviewer"
}
```

#### /review
* `GET` : Get all review
* `POST` : Create a new review

#### /review/:id
* `GET` : Get a review
* `PUT` : Update a review
* `DELETE` : Delete a review

##### example req body for review
* `POST` or `PUT` : 
```json
{
  "user_id" : 420,
  "book_id": 69,
  "description": "Well written book",
  "stars": 5
}
```

## Role Based Access

#### Public routes
* `/books` : `GET`
* `/categories` : `GET`
* `/users` : `GET`
* `/reviews` : `GET`

#### user with role `superadmin`
* `/books` : `POST` , `PUT` and `Delete`
* `/categories` : `POST` , `PUT` and `Delete`
* `/users` : `POST` , `PUT` and `Delete`
* `/reviews` : `POST` , `PUT` and `Delete`

#### user with role `author`
* `/books` : `POST` , `PUT` and `Delete`
* `/users` : `POST` , `PUT` and `Delete`

#### user with role `reviewer`
* `/users` : `POST` , `PUT` and `Delete`
* `/reviews` : `POST` , `PUT` and `Delete`

 