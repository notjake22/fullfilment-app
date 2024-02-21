package load

const (
	baseApiUrl  string = "https://glum-backend-production.up.railway.app/b2"
	buildNumber int    = 2
)

type LoginResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
}

type BuildResponse struct {
	BuildNumber int `json:"build"`
}
