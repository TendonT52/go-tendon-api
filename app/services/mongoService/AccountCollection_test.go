package mongoService

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	// firstName := fmt.Sprintf("TestFirstName%dTest", time.Now().Unix())
	// lastName := fmt.Sprintf("TestLastName%dTest", time.Now().Unix())
	// email := fmt.Sprintf("Test%d@email.comTest", time.Now().Unix())
	// hashPassword := fmt.Sprintf("Test%dTest", time.Now().Unix())

	// _, err := AccountCollection.AddUser(firstName, lastName, email, hashPassword)
	// if err != nil {
	// 	t.Fatalf("got %v expect %v", err, nil)
	// }
}

// func TestGetUserById(t *testing.T) {
// 	firstName := fmt.Sprintf("TestFirstName%dTest", time.Now().Unix())
// 	lastName := fmt.Sprintf("TestLastName%dTest", time.Now().Unix())
// 	email := fmt.Sprintf("Test%d@email.comTest", time.Now().Unix())
// 	hashPassword := fmt.Sprintf("Test%dTest", time.Now().Unix())
// 	id, err := mongoService.AccountCollection.AddUser(firstName, lastName, email, hashPassword)

// 	if err != nil {
// 		t.Fatalf("got %v expect %v", err, nil)
// 	}

// 	user, err := mongoService.AccountCollection.GetUserById(id)
// 	if err != nil {
// 		t.Fatalf("got %v expect %v", err, nil)
// 	}
// 	if user.FirstName != firstName {
// 		t.Fatalf("got %v expect %v", user.FirstName, firstName)
// 	}
// }

// func TestGetUserByEmail(t *testing.T) {
// 	firstName := fmt.Sprintf("TestFirstName%dTest", time.Now().Unix())
// 	lastName := fmt.Sprintf("TestLastName%dTest", time.Now().Unix())
// 	email := fmt.Sprintf("Test%d@email.comTest", time.Now().Unix())
// 	hashPassword := fmt.Sprintf("Test%dTest", time.Now().Unix())
// 	_, err := mongoService.AccountCollection.AddUser(firstName, lastName, email, hashPassword)

// 	if err != nil {
// 		t.Fatalf("got %v expect %v", err, nil)
// 	}

// 	user, err := mongoService.AccountCollection.GetUserByEmail(email)
// 	if err != nil {
// 		t.Fatalf("got %v expect %v", err, nil)
// 	}
// 	if user.FirstName != firstName {
// 		t.Fatalf("got %v expect %v", user.FirstName, firstName)
// 	}
// }

// func TestDeleteUserById(t *testing.T) {
// 	firstName := fmt.Sprintf("TestFirstName%dTest", time.Now().Unix())
// 	lastName := fmt.Sprintf("TestLastName%dTest", time.Now().Unix())
// 	email := fmt.Sprintf("Test%d@email.comTest", time.Now().Unix())
// 	hashPassword := fmt.Sprintf("Test%dTest", time.Now().Unix())
// 	id, err := mongoService.AccountCollection.AddUser(firstName, lastName, email, hashPassword)
// 	if err != nil {
// 		t.Log("Error while AddUser")
// 		t.Fatalf("got %v expect %v", err, nil)
// 	}
// 	delCount, err := mongoService.AccountCollection.DeleteUserById(id)
// 	if err != nil {
// 		t.Fatalf("got %v expect %v", err, nil)
// 	} else if delCount != 1 {
// 		t.Fatalf("got %v expect %v", delCount, 1)
// 	}
// }
