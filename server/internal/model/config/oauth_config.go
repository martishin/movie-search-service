package config

type OAuthConfig struct {
	ClientID      string
	ClientSecret  string
	CallbackURL   string
	RedirectURL   string
	SessionSecret string
	Domain        string
	IsProduction  bool
}
