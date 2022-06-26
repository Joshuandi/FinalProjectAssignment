package main

import (
	"FinalProjectAssignment/config"
	"FinalProjectAssignment/handler"
	"FinalProjectAssignment/middleware"
	"FinalProjectAssignment/repo"
	"FinalProjectAssignment/router"
	"FinalProjectAssignment/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
)

var cfg config.Config

func main() {
	//database initilization
	_ = cleanenv.ReadConfig(".env", &cfg)
	config.Db, config.Err = sql.Open("postgres", ConnectDbPsql(
		cfg.Db_Host,
		cfg.Db_Dbname,
		cfg.Db_User,
		cfg.Db_Password,
		cfg.Db_Port,
	))
	defer config.Db.Close()
	if config.Err != nil {
		panic(config.Err)
	}
	config.Err = config.Db.Ping()
	if config.Err != nil {
		panic(config.Err)
	}
	fmt.Println("Successfully Connect to Database")

	//interface user
	userRepo := repo.NewUserRepo(config.Db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserhandler(userService)

	//interface photo
	photoRepo := repo.NewPhotoRepo(config.Db)
	photoService := service.NewPhotoService(photoRepo)
	photoHandler := handler.NewPhotohandler(photoService)

	//router
	r := mux.NewRouter()
	//middleware use
	r.Use(middleware.LoginMiddleware)
	//router connect
	router.UserRouter(r, userHandler)
	router.PhotoRouter(r, photoHandler)
	fmt.Println("Now Loading on Port", cfg.PORT)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8088",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func ConnectDbPsql(host, user, password, dbname string, port int) string {
	_ = cleanenv.ReadConfig(".env", &cfg)
	psqlInfo := fmt.Sprintf("host= %s port= %d user= %s "+
		" password= %s dbname= %s sslmode=disable",
		cfg.Db_Host,
		cfg.Db_Port,
		cfg.Db_User,
		cfg.Db_Password,
		cfg.Db_Dbname)
	return psqlInfo
}
