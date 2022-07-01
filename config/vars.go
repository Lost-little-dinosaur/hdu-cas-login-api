package config

type GlobalConfig struct {
	MODE          string
	ProgramName   string
	AUTHOR        string
	VERSION       string
	FrontendLogin string
	Auth          struct {
		Secret string
		Issuer string
	}

	OAUTH struct {
		GITHUB struct {
			ClientId     string
			ClientSecret string
			RedirectUri  string
			Scope        string
		}
	}
}
