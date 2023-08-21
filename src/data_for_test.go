package main

type DataForTest struct {
	IsbnSample    string
	DvdCodeSample string
	CdCodeSample  string
}

func NewDataForTest() DataForTest {
	return DataForTest{
		IsbnSample:    "9789571313887",
		DvdCodeSample: "4715219794386",
		CdCodeSample:  "602508588662",
	}
}
