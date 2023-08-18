package main

type RawDataConvertor struct {
	getSingleBook Book
	getBooks      []Book

	getSingleDvd Dvd
	getDvds      []Dvd

	getSingleCd Cd
	getCds      []Cd
}

func NewRawDataConvertor(raw RawDataLoader) RawDataConvertor {

	singleBook := func(raw RawBook) Book {
		return Book{
			Product_: Product{
				Barcode:       raw.Isbn,
				Price:         raw.Price,
				Product_title: raw.Title,
			},
			PublicationDate: raw.PublicationDate,
			Author:          raw.Author,
			Translator:      raw.Translator,
			Language:        raw.Language,
		}
	}

	books := func(raws []RawBook) []Book {
		books := make([]Book, len(raws))
		for i, v := range raws {
			books[i] = singleBook(v)
		}
		return books
	}

	dvd := func(raw RawDvd) Dvd {
		return Dvd{
			Product_: Product{
				Barcode:       raw.Barcode,
				Price:         raw.Price,
				Product_title: raw.Title,
			},
			PublicationDate: raw.PublicationDate,
			Publisher:       raw.Publisher,
		}
	}

	dvds := func(raws []RawDvd) []Dvd {
		dvds := make([]Dvd, len(raws))

		for i, v := range raws {
			dvds[i] = dvd(v)
		}
		return dvds
	}

	cd := func(raw RawCd) Cd {
		return Cd{
			Product_: Product{
				Barcode:     raw.Barcode,
				Price:       raw.Price,
				Description: raw.Description,
			},
			PublicationDate: raw.PublicationDate,
			Publisher:       raw.Publisher,
		}
	}

	cds := func(raws []RawCd) []Cd {
		cds := make([]Cd, len(raws))

		for i, v := range raws {
			cds[i] = cd(v)
		}
		return cds
	}

	return RawDataConvertor{
		getSingleBook: singleBook(raw.getSingleRawBook),
		getBooks:      books(raw.getRawBooks),

		getSingleDvd: dvd(raw.getSingleRawDvd),
		getDvds:      dvds(raw.getRawDvds),

		getSingleCd: cd(raw.getSingleRawCd),
		getCds:      cds(raw.getRawCds),
	}
}
