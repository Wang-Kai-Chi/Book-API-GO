package main

func DefaultQuantity() int {
	return 1
}

type RawDataConvertor[T Data] interface {
	ConvertRaw() T
	ConvertRaws() T
}

func (book Book) ConvertRaw(raw RawBook) Book {
	return Book{
		Product_: Product{
			Barcode:          raw.Isbn,
			Price:            raw.Price,
			Product_title:    raw.Title,
			Quantity:         DefaultQuantity(),
			Publication_date: raw.Publication_date,
			Publisher:        raw.Publisher,
			Description:      "",
		},
		Author:     raw.Author,
		Translator: raw.Translator,
		Language:   raw.Language,
	}
}
func (book Book) ConvertRaws(raws []RawBook) []Book {
	books := make([]Book, len(raws))
	for i, v := range raws {
		books[i] = book.ConvertRaw(v)
	}
	return books
}

func (dvd Dvd) ConvertRaw(raw RawDvd) Dvd {
	return Dvd{
		Product_: Product{
			Barcode:          raw.Barcode,
			Price:            raw.Price,
			Product_title:    raw.Title,
			Quantity:         DefaultQuantity(),
			Publication_date: raw.Publication_date,
			Publisher:        raw.Publisher,
			Description:      raw.Description,
		},
		Director: "",
		Category: "",
	}
}

func (dvd Dvd) ConvertRaws(raws []RawDvd) []Dvd {
	dvds := make([]Dvd, len(raws))
	for i, v := range raws {
		dvds[i] = dvd.ConvertRaw(v)
	}
	return dvds
}

func (cd Cd) ConvertRaw(raw RawCd) Cd {
	return Cd{
		Product_: Product{
			Barcode:          raw.Barcode,
			Price:            raw.Price,
			Description:      raw.Description,
			Quantity:         DefaultQuantity(),
			Publication_date: raw.Publication_date,
			Publisher:        raw.Publisher,
			Product_title:    raw.Title,
		},
		Performer: "",
		Genre:     "",
	}
}

func (cd Cd) ConvertRaws(raws []RawCd) []Cd {
	cds := make([]Cd, len(raws))
	for i, v := range raws {
		cds[i] = cd.ConvertRaw(v)
	}
	return cds
}
