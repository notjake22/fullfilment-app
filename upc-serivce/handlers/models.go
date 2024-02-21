package handlers

type ItemResponse struct {
	Name     string `json:"name"`
	ImageURI string `json:"imageURI"`
	Success  bool   `json:"success"`
}
