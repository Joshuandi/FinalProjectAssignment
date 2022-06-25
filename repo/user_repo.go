package repo

import (
	"FinalProjectAssignment/config"
	"FinalProjectAssignment/model"
	"FinalProjectAssignment/util"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type UserRepoInterface interface {
	UserRepoRegister(ctx context.Context, users *model.User) (*model.User, error)
	UserRepoLogin(ctx context.Context, users *model.User) (*model.User, error)
	UserCheck(ctx context.Context, users *model.User) (*model.User, error)
	UserRepoUpdate(ctx context.Context, users *model.User) (*model.User, error)
	UserRepoDelete(ctx context.Context, users *model.User) (*model.User, error)
	UserRepoGetId(ctx context.Context, user_id string) (*model.User, error)
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
	sqlSt := `insert into users (u_username, u_email, u_pass, u_age,u_created_date, u_updated_date)
		values ($1, $2, $3, $4, $5, $5)
		returning u_id;`
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
	sqlSt := "select u_id, u_email, u_pass from users where u_email = $1"
	rows, err := config.Db.Query(sqlSt, users.Email)
	if err != nil {
		panic(err)
	}
	fmt.Println("ini login repo:", users)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(
			&users.User_id,
			&users.Email,
			&users.Password,
		); err != nil {
			fmt.Println("No Data", err)
			errors.New("No Data")
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

func (u *UserRepo) UserRepoUpdate(ctx context.Context, users *model.User) (*model.User, error) {
	sqlSt := `update users set u_email = $1, u_username = $2, u_updated_date = $3 where u_id = $4`
	res, err := config.Db.Exec(sqlSt,
		&users.Email,
		&users.Username,
		time.Now(),
		&users.User_id,
	)
	if err != nil {
		fmt.Errorf("Error Update User: " + err.Error())
		return nil, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Errorf("Error Update User: " + err.Error())
		return nil, err
	}
	fmt.Println("updated data : ", count)
	return users, nil
}

func (u *UserRepo) UserRepoDelete(ctx context.Context, users *model.User) (*model.User, error) {
	sqlSt := `delete from users where u_id = $1`
	res, err := config.Db.Exec(sqlSt, users.User_id)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted Data : ", count)
	return users, nil
}

func (u *UserRepo) UserRepoGetId(ctx context.Context, user_id string) (*model.User, error) {
	var users model.User
	sqlSt := `Select u_id, u_username, u_email, u_age, u_updated_date from users where u_id = $1;`
	row, err := config.Db.Query(sqlSt, user_id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if err = row.Scan(
			&users.User_id,
			&users.Username,
			&users.Email,
			&users.Age,
			&users.Updated_at,
		); err != nil {
			fmt.Println("No Data", err)
		}
	}
	return &users, nil
}
