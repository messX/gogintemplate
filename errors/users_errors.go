package errors

import "fmt"

type UserTakenError struct {
	Username string
}

func (m *UserTakenError) Error() string {
	return fmt.Sprintf("Username %s already taken", m.Username)
}

type UserNameInvalid struct {
	Username string
}

func (m *UserNameInvalid) Error() string {
	return fmt.Sprintf("Username %s is invalid", m.Username)
}

type UserPasswordInvalid struct {
	Username string
}

func (m *UserPasswordInvalid) Error() string {
	return fmt.Sprintf("Username %s provided invalid password", m.Username)
}
