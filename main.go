package quickscan_backend

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rithikjain/quickscan-backend/api/handler"
	"github.com/rithikjain/quickscan-backend/pkg/entities"
	"github.com/rithikjain/quickscan-backend/pkg/user"
	"log"
	"net/http"
	"os"
)

func dbConnect(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {
	// In the case of heroku
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	}
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode),
	)

	return db, err
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("INFO: No PORT environment variable detected, defaulting to 4000")
		return "localhost:4000"
	}
	return ":" + port
}

func main() {
	if os.Getenv("onServer") != "True" {
		// Loading the .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Setting up DB
	db, err := dbConnect(
		os.Getenv("dbHost"),
		os.Getenv("dbPort"),
		os.Getenv("dbUser"),
		os.Getenv("dbName"),
		os.Getenv("dbPass"),
		os.Getenv("sslmode"),
	)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}

	// Creating the tables
	db.AutoMigrate(&entities.User{})

	defer db.Close()
	fmt.Println("Connected to DB...")

	// Setting up the router
	r := http.NewServeMux()

	// Users
	userRepo := user.NewRepo(db)
	userSvc := user.NewService(userRepo)
	handler.MakeUserHandler(r, userSvc)

	// To check if server up or not
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello There"))
		return
	})

	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(GetPort(), r))
}
