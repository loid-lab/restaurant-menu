package initializers

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/loid-lab/restaurant-menu/models"
)

var Env models.SMTConfig

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load Environment variable: %s", err)
	}

	smtPortStr := os.Getenv("SMTP_PORT")
	smtpPort, err := strconv.Atoi(smtPortStr)
	if err != nil {
		log.Fatalf("Invalid SMT_PORT: must be integer. Got %s", smtPortStr)
	}

	Env = models.SMTConfig{
		SMTPHost: os.Getenv("SMTP_Host"),
		SMTPPort: smtpPort,
		SMTPUser: os.Getenv("SMTP_USER"),
		SMTPPass: os.Getenv("SMTP_PASS"),
	}
}
