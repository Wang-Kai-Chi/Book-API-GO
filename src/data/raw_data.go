package data

type RawData interface {
	RawBook | []RawBook | RawDvd | []RawDvd | RawCd | []RawCd
}

type RawBook struct {
	Isbn            string
	Title           string
	PublicationDate string `json:"publicationDate"`
	Publisher       string
	Price           string
	Author          string
	Translator      string
	Language        string
}
type RawDvd struct {
	Barcode         string
	Title           string
	PublicationDate string `json:"publicationDate"`
	Price           string
	Publisher       string
	Description     string
}
type RawCd struct {
	Barcode         string
	Title           string
	PublicationDate string `json:"publicationDate"`
	Price           string
	Publisher       string
	Description     string
}
