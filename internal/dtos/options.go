package dtos

type ListOpts struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
	Filter any `form:"filter"`
}
