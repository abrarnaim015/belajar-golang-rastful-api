package web

type  CategoryCreateReq struct {
	Name string `validate:"required,max=200,min=1" json:"name"`
}

