package database

import (
	"database/sql"
	"knocker/utils"
	"net/smtp"
	"time"
)

type UserEmail struct {
	Email string
}

type UserResult struct {
	ID      int
	Name    string
	Surname string
	Test    Test
	Result  string
}

type User struct {
	Id       int
	Name     string
	Surname  string
	State    string
	Email    string
	Password string
	Role     string
}

type Session struct {
	Hash   string
	User   User
	Date   string
	Exists bool
}

type Password struct {
	OldPassword string
	NewPassword string
}

var query map[string]*sql.Stmt
var sessionMap map[string]Session

func prepareUser() []string {
	errors := make([]string, 0)
	if query == nil {
		query = make(map[string]*sql.Stmt)
	}
	sessionMap = make(map[string]Session)

	var e error

	query["checkEmail"], e = Link.Prepare(`SELECT email FROM "users" WHERE email = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["addUser"], e = Link.Prepare(`INSERT INTO "users" ("name", "surname", "state", "email", "password", "role") VALUES($1, $2, $3, $4, $5, $6)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["addUserConfirm"], e = Link.Prepare(`INSERT INTO "users_confirm" ("hash", "name", "surname", "state", "email", "password", "role") VALUES($1, $2, $3, $4, $5, $6, $7)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["DeleteUserConfirm"], e = Link.Prepare(`DELETE FROM "users_confirm" WHERE "hash" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetUserConfirmData"], e = Link.Prepare(`SELECT "name", "surname", "state", "email", "password", "role" FROM "users_confirm" WHERE "hash" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SessionSelect"], e = Link.Prepare(`SELECT "hash", "id", "name", "surname", "state", "email", "role", "date" FROM "sessions" AS s INNER JOIN "users" AS u ON u."id" = s."user"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SessionInsert"], e = Link.Prepare(`INSERT INTO "sessions" ("hash", "user", "date") VALUES ($1, $2, CURRENT_TIMESTAMP)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SessionDelete"], e = Link.Prepare(`DELETE FROM "sessions" WHERE "hash" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["Login"], e = Link.Prepare(`SELECT "id", "name", "surname", "state", "role" FROM "users" WHERE "email" = $1 AND "password" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetTestsResults"], e = Link.Prepare(`SELECT "test_id", "result" FROM "result" WHERE "user_id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetTestsName"], e = Link.Prepare(`SELECT "name" FROM "title" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetTestID"], e = Link.Prepare(`SELECT "id" FROM "title" WHERE "name" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SearchTestsResults"], e = Link.Prepare(`SELECT "user_id", "result" FROM "result" WHERE "test_id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetUserNameSurname"], e = Link.Prepare(`SELECT "name", "surname" FROM "users" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["ChangeName"], e = Link.Prepare(`UPDATE "users" SET "name" = $1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["ChangeSurname"], e = Link.Prepare(`UPDATE "users" SET "surname" = $1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["addChangeEmailConfirm"], e = Link.Prepare(`INSERT INTO "change_email" ("hash", "old_email", "new_email") VALUES ($1, $2, $3)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetChangeEmailConfirmData"], e = Link.Prepare(`SELECT "new_email" FROM "change_email" WHERE "hash" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["DeleteChangeEmailConfirmData"], e = Link.Prepare(`DELETE FROM "change_email" WHERE "hash" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["ChangeEmail"], e = Link.Prepare(`UPDATE "users" SET "email" = $1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetPassword"], e = Link.Prepare(`SELECT "password" FROM "users" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["ChangePassword"], e = Link.Prepare(`UPDATE "users" SET "password" = $1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}
	return errors
}

func (s *Session) ChangePassword(newPassword string) string {
	stmt, ok := query["ChangePassword"]
	if !ok {
		return ""
	}

	_, e := stmt.Exec(newPassword, s.User.Id)
	if e != nil {
		utils.Logger.Println(e)
		return ""
	}

	return "Пароль изменен"
}

func (s *Session) CheckPassword(oldPassword string) string {
	stmt, ok := query["GetPassword"]
	if !ok {
		return ""
	}
	var currentPassword string
	row := stmt.QueryRow(s.User.Id)
	e := row.Scan(&currentPassword)
	if e != nil {
		utils.Logger.Println(e)
		return ""
	}
	if currentPassword == oldPassword {
		return "Пароль совпадает"
	} else {
		return "Пароль не совпадает"
	}
}

func (s *Session) ChangeEmail(hash string) bool {
	var newEmail string
	stmt, ok := query["GetChangeEmailConfirmData"]
	if !ok {
		return false
	}

	row := stmt.QueryRow(hash)
	e := row.Scan(&newEmail)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	stmt, ok = query["DeleteChangeEmailConfirmData"]
	if !ok {
		return false
	}

	_, e = stmt.Exec(hash)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	stmt, ok = query["ChangeEmail"]
	if !ok {
		return false
	}

	_, e = stmt.Exec(newEmail, s.User.Id)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	sessionMap[s.Hash] = Session{
		Hash: s.Hash,
		User: User{
			Id:       s.User.Id,
			Name:     s.User.Name,
			Surname:  s.User.Surname,
			State:    s.User.State,
			Email:    newEmail,
			Password: "",
			Role:     s.User.Role,
		},
		Date: time.Now().String()[:19],
	}

	return true
}

func (s *Session) ChangeEmailConfirm(newEmail string) bool {
	hash, e := utils.GenerateHash(s.User.Email)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	stmt, ok := query["addChangeEmailConfirm"]
	if !ok {
		return false
	}

	_, e = stmt.Exec(hash, s.User.Email, newEmail)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	auth := smtp.PlainAuth("", "testsonlinesender@yandex.ru", "Root1_Root2", "smtp.yandex.ru")
	to := []string{s.User.Email}
	msg := []byte("From: testsonlinesender@yandex.ru\r\n" +
		"To: " + s.User.Email + "\r\n" +
		"Subject: Изменение почты\r\n\r\n" +
		"Для изминения почты  на " + newEmail + " перейдите по ссылке localhost:8080/profile/changeEmail/confirm/" + hash + "\r\n")
	e = smtp.SendMail("smtp.yandex.ru:587", auth, "testsonlinesender@yandex.ru", to, msg)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	return true
}

func (s *Session) ChangeSurname(newSurname string) bool {
	stmt, ok := query["ChangeSurname"]
	if !ok {
		return false
	}

	_, e := stmt.Exec(newSurname, s.User.Id)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	sessionMap[s.Hash] = Session{
		Hash: s.Hash,
		User: User{
			Id:       s.User.Id,
			Name:     s.User.Name,
			Surname:  newSurname,
			State:    s.User.State,
			Email:    s.User.Email,
			Password: "",
			Role:     s.User.Role,
		},
		Date: time.Now().String()[:19],
	}

	return true
}

func (s *Session) ChangeName(newName string) bool {
	stmt, ok := query["ChangeName"]
	if !ok {
		return false
	}

	_, e := stmt.Exec(newName, s.User.Id)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	sessionMap[s.Hash] = Session{
		Hash: s.Hash,
		User: User{
			Id:       s.User.Id,
			Name:     newName,
			Surname:  s.User.Surname,
			State:    s.User.State,
			Email:    s.User.Email,
			Password: "",
			Role:     s.User.Role,
		},
		Date: time.Now().String()[:19],
	}

	return true
}

func SearchResult(testName string) []UserResult {
	var results []UserResult
	stmt, ok := query["GetTestID"]
	if !ok {
		return nil
	}

	row := stmt.QueryRow(testName)
	var testID int
	e := row.Scan(&testID)
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}

	stmt, ok = query["SearchTestsResults"]
	if !ok {
		return nil
	}

	rows, e := stmt.Query(testID)
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var result UserResult
		e = rows.Scan(&result.ID, &result.Result)
		results = append(results, result)
	}

	stmt, ok = query["GetUserNameSurname"]
	if !ok {
		return nil
	}

	for i := 0; i < len(results); i++ {
		row = stmt.QueryRow(results[i].ID)
		e = row.Scan(&results[i].Name, &results[i].Surname)
		if e != nil {
			utils.Logger.Println(e)
			return nil
		}
	}
	return results
}

func TestResults(userID int) []UserResult {
	var results []UserResult

	stmt, ok := query["GetTestsResults"]
	if !ok {
		return nil
	}

	rows, e := stmt.Query(userID)
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var result UserResult
		e = rows.Scan(&result.Test.TestID, &result.Result)
		if e != nil {
			utils.Logger.Println(e)
			return nil
		}
		results = append(results, result)
	}

	stmt, ok = query["GetTestsName"]
	if !ok {
		return nil
	}

	for i := 0; i < len(results); i++ {
		row := stmt.QueryRow(results[i].Test.TestID)
		e = row.Scan(&results[i].Test.TestName)
		if e != nil {
			utils.Logger.Println(e)
			return nil
		}
	}

	return results
}

func (user *User) LoginCheck() bool {
	stmt, ok := query["Login"]
	if !ok {
		return false
	}
	row := stmt.QueryRow(user.Email, user.Password)
	e := row.Scan(&user.Id, &user.Name, &user.Surname, &user.State, &user.Role)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}
	return true
}

func GetSession(hash string) *Session {
	session, ok := sessionMap[hash]
	if ok {
		return &session
	}

	return nil
}

func (s *Session) DeleteSession() {
	stmt, ok := query["SessionDelete"]
	if !ok {
		return
	}

	_, e := stmt.Exec(s.Hash)
	if e != nil {
		utils.Logger.Println(e)
	}

	return
}

func CreateSession(user *User) (string, bool) {
	stmt, ok := query["SessionInsert"]
	if !ok {
		return "", false
	}

	hash, e := utils.GenerateHash(user.Email)
	if e != nil {
		utils.Logger.Println(e)
		return "", false
	}

	_, e = stmt.Exec(hash, user.Id)
	if e != nil {
		utils.Logger.Println(e)
		return "", false
	}

	if sessionMap != nil {
		sessionMap[hash] = Session{
			Hash: hash,
			User: User{
				Id:       user.Id,
				Name:     user.Name,
				Surname:  user.Surname,
				State:    user.State,
				Email:    user.Email,
				Password: "",
				Role:     user.Role,
			},
			Date: time.Now().String()[:19],
		}
	}

	return hash, true
}

func LoadSession(m map[string]Session) {
	stmt, ok := query["SessionSelect"]
	if !ok {
		return
	}

	rows, e := stmt.Query()
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var session Session
		e = rows.Scan(&session.Hash, &session.User.Id, &session.User.Name, &session.User.Surname, &session.User.State, &session.User.Email, &session.User.Role, &session.Date)
		if e != nil {
			utils.Logger.Println(e)
			return
		}

		m[session.Hash] = session
	}
}

func ConfirmUser(hash string) {
	var user User

	stmt, ok := query["GetUserConfirmData"]
	if !ok {
		return
	}

	row := stmt.QueryRow(hash)
	e := row.Scan(&user.Name, &user.Surname, &user.State, &user.Email, &user.Password, &user.Role)
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	stmt, ok = query["addUser"]
	if !ok {
		return
	}

	_, e = stmt.Exec(user.Name, user.Surname, user.State, user.Email, user.Password, user.Role)
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	stmt, ok = query["DeleteUserConfirm"]
	if !ok {
		return
	}

	_, e = stmt.Exec(hash)
	if e != nil {
		utils.Logger.Println(e)
		return
	}
}

func (user *User) RegCheckForm() bool {
	hash, e := utils.GenerateHash(user.Email)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	stmt, ok := query["addUserConfirm"]
	if !ok {
		return false
	}

	_, e = stmt.Exec(hash, user.Name, user.Surname, user.State, user.Email, user.Password, user.Role)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	auth := smtp.PlainAuth("", "testsonlinesender@yandex.ru", "Root1_Root2", "smtp.yandex.ru")
	to := []string{user.Email}
	msg := []byte("From: testsonlinesender@yandex.ru\r\n" +
		"To: " + user.Email + "\r\n" +
		"Subject: Подтверждение почты\r\n\r\n" +
		"Для подтверждения почты перейдите по ссылке localhost:8080/registration/confirm/" + hash + "\r\n")
	e = smtp.SendMail("smtp.yandex.ru:587", auth, "testsonlinesender@yandex.ru", to, msg)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	return true
}

func (email *UserEmail) RegCheckEmil() bool {
	stmt, ok := query["checkEmail"]
	if !ok {
		return false
	}

	row := stmt.QueryRow(email.Email)

	e := row.Scan(&email.Email)
	if e != nil {
		return true
	}

	return false
}
