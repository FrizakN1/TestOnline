package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"html/template"
	"knocker/database"
	"knocker/setting"
	"knocker/utils"
)

var privateJS template.JS

func main() {
	privateJS = template.JS(utils.LoadAssets("private/js/private.js"))

	settings := setting.LoadSetting("setting.json")
	database.Connect(settings)

	router := gin.Default()
	word := sessions.NewCookieStore([]byte("SecretX"))
	router.Use(sessions.Sessions("session", word))

	router.LoadHTMLGlob("template/*.html")
	router.Static("assets/", "assets/")

	router.GET("/", index)
	router.GET("/registration", reg)
	router.POST("/registration/checkEmail", regCheckEmail)
	router.POST("/registration/checkForm", regCheckForm)
	router.GET("/registration/success", success)
	router.GET("/login", login)
	router.POST("/login/check", loginCheck)
	router.GET("/profile", profile)
	router.POST("/exit", exit)
	router.GET("/create", create)
	router.POST("/create/save", saveTest)
	router.GET("/test/:id", testLoad)
	router.POST("/test/result", testResult)
	router.GET("/test_results", testResults)
	router.GET("/search_request", searchRequest)
	router.POST("/search_result", searchResult)
	router.GET("/registration/confirm/:hash", confirmUser)
	router.POST("/profile/changeName", changeName)
	router.POST("/profile/changeSurname", changeSurname)
	router.POST("/profile/changeEmail", changeEmailConfirm)
	router.GET("/profile/changeEmail/confirm/:hash", changeEmail)
	router.POST("/profile/changePassword", changePassword)
	router.POST("/search", search)
	_ = router.Run(settings.Address + ":" + settings.Port)
}

func search(c *gin.Context) {
	var value database.Search
	e := c.BindJSON(&value)
	if e != nil {
		utils.Logger.Println(e)
		return
	}
	var tests []database.TestList

	if value.Search == "" {
		tests = database.GetTestList()
	} else {
		tests = database.SearchRes(value.Search)
	}

	c.JSON(200, tests)
}

func changePassword(c *gin.Context) {
	session := getSession(c)
	var password database.Password
	e := c.BindJSON(&password)
	if e != nil {
		utils.Logger.Println(e)
		return
	}
	password.OldPassword, e = utils.Encrypt(password.OldPassword)
	if e != nil {
		utils.Logger.Println(e)
		return
	}
	res := session.CheckPassword(password.OldPassword)
	if res == "Пароль совпадает" {
		password.NewPassword, e = utils.Encrypt(password.NewPassword)
		if e != nil {
			utils.Logger.Println(e)
			return
		}
		c.JSON(200, session.ChangePassword(password.NewPassword))
	} else {
		c.JSON(200, res)
	}

}

func changeEmail(c *gin.Context) {
	hash := c.Param("hash")
	session := getSession(c)

	go session.ChangeEmail(hash)

	c.HTML(200, "successChangeEmail", gin.H{
		"Title": "Почта изменена",
		"Role":  session.User.Role,
		"Id":    session.User.Id,
		"Name":  session.User.Name,
	})
}

func changeEmailConfirm(c *gin.Context) {
	session := getSession(c)

	var newEmail database.User
	e := c.BindJSON(&newEmail)
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	c.JSON(200, session.ChangeEmailConfirm(newEmail.Email))
}

func changeSurname(c *gin.Context) {
	session := getSession(c)

	var newSurname database.User
	e := c.BindJSON(&newSurname)
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	c.JSON(200, session.ChangeSurname(newSurname.Surname))
}

func changeName(c *gin.Context) {
	session := getSession(c)

	var newName database.User
	e := c.BindJSON(&newName)
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	c.JSON(200, session.ChangeName(newName.Name))
}

func searchResult(c *gin.Context) {
	var searchData database.UserResult
	e := c.BindJSON(&searchData.Test)
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	searchResults := database.SearchResult(searchData.Test.TestName)
	c.JSON(200, searchResults)
}

func searchRequest(c *gin.Context) {
	session := getSession(c)
	if session.User.Role == "Admin" {
		c.HTML(200, "searchRequest", gin.H{
			"Title": "Поиск резултатов",
			"Role":  session.User.Role,
			"Id":    session.User.Id,
			"Name":  session.User.Name,
		})
	} else {
		c.Redirect(200, "/")
	}
}

func testResults(c *gin.Context) {
	session := getSession(c)

	var results []database.UserResult
	results = database.TestResults(session.User.Id)

	c.HTML(200, "results", gin.H{
		"Title":   "Результаты",
		"Results": results,
		"Role":    session.User.Role,
		"Id":      session.User.Id,
		"Name":    session.User.Name,
	})
}

func testResult(c *gin.Context) {
	session := getSession(c)
	form, e := c.MultipartForm()
	if e != nil {
		c.JSON(400, nil)
		return
	}

	test := database.Test{
		TestID:                     form.Value["TestID"][0],
		QuestionID:                 form.Value["QuestionID"],
		QuestionType:               form.Value["QuestionType"],
		DefBlockQuestion:           form.Value["QuestionID defBlock"],
		DefBlockAnswer:             form.Value["QuestionAnswer defBlock"],
		CheckBoxQuestion:           form.Value["QuestionID checkBox"],
		CheckBoxAnswer:             form.Value["QuestionAnswer checkBox"],
		CheckBoxAnswerCount:        form.Value["QuestionCount checkBox"],
		SelectBlockValue:           form.Value["QuestionValue selectBlock"],
		SelectBlockAnswer:          form.Value["QuestionAnswer selectBlock"],
		SelectBlockValeAnswerCount: form.Value["QuestionAnswerCount selectBlock"],
	}

	result := database.UserResult{
		Name:    session.User.Name,
		Surname: session.User.Surname,
		Result:  test.TestResult(session.User),
	}
	c.JSON(200, result)
}

func testLoad(c *gin.Context) {
	session := getSession(c)
	id := c.Param("id")

	var test []database.GetTest
	test = database.TestLoad(id)

	c.HTML(200, "test", gin.H{
		"Title":  "Тест",
		"Test":   test,
		"TestID": id,
		"Role":   session.User.Role,
		"Id":     session.User.Id,
		"Name":   session.User.Name,
	})
}

func saveTest(c *gin.Context) {

	form, e := c.MultipartForm()
	if e != nil {
		c.JSON(400, nil)
		return
	}

	test := database.Test{
		TestName:                   form.Value["Title"][0],
		DefBlockQuestion:           form.Value["defBlock Question"],
		DefBlockValue:              form.Value["defBlock Value"],
		DefBlockAnswer:             form.Value["defBlock Answer"],
		CheckBoxQuestion:           form.Value["checkBoxBlock Question"],
		CheckBoxValue:              form.Value["checkBoxBlock Value"],
		CheckBoxAnswer:             form.Value["checkBoxBlock Answer"],
		CheckBoxAnswerCount:        form.Value["checkBoxBlock Answer Count"],
		CheckBoxValueCount:         form.Value["checkBoxBlock Value Count"],
		SelectBlockQuestionCount:   form.Value["selectBlock Question Count"],
		SelectBlockValeAnswerCount: form.Value["selectBlock ValueAnswer Count"],
		SelectBlockValue:           form.Value["selectBlock Value"],
		SelectBlockAnswer:          form.Value["selectBlock Answer"],
	}

	c.JSON(200, test.SaveTest())
}

func getSession(c *gin.Context) *database.Session {
	_session := sessions.Default(c)

	sessionHash, ok := _session.Get("SessionSecretKey").(string)
	if ok {
		session := database.GetSession(sessionHash)
		if session != nil {
			session.Exists = true
			return session
		}
	}

	return &database.Session{
		Exists: false,
	}
}

func create(c *gin.Context) {
	session := getSession(c)

	if session.User.Role == "Admin" {
		c.HTML(200, "create", gin.H{
			"Title": "Создание теста",
			"Role":  session.User.Role,
			"Id":    session.User.Id,
			"Name":  session.User.Name,
			"Js":    privateJS,
		})
	} else {
		c.Redirect(301, "/")
	}
}

func profile(c *gin.Context) {
	session := getSession(c)
	c.HTML(200, "profile", gin.H{
		"Title":   "Профиль",
		"Role":    session.User.Role,
		"Id":      session.User.Id,
		"Name":    session.User.Name,
		"Surname": session.User.Surname,
		"State":   session.User.State,
		"Email":   session.User.Email,
	})
}

func exit(c *gin.Context) {
	session := sessions.Default(c)
	_session := getSession(c)

	_, ok := session.Get("SessionSecretKey").(string)
	if ok {
		session.Clear()
		_ = session.Save()
		c.SetCookie("hello", "", -1, "/", c.Request.URL.Hostname(), false, true)
		session.Delete("SessionSecretKey")
	}

	_session.DeleteSession()

	c.JSON(301, true)
}

func loginCheck(c *gin.Context) {
	session := sessions.Default(c)

	var user database.User
	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(400, nil)
		return
	}

	user.Password, e = utils.Encrypt(user.Password)
	if e != nil {
		c.JSON(400, nil)
		return
	}

	if user.LoginCheck() {
		hash, ok := database.CreateSession(&user)
		if ok {
			session.Set("SessionSecretKey", hash)
			e = session.Save()
			if e != nil {
				utils.Logger.Println(e)
			}

			c.JSON(200, true)

			return
		}
	}

	c.JSON(400, nil)
}

func login(c *gin.Context) {
	session := getSession(c)

	c.HTML(200, "login", gin.H{
		"Title": "Вход",
		"Role":  session.User.Role,
		"Id":    session.User.Id,
		"Name":  session.User.Name,
	})
}

func success(c *gin.Context) {
	session := getSession(c)

	c.HTML(200, "successConfirm", gin.H{
		"Title": "Успех",
		"Role":  session.User.Role,
		"Id":    session.User.Id,
		"Name":  session.User.Name,
	})
}

func confirmUser(c *gin.Context) {
	hash := c.Param("hash")
	session := getSession(c)

	go database.ConfirmUser(hash)

	c.HTML(200, "successConfirm", gin.H{
		"Title": "Подтверждение почты",
		"Role":  session.User.Role,
		"Id":    session.User.Id,
		"Name":  session.User.Name,
	})
}

func regCheckForm(c *gin.Context) {
	var user database.User

	e := c.BindJSON(&user)
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	user.Role = "Default"
	user.Password, e = utils.Encrypt(user.Password)
	if e != nil {
		utils.Logger.Println(e)
		c.JSON(400, nil)
		return
	}

	c.JSON(200, user.RegCheckForm())
}

func regCheckEmail(c *gin.Context) {
	var email database.UserEmail

	e := c.BindJSON(&email)
	if e != nil {
		utils.Logger.Println(e)
		c.JSON(400, nil)
		return
	}

	c.JSON(200, email.RegCheckEmil())
}

func reg(c *gin.Context) {
	session := getSession(c)

	c.HTML(200, "registration", gin.H{
		"Title": "Регистрация",
		"Role":  session.User.Role,
		"Id":    session.User.Id,
		"Name":  session.User.Name,
	})
}

func index(c *gin.Context) {
	session := getSession(c)

	tests := database.GetTestList()

	c.HTML(200, "index", gin.H{
		"Title": "Главная",
		"Role":  session.User.Role,
		"Id":    session.User.Id,
		"Name":  session.User.Name,
		"Tests": tests,
	})
}
