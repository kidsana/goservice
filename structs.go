package structs
// Basic response struct
type BasicResponse struct {
    Error       int `json:"error"`
    Message     string `json:"message"`
}
