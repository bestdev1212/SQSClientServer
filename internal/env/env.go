package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load("./.env")
	if(err != nil){
		fmt.Println("Error to load env file");
		os.Exit(1)
	}
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}