package repository

import (
	"fmt"
	"testing"
)

const test = "[TEST] "

func TestPostAuthUser(t *testing.T) {
	db := New()
	originalAmount := len(db.UsersCollection)
	db.PostAuthUser(User{})

	if len(db.UsersCollection) == originalAmount {
		t.Errorf(test + "User was not added")
	}
}

func TestGetAuthUsers(t *testing.T) {
	db := New()
	var auths = GetAuthorized(db)

	if len(auths) <= 0 {
		t.Errorf(test + "Couldn't get any authorized user")
	}
}

func TestSendMessage(t *testing.T) {
	// assuming there's more than 1 user, and a ton of other things.

	db := New()
	originalAmount := len(db.UsersCollection[0].UnseenMessages)
	var user = &db.UsersCollection[1]
	message := Message{db.UsersCollection[0].Name, user.Name, "TestMessage"}

	ok := db.SendMessage(user, message)

	if len(db.UsersCollection[0].UnseenMessages) == originalAmount {
		t.Errorf(test + "Couldn't send the message!")
		fmt.Println("SendMessage returned: /v", ok)
	}
}

func TestGetSeenMessages(t *testing.T) {
	db := New()
	messages := db.GetSeenMessages(&db.UsersCollection[0])

	if len(messages) <= 0 {
		t.Errorf(test + "Couldn't retrieve the messages from the user!")
	}
}

func TestGetUnseenMessages(t *testing.T) {
	db := New()
	messages := db.GetUnseenMessages(&db.UsersCollection[0])

	if len(messages) <= 0 {
		t.Errorf(test + "Couldn't retrieve the unseen messages from the user!")
	}
}

func ValidateAuthorizedUser(t *testing.T) {
	db := New()
	userToValidate := User{Name: "TestUser", Password: "pass"} // we could retrieve the user from db... but it would later validate against exactly the same thing. would it differ?
	if ok := db.ValidateAuthorizedUser(userToValidate); !ok {
		t.Errorf(test + "Failed to validate user password")
	}
}
