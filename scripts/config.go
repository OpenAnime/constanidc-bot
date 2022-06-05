package scripts

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetBotToken() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env dosyası yuklenmedi hata :", err.Error())
	}
	return os.Getenv("TOKEN")
}
func GetMong() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env dosyası yuklenmedi hata :", err.Error())
	}
	return os.Getenv("MONGO_URL")
}
