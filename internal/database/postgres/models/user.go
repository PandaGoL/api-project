package models

import "strings"

type User struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

var userColumnList = []string{
	"user_id", "first_name",
	"last_name", "email",
	"phone",
}

var userEngine = NewModelEngine(userColumnList)

func (u *User) GetFields() []interface{} {
	return []interface{}{
		&u.UserID, &u.FirstName,
		&u.LastName, &u.Email,
		&u.Phone,
	}
}

func (u *User) GetFieldValue() []interface{} {
	return []interface{}{
		u.UserID, u.FirstName,
		u.LastName, u.Email,
		u.Phone,
	}
}

func (c *User) getEngine() *ModelEngine {
	return userEngine
}

func (u *User) GetColumns() string {
	return getColumnList(u.getEngine().Columns)
}

func (u *User) GetColumnListWhithTableAbbreviation(abbr string) string {
	return getColumnListWhithTableAbbreviation(abbr, u.getEngine().Columns)
}

func (u *User) GetPlaceholders() (result string) {
	return strings.Join(u.getEngine().Placeholders, ", ")
}

func (u *User) GetPlaceholder(column string) string {
	result, ok := u.getEngine().PlaceholderMap[column]
	if ok {
		return result
	}
	return "$ERROR$"
}

func (u *User) ToMap() map[string]interface{} {
	fields := make(map[string]interface{})
	fields["id"] = u.UserID
	fields["first_name"] = u.FirstName
	fields["last_name"] = u.LastName
	fields["email"] = u.Email
	fields["phone"] = u.Phone

	return fields
}
