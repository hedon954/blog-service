package model

import "blog-service/pkg/app"

//标签
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (a Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List []*Tag
	Page *app.Pager
}
