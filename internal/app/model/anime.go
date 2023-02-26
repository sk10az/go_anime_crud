package model

type IdAC uint8

type AnimeCharacter struct {
	Id          IdAC   `json:"id" xml:"id" form:"id" query:"id"`
	Name        string `json:"name" xml:"name" form:"name" query:"name"`
	Age         uint8  `json:"age" xml:"age" form:"age" query:"age"`
	Description string `json:"desc" xml:"desc" form:"desc" query:"desc"`
}
