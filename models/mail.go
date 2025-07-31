package models

type EmailData struct {
	To        string
	Subject   string
	HTMLBody  string
	From      string
	ImagePath string
	SMTConfig SMTConfig
}

type SMTConfig struct {
	SMTPHost string
	SMTPPort int
	SMTPUser string
	SMTPPass string
}
