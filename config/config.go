package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

//PORT has something to do
var (
	PORT      = 0
	DBURI     = ""
	RDSDRIVER = ""
	JWTKEY    = ""
)

//Load loads configurations
func Load() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		PORT = 5000
	}

	RDSDRIVER = os.Getenv("RDS_DRIVER")

	RDSHOSTNAME := os.Getenv("RDS_HOSTNAME")
	RDSUSERNAME := os.Getenv("RDS_USERNAME")
	RDSPASSWORD := os.Getenv("RDS_PASSWORD")
	RDSDATABASE := os.Getenv("RDS_DATABASE")
	RDSPORT := os.Getenv("RDS_PORT")

	//DBURI Format := id:password@tcp(your-amazonaws-uri.com:3306)/dbname
	DBURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		RDSUSERNAME,
		RDSPASSWORD,
		RDSHOSTNAME,
		RDSPORT,
		RDSDATABASE)

	// fmt.Println(DBURI)
	if err != nil {
		log.Fatal(err)
	}

}
