package user

import (
	
)

type User struct {
	ID    string
	Name  string
	Books []Book
}

func (u *User) BorrowBook(bookID string, books []Book) error {

	// olmoqchi bo'lgan kitobi bo'lmasa, "bunday kitob mavjud emas degan" error qaytarsin

	return nil
}

func (u *User) ReturnBook(bookID string)(Book,  error) {

	// o'zini kitoblari ichidan olmoqchi bo'lgan kitobi bo'lmasa error qaytarsin "sizda bunday kitob yoq"

	return Book, nil
}

func (u *User) DisplayBooks() (string, error) {
	books := "Books Borrowed:\n"  // bitta stringa hamma kitoblarni ketma ket qilib joylang
	/*
		1. Xalqa, Akrom Malik, 2019
		2. Xalqa, Akrom Malik, 2019
		3. Xalqa, Akrom Malik, 2019
	*/

	//  olingan kitoblarni qaytaradi, agar kitoblari bo'lmasa error qaytarsin "sizda olingan kitoblar yoq"

	return books, nil
}
