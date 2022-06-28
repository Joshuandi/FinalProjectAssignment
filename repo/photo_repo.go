package repo

import (
	"FinalProjectAssignment/config"
	"FinalProjectAssignment/model"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type PhotoRepoInterface interface {
	PhotoRepoRegister(ctx context.Context, photos *model.Photo) (*model.Photo, error)
	PhotoRepoUpdate(ctx context.Context, id string, photo *model.Photo) (*model.Photo, error)
	PhotoRepoDelete(ctx context.Context, id string, photos *model.Photo) (*model.Photo, error)
	PhotoRepoGet() ([]*model.PhotoGet, error)
}

type PhotoRepo struct {
	db *sql.DB
}

func NewPhotoRepo(db *sql.DB) PhotoRepoInterface {
	return &PhotoRepo{db: db}
}

func (p *PhotoRepo) PhotoRepoRegister(ctx context.Context, photos *model.Photo) (*model.Photo, error) {
	sqlSt := `insert into photo (p_title, p_caption, p_url, user_id, p_created_date, p_updated_date)
	values ($1, $2, $3, $4, $5, $5)
	returning p_id;`
	err := config.Db.QueryRow(sqlSt,
		photos.Title,
		photos.Caption,
		photos.Photo_url,
		photos.User_id,
		time.Now(),
	).Scan(&photos.Photo_id)
	fmt.Println("ini photo Repo : ", photos)
	if err != nil {
		fmt.Println("error scan:", err)
	}
	fmt.Println("repo photo:", photos)
	fmt.Println("repo Photo User_id:", photos.User_id)
	fmt.Println("repo Photo_id:", photos.Photo_id)
	return photos, nil
}

func (p *PhotoRepo) PhotoRepoGet() ([]*model.PhotoGet, error) {
	sqlSt := `select
	p.p_id,
	p.p_title,
	p.p_caption,
	p.p_url,
	p.user_id,
	p.p_created_date,
	p.p_updated_date,
	u.u_email,
	u.u_username
	from photo p join users u on p.user_id = u.u_id;`

	rows, err := config.Db.Query(sqlSt)
	if err != nil {
		fmt.Println("Query row error")
	}
	fmt.Println("ini rows", rows)
	defer rows.Close()

	photo := []*model.PhotoGet{}
	for rows.Next() {
		var photos model.PhotoGet
		if err = rows.Scan(
			&photos.Photo_id,
			&photos.Title,
			&photos.Caption,
			&photos.Photo_url,
			&photos.User_id,
			&photos.Created_at,
			&photos.Updated_at,
			&photos.User.Email,
			&photos.User.Username,
		); err != nil {
			fmt.Println(err)
		}
		photo = append(photo, &photos)
		fmt.Println("ini user email :", photos.User.Email)
		fmt.Println("ini user username :", photos.User.Username)
	}
	fmt.Println("ini rows", rows)
	fmt.Println("ini photos", photo)

	return photo, nil
}

func (p *PhotoRepo) PhotoRepoUpdate(ctx context.Context, id string, photo *model.Photo) (*model.Photo, error) {
	sqlSt := `update photo set p_title = $1, p_caption =$2, p_url =$3, p_updated_date = $4
	where P_id = $5`
	_, err := config.Db.Exec(sqlSt,
		photo.Title,
		photo.Caption,
		photo.Photo_url,
		time.Now(),
		id,
	)
	if err != nil {
		fmt.Errorf("Error Update Photo: " + err.Error())
		return nil, err
	}
	if err != nil {
		fmt.Errorf("Error Update Photo: " + err.Error())
		return nil, err
	}
	sqlSt2 := `select p_id, p_title, p_caption, p_url, user_id, p_updated_date from photo where p_id = $1`
	err2 := config.Db.QueryRow(sqlSt2, id).Scan(
		&photo.Photo_id,
		&photo.Title,
		&photo.Caption,
		&photo.Photo_url,
		&photo.User_id,
		&photo.Updated_at,
	)
	if err2 != nil {
		fmt.Errorf("Error Update Photo: " + err2.Error())
		return nil, err2
	}
	return photo, nil
}

func (p *PhotoRepo) PhotoRepoDelete(ctx context.Context, id string, photos *model.Photo) (*model.Photo, error) {
	sqlSt := `delete from photo where p_id = $1`
	_, err := config.Db.Exec(sqlSt, id)
	if err != nil {
		panic(err)
	}
	return photos, nil
}

//
