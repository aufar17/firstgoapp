package models

import (
  "github.com/goravel/framework/database/orm"
)

type Post struct {
  orm.Model
  Title   string
  Body string
  Author string
}
