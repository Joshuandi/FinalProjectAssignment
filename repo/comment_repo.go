package repo

import (
	"FinalProjectAssignment/config"
	"FinalProjectAssignment/model"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type CommentRepoInterface interface {
	CommentRepoRegister(ctx context.Context, comments *model.Comment) (*model.Comment, error)
	CommentRepoUpdate(ctx context.Context, id string, comments *model.Comment) (*model.CommentShow, error)
	CommentRepoDelete(ctx context.Context, id string, comments *model.Comment) (*model.Comment, error)
	CommentRepoGet(ctx context.Context, id string) ([]*model.CommentGet, error)
}

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) CommentRepoInterface {
	return &CommentRepo{db: db}
}

func (c *CommentRepo) CommentRepoRegister(ctx context.Context, comments *model.Comment) (*model.Comment, error) {
	sqlSt := `insert into commentss (c_message, photo_id, user_id, c_created_date)
	values ($1, $2, $3, $4)
	returning c_id;`
	err := config.Db.QueryRow(sqlSt,
		comments.Message,
		comments.Photo_id,
		comments.User_id,
		time.Now(),
	).Scan(&comments.Comment_id)
	fmt.Println("ini Comment Repo : ", comments)
	if err != nil {
		fmt.Println("Query row error")
	}
	fmt.Println("repo Comment:", comments)
	fmt.Println("repo Comment User_id:", comments.User_id)
	fmt.Println("repo Comment_id:", comments.Comment_id)
	return comments, nil
}

func (c *CommentRepo) CommentRepoGet(ctx context.Context, id string) ([]*model.CommentGet, error) {
	sqlSt := `select
	c.c_id,
	c.c_message,
	c.photo_id,
	c.user_id,
	c.c_created_date,
	c.c_updated_date,
	u.u_id,
	u.u_email,
	u.u_username,
	p.p_id,
	p.p_title,
	p.p_caption,
	p.p_url,
	p.user_id
	from commentss c left join users u on c.user_id = u.u_id left join photo p on c.photo_id = p.p_id
	where u.u_id = $1
	group by c.c_idleft join photo p on c.photo_id = p.p_id;`
	rows, err := config.Db.Query(sqlSt, id)
	if err != nil {
		fmt.Println("Query row error")
	}
	fmt.Println("ini rows", rows)
	defer rows.Close()

	comment := []*model.CommentGet{}

	for rows.Next() {
		var comments model.CommentGet
		if err = rows.Scan(
			&comments.Comment_id,
			&comments.Message,
			&comments.Photo_id,
			&comments.User_id,
			&comments.Created_at,
			&comments.Updated_at,
			&comments.User.User_id,
			&comments.User.Email,
			&comments.User.Username,
			&comments.Photo.Photo_id,
			&comments.Photo.Title,
			&comments.Photo.Caption,
			&comments.Photo.User_id,
		); err != nil {
			fmt.Println("Scan row error")
		}
		comment = append(comment, &comments)
	}
	fmt.Println("ini rows", rows)
	fmt.Println("ini Comments", comment)

	return comment, nil
}

func (c *CommentRepo) CommentRepoUpdate(ctx context.Context, id string, comments *model.Comment) (*model.CommentShow, error) {
	sqlSt := `update commentss set c_message = $1, c_updated_date = $2
	where c_id = $3`
	res, err := config.Db.Exec(sqlSt,
		comments.Message,
		time.Now(),
		id,
	)
	if err != nil {
		fmt.Errorf("Error Update Photo: " + err.Error())
		return nil, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Errorf("Error Update Photo: " + err.Error())
		return nil, err
	}

	sc := model.CommentShow{}

	sqlSt2 := `select
	c.c_id,
	p.p_title,
	p.p_caption,
	p.p_caption,
	p.user_id,
	c.c_updated_date
	from commentss c join photo p on c.photo_id = p.p_id
	where c.c_id = $1
	group by c.c_id ,p.p_id ;
	`
	err2 := config.Db.QueryRow(sqlSt2, id).Scan(
		&sc.Comment_id,
		&sc.Title,
		&sc.Caption,
		&sc.Photo_url,
		&sc.User_id,
		&sc.Updated_at,
	)
	if err2 != nil {
		fmt.Errorf("Error Update Photo: " + err2.Error())
		return nil, err2
	}
	fmt.Println("updated data : ", count)
	return &sc, nil
}

func (c *CommentRepo) CommentRepoDelete(ctx context.Context, id string, comments *model.Comment) (*model.Comment, error) {
	sqlSt := `delete from commentss where c_id = $1`
	res, err := config.Db.Exec(sqlSt, id)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted Data : ", count)
	return comments, nil
}
