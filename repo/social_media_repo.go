package repo

import (
	"FinalProjectAssignment/config"
	"FinalProjectAssignment/model"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type SocialMediaRepoInterface interface {
	SocialMediaRepoRegister(ctx context.Context, sm *model.SocialMedia) (*model.SocialMedia, error)
	SocialMediaRepoUpdate(ctx context.Context, id string, sm *model.SocialMedia) (*model.SocialMedia, error)
	SocialMediaRepoDelete(ctx context.Context, id string, sm *model.SocialMedia) (*model.SocialMedia, error)
	SocialMediaRepoGet(ctx context.Context, id string) ([]*model.SocialMediaShow, error)
}

type SocialMediaRepo struct {
	db *sql.DB
}

func NewSocialMediaRepo(db *sql.DB) SocialMediaRepoInterface {
	return &SocialMediaRepo{db: db}
}

func (s *SocialMediaRepo) SocialMediaRepoRegister(ctx context.Context, sm *model.SocialMedia) (*model.SocialMedia, error) {
	sqlSt := `insert into social_media (sm_name, sm_url, user_id, sm_created_date)
	values ($1, $2, $3, $4)
	returning sm_id;`
	err := config.Db.QueryRow(sqlSt,
		sm.Name,
		sm.SocialMedia_url,
		sm.User_id,
		time.Now(),
	).Scan(&sm.Sm_Id)
	fmt.Println("ini SocialMedia Repo : ", sm)
	if err != nil {
		fmt.Println("Query row error")
	}
	fmt.Println("repo SocialMedia:", sm)
	fmt.Println("repo SocialMedia User_id:", sm.User_id)
	fmt.Println("repo SocialMedia_id:", sm.Sm_Id)
	return sm, nil
}

func (s *SocialMediaRepo) SocialMediaRepoGet(ctx context.Context, id string) ([]*model.SocialMediaShow, error) {
	sqlSt := `
	select
	sm.sm_id, sm.sm_name, sm.sm_url,
	sm.user_id,
	sm.sm_created_date,
	sm.sm_updated_date,
	u.u_id, u.u_username
	from social_media sm left join users u on sm.user_id = u.u_id
	where sm.sm_id = $1
	group by sm.sm_id , u.u_id ;`
	rows, err := config.Db.Query(sqlSt, id)
	if err != nil {
		fmt.Println("Query row error")
	}
	fmt.Println("ini rows", rows)
	defer rows.Close()

	socmed := []*model.SocialMediaShow{}

	for rows.Next() {
		var sm model.SocialMediaShow
		if err = rows.Scan(
			&sm.Social_medias.Sm_Id,
			&sm.Social_medias.Name,
			&sm.Social_medias.SocialMedia_url,
			&sm.Social_medias.User_id,
			&sm.Social_medias.Created_at,
			&sm.Social_medias.Updated_at,
			&sm.Social_medias.User.User_id,
			&sm.Social_medias.User.Username,
		); err != nil {
			fmt.Println("Scan row error")
		}
		socmed = append(socmed, &sm)
	}
	fmt.Println("ini rows", rows)
	fmt.Println("ini socmed", socmed)
	return socmed, nil
}

func (s *SocialMediaRepo) SocialMediaRepoUpdate(ctx context.Context, id string, sm *model.SocialMedia) (*model.SocialMedia, error) {
	sqlSt := `update social_media set sm_name = $1, sm_url = $2, sm_updated_date = $3
	where sm_id = $4`
	res, err := config.Db.Exec(sqlSt,
		sm.Name,
		sm.SocialMedia_url,
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
	socmed := model.SocialMedia{}
	sqlSt2 := `select
	sm_id,
	sm_name,
	sm_url,
	user_id,
	sm_updated_date
	from social_media where sm_id = $1`
	err2 := config.Db.QueryRow(sqlSt2, id).Scan(
		&socmed.Sm_Id,
		&socmed.Name,
		&socmed.SocialMedia_url,
		&socmed.User_id,
		&socmed.Updated_at,
	)
	if err2 != nil {
		fmt.Errorf("Error Update Photo: " + err2.Error())
		return nil, err2
	}
	fmt.Println("updated data : ", count)
	return &socmed, nil
}

func (s *SocialMediaRepo) SocialMediaRepoDelete(ctx context.Context, id string, sm *model.SocialMedia) (*model.SocialMedia, error) {
	sqlSt := `delete from social_media where sm_id = $1`
	res, err := config.Db.Exec(sqlSt, id)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted Data : ", count)
	return sm, nil
}
