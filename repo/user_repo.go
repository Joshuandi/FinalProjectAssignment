package repo

import (
	"FinalProjectAssignment/config"
	"FinalProjectAssignment/model"
	"FinalProjectAssignment/util"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type UserRepoInterface interface {
	UserRepoRegister(ctx context.Context, users *model.User) (*model.User, error)
	UserRepoLogin(ctx context.Context, users *model.User) (*model.User, error)
	UserCheck(ctx context.Context, users *model.User) (*model.User, error)
	//UserRepoUpdate(ctx context.Context, users *model.User) (*model.User, error)
	//UserRepoGetId(ctx context.Context, id string) (*model.User, error)
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepoInterface {
	return &UserRepo{db: db}
}

func (u *UserRepo) UserRepoRegister(ctx context.Context, users *model.User) (*model.User, error) {
	pass, errHash := util.GenerateHashPassword(users.Password)
	if errHash != nil {
		fmt.Println("Error Hash : " + errHash.Error())
		return nil, errHash
	}
	users.Password = pass
	sqlSt := "insert into users (u_username, u_email, u_pass, u_age,u_created_date, u_updated_date)" +
		"values ($1, $2, $3, $4, $5, $5)" +
		"returning u_id;"

		//ctx nya tidak bisa
	rows, err := config.Db.Query(sqlSt,
		users.Username,
		users.Email,
		users.Password,
		users.Age,
		time.Now(),
	)
	if err != nil {
		fmt.Println("Query row error")
	}
	defer rows.Close()
	fmt.Println("repo user:", users)

	for rows.Next() {
		var id int
		if err1 := rows.Scan(&id); err1 != nil {
			fmt.Println("Scan Id Error")
			return nil, err1
		}
		users.User_id = id
	}
	fmt.Println("repo user_id:", users.User_id)

	return users, nil
}

func (u *UserRepo) UserRepoLogin(ctx context.Context, users *model.User) (*model.User, error) {
	sqlSt := "select u_email, u_pass from users"
	rows, err := config.Db.Query(sqlSt)
	if err != nil {
		panic(err)
	}
	fmt.Println("ini login repo:", users)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(
			&users.Email,
			&users.Password,
		); err != nil {
			fmt.Println("No Data", err)
		}
	}
	fmt.Println("ini login repo:", users)
	return users, nil
}

func (u *UserRepo) UserCheck(ctx context.Context, users *model.User) (*model.User, error) {
	sqlSt := "select u_email, u_username from users"
	rows, err := config.Db.Query(sqlSt)
	if err != nil {
		panic(err)
	}
	fmt.Println("ini email repo:", users)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&users.Email, &users.Username); err != nil {
			fmt.Println("No Data", err)
		}
	}
	fmt.Println("ini email repo:", users)
	return users, nil
}
