package authService

import "testing"

func TestValidHashAndValidatePassword(t *testing.T) {
	password := "somePassword123!@#$%^&*()_+"
	hashPassword := HashPassword(password)
	valid := ValidatePassword(password, hashPassword)
	if !valid {
		t.Fatalf("got %v expect %v", valid, true)
	}
}

func TestInvalidHashAndValidatePassword(t *testing.T) {
	password := "somePassword123!@#$%^&*()_+"
	hashPassword := HashPassword(password+" ")
	valid := ValidatePassword(password, hashPassword)
	if valid {
		t.Fatalf("got %v expect %v", valid, false)
	}
}
