package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251209063211CreatePostsTable struct{}

// Signature The unique signature for the migration.
func (r *M20251209063211CreatePostsTable) Signature() string {
	return "20251209063211_create_posts_table"
}

// Up Run the migrations.
func (r *M20251209063211CreatePostsTable) Up() error {
	if !facades.Schema().HasTable("posts") {
		return facades.Schema().Create("posts", func(table schema.Blueprint) {
			table.ID()
			table.String("title")
			table.Text("body")
			table.String("author")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20251209063211CreatePostsTable) Down() error {
 	return facades.Schema().DropIfExists("posts")
}
