package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bmviniciuss/golang-rest-api/internal/comment"
	"github.com/google/uuid"
)

type CommentRow struct {
	ID     string         `db:"id"`
	Slug   sql.NullString `db:"slug"`
	Body   sql.NullString `db:"body"`
	Author sql.NullString `db:"author"`
}

func mapCommentRowToComment(cr CommentRow) comment.Comment {
	return comment.Comment{
		ID:   cr.ID,
		Slug: cr.Slug.String,
		Body: cr.Body.String,
	}
}

func (d *Database) GetComment(
	ctx context.Context,
	uuid string,
) (comment.Comment, error) {
	var cr CommentRow
	r := d.Client.QueryRowContext(
		ctx,
		"SELECT id, slug, body, author FROM comments WHERE id = $1",
		uuid,
	)

	err := r.Scan(&cr.ID, &cr.Slug, &cr.Body, &cr.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to scan comment: %w", err)
	}

	return mapCommentRowToComment(cr), nil
}

func (d *Database) CreateComment(ctx context.Context, cmt *comment.Comment) error {
	cmt.ID = uuid.New().String()
	pr := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments (id, slug, author, body) VALUES (:id, :slug, :author, :body)`,
		&pr,
	)

	if err != nil {
		return fmt.Errorf("failed to insert comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return fmt.Errorf("failed to close rows: %w", err)
	}

	fmt.Println("Comment inserted")

	return nil
}

func (d *Database) DeleteComment(ctx context.Context, uuid string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments where id = $1`,
		uuid,
	)

	if err != nil {
		return fmt.Errorf("failed to delete comment from the database: %w", err)
	}

	return nil
}

func (d *Database) UpdateComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	pr := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET slug = :slug, body = :body, author = :author WHERE id = :id`,
		&pr,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to update comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return mapCommentRowToComment(pr), nil

}
