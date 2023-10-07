package models

type HTML struct {
	Content string `json:"content" form:"content" query:"content" validate:"required"`
}

type URL struct {
	Link string `json:"link" form:"link" query:"link" validate:"required,url"`
}

type PDF struct {
	Content string `json:"content" form:"content" query:"content" validate:"required"`
}
