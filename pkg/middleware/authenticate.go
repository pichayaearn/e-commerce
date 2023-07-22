package middleware

func authenticate(username, password string) (bool, error) {
	// Here, you would typically perform the authentication logic against your database or any other user store
	// For simplicity, we are using hardcoded values
	validUsername := "user123"
	validPassword := "password123"

	return username == validUsername && password == validPassword, nil
}
