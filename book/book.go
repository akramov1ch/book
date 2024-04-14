package book

type Book struct {
	ID     string
	Title  string
	Author string
	Year   int
	Cnt    int
}

type Books struct {
	Books []Book
}

func (b *Books) DisplayAllBooks() error {

	// barcha kitoblarni chiqarsin agar bo'sh bo'lsa error type  qaytarsin "kitoblar mavjud emas"

	return nil
}

func (b *Books) AddBooksToShelf(book Book) error {

	// kitoblar javoniga kitob qo'shsin, agar qo'shmoqchi bo'lsan kitobni biror bir fieldi bo'sh bo'lsa,
	// "qo'shmoqchi bo'lgan kitob ma`lumotlari yaroqli emas" degan error qaytarsin

	return nil
}

func (b *Books) RemoveBookFromBooks(bookId string) error {

	// kitoblar ichidan kitobni olib tashlash, agar bunday kitob bo'lmasa,
	// "olib tashlanishi kerak bo'lgan kitob kitoblar ichida mavjuda emas" degan error qaytarsin

	return nil
}
