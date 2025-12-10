package models

import "github.com/goravel/framework/database/orm"

type Post struct {
    orm.Model
    Id   string `json:"id"`
    Title   string `json:"title"`
    Body    string `json:"body"`
    Author  string `json:"author"`
}
