package service

import (
	"FinalProjectAssignment/model"
	"FinalProjectAssignment/repo"
	"FinalProjectAssignment/util"
	"context"
	"errors"
	"fmt"
)

type UserServiceInterface interface {
	UserRegister(ctx context.Context, users *model.User) (*model.User, error)
	UserLogin(ctx context.Context, login *model.UserPostLogin) (*model.User, error)
	UserGetId(ctx context.Context, user_id string) (*model.User, error)
	UserUpdate(ctx context.Context, id string, users *model.UserUpdateInput) (*model.User, error)
	UserDelete(ctx context.Context, id string, users *model.User) (*model.User, error)
}

type UserService struct {
	userRepo repo.UserRepoInterface
}

func NewUserService(userRepo repo.UserRepoInterface) UserServiceInterface {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) UserRegister(ctx context.Context, users *model.User) (*model.User, error) {
	email := users.Email
	username := users.Username
	//digunakan saat data awal sudah ada untuk check data
	// userCheck, err := u.userRepo.UserCheck(ctx, &model.User{
	// 	Email:    email,
	// 	Username: username,
	// })
	// fmt.Println("ini login srvc email masuk dari handler:", email)
	// fmt.Println("ini hasil database:", userCheck.Email)
	// if email == userCheck.Email {
	// 	return nil, errors.New("Email already registered")
	// }
	if _, ok := util.ValidateEmail(users.Email); !ok {
		return nil, errors.New("Email must valid")
	}
	if email == "" {
		return nil, errors.New("Email must be input")
	}
	if username == "" {
		return nil, errors.New("Username must be input")
	}
	// fmt.Println("ini login srvc email masuk dari handler:", username)
	// fmt.Println("ini hasil database:", userCheck.Username)

	//gunakan saat sudah ada user nya
	// if username == userCheck.Username {
	// 	return nil, errors.New("Username already registered")
	// }
	if users.Password == "" {
		return nil, errors.New("Password must be input")
	}
	if len(users.Password) < 6 {
		return nil, errors.New("Password must more than 6 character")
	}
	if users.Age == 0 {
		return nil, errors.New("age must be input")
	}
	if users.Age <= 8 {
		return nil, errors.New("age must more than 8 years")
	}
	userRegister, err := u.userRepo.UserRepoRegister(ctx, users)
	if err != nil {
		fmt.Println("Error While Register", err.Error())
		return nil, err
	}
	fmt.Println("service user:", userRegister)
	return userRegister, nil
}

func (u *UserService) UserLogin(ctx context.Context, login *model.UserPostLogin) (*model.User, error) {
	email := login.Email
	pass := login.Password
	//fmt.Println("ini login srvc email masuk dari handler:", email)
	//fmt.Println("ini login srvc password masuk dari handler:", pass)
	//data tidak ter process ke repo tapi di service muncul
	userLogin, err := u.userRepo.UserRepoLogin(ctx, &model.User{
		Email:    email,
		Password: pass,
	})
	//fmt.Println("ini login srvc :", userLogin)
	if err != nil {
		fmt.Println("Error While Login", err.Error())
		return nil, err
	}
	//fmt.Println("ini login email database check :", userLogin.Email)
	//fmt.Println("ini login email input check :", email)
	if email != userLogin.Email {
		return nil, errors.New("Email doesnt match")
	}
	if _, ok := util.ValidateEmail(email); !ok {
		return nil, errors.New("Email must valid")
	}
	if email == "" {
		return nil, errors.New("Email must be input")
	}
	if pass == "" {
		return nil, errors.New("Password must be input")
	}
	if len(pass) < 6 {
		return nil, errors.New("Password must more than 6 character")
	}
	if !util.CheckHashPassword(pass, userLogin.Password) {
		return nil, errors.New("Password doesnt match")
	}
	//fmt.Println("service user:", userLogin)
	return userLogin, nil
}

func (u *UserService) UserUpdate(ctx context.Context, id string, users *model.UserUpdateInput) (*model.User, error) {
	user_id, err := u.userRepo.UserRepoGetId(ctx, id)
	fmt.Println("UserId : ", user_id)
	if err != nil {
		return nil, err
	}
	if users.U_email == "" {
		return nil, errors.New(" Email must be input")
	}
	if users.U_username == "" {
		return nil, errors.New(" Username must be input")
	}
	user_id.Email = users.U_email
	user_id.Username = users.U_username
	user_id.Updated_at = users.U_Updated_at
	//fmt.Println("UserId Email: ", user_id.Email)
	//fmt.Println("UserId Username: ", user_id.Username)
	updateUser, err := u.userRepo.UserRepoUpdate(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return updateUser, nil
}

func (u *UserService) UserDelete(ctx context.Context, id string, users *model.User) (*model.User, error) {
	user_id, err := u.userRepo.UserRepoGetId(ctx, id)
	if err != nil {
		return nil, err
	}
	user_id.User_id = users.User_id

	deleteUser, err := u.userRepo.UserRepoDelete(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return deleteUser, nil
}

func (u *UserService) UserGetId(ctx context.Context, user_id string) (*model.User, error) {
	users_id, err := u.userRepo.UserRepoGetId(ctx, user_id)
	if err != nil {
		return nil, err
	}
	fmt.Println(users_id)
	return users_id, nil
}

//