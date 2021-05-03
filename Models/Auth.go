package Models

type User struct {
	Api_key    string `json:"api_key"`
	Hash_value string `json:"hash_value"`
}

type Email struct {
	Email string `json:"email"`
}
