package controllers

import (
	"strconv"
	"bufio"

	 "encoding/json"

    "github.com/goravel/framework/contracts/http"
    "github.com/goravel/framework/facades"
	"goravel/app/models"
)


type PostRow struct {
    ID        uint   `json:"id"`
    Title     string `json:"title"`
    Body      string `json:"body"`
    Author    string `json:"author"`
    CreatedAt string `json:"created_at"`
}

type PostController struct {
}

func NewPostController() *PostController {
	return &PostController{}
}


func (r *PostController) Index(ctx http.Context) http.Response {
	var rows []PostRow

	err := facades.DB().Select(
		&rows,
		"SELECT id, title, body, author, created_at FROM posts",
	)

	if err != nil {
		return ctx.Response().Json(500, map[string]any{
			"error": err.Error(),
		})
	}

	// stream hasil ke client (pakai bufio writer untuk performa)
	writer := ctx.Response().Writer()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)

	bw := bufio.NewWriter(writer)
	defer bw.Flush()

	// buka array JSON
	bw.WriteByte('[')

	for i, row := range rows {
		if i != 0 {
			bw.WriteByte(',')
		}

		// encode tiap row (map) jadi JSON
		b, jErr := json.Marshal(row)
		if jErr != nil {
			// kalau marshal gagal, skip atau return error
			_ = jErr
			continue
		}
		bw.Write(b)
	}

	// tutup array JSON
	bw.WriteByte(']')

	// kembalikan nil karena respon sudah ditulis langsung
	return nil
}





// Create: Buat post baru
func (r *PostController) Create(ctx http.Context) http.Response {
	title := ctx.Request().Input("title")
	body := ctx.Request().Input("body")
	author := ctx.Request().Input("author")

	if title == "" || body == "" || author == "" {
		return ctx.Response().Json(400, map[string]any{
			"error": "title, body, and author are required",
		})
	}

	post := models.Post{
		Title:  title,
		Body:   body,
		Author: author,
	}

	if err := facades.Orm().Query().Create(&post); err != nil {
		return ctx.Response().Json(500, map[string]any{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(201, map[string]any{
		"post": post,
	})
}

// Update: Update post berdasarkan ID
func (r *PostController) Update(ctx http.Context) http.Response {
	idStr := ctx.Request().Input("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Response().Json(400, map[string]any{
			"error": "invalid id",
		})
	}

	var post models.Post
	if facades.Orm().Query().Find(&post, id); post.ID == 0 {
		return ctx.Response().Json(404, map[string]any{
			"error": "post not found",
		})
	}

	title := ctx.Request().Input("title")
	body := ctx.Request().Input("body")
	author := ctx.Request().Input("author")

	if title != "" {
		post.Title = title
	}
	if body != "" {
		post.Body = body
	}
	if author != "" {
		post.Author = author
	}

	if err := facades.Orm().Query().Save(&post); err != nil {
		return ctx.Response().Json(500, map[string]any{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(200, map[string]any{
		"post": post,
	})
}

func (r *PostController) Delete(ctx http.Context) http.Response {
	idStr := ctx.Request().Input("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Response().Json(400, map[string]any{
			"error": "invalid id",
		})
	}

	var post models.Post
	if facades.Orm().Query().Find(&post, id); post.ID == 0 {
		return ctx.Response().Json(404, map[string]any{
			"error": "post not found",
		})
	}

	// Hapus post
	result, err := facades.Orm().Query().Delete(&post)
	if err != nil {
		return ctx.Response().Json(500, map[string]any{
			"error": err.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return ctx.Response().Json(404, map[string]any{
			"error": "post not found or already deleted",
		})
	}

	return ctx.Response().Json(200, map[string]any{
		"message": "post deleted successfully",
	})
}