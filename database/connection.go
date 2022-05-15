package database

import (
	"database/sql"
	"fmt"
	"knocker/setting"
	"knocker/utils"

	_ "github.com/lib/pq"
)

var Link *sql.DB

/*"DbHost": "10.14.206.27",*/

func Connect(options *setting.Setting) {
	var e error
	Link, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		options.DbHost,
		options.DbPort,
		options.DbUser,
		options.DbPass,
		options.DbName))
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	e = Link.Ping()
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	errors := make([]string, 0)

	errors = append(errors, prepareRequest()...)
	errors = append(errors, prepareUser()...)

	if len(errors) > 0 {
		for _, i := range errors {
			utils.Logger.Println(i)
		}
	}

	LoadSession(sessionMap)
}
