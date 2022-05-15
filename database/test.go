package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"knocker/utils"
	"strconv"
)

type Test struct {
	TestID                     string
	TestName                   string
	QuestionID                 []string
	QuestionType               []string
	DefBlockQuestion           []string
	DefBlockValue              []string
	DefBlockAnswer             []string
	CheckBoxQuestion           []string
	CheckBoxValue              []string
	CheckBoxAnswer             []string
	CheckBoxAnswerCount        []string
	CheckBoxValueCount         []string
	SelectBlockQuestionCount   []string
	SelectBlockValeAnswerCount []string
	SelectBlockValue           []string
	SelectBlockAnswer          []string
}

type TrueAnswer struct {
	QuestionID   string
	QuestionType string
	Answer       []string
}

type GetTest struct {
	QuestionID   int
	Question     string
	QuestionType string
	Value        []string
	Answer       []string
}

type TestList struct {
	Id   int
	Name string
}

type Search struct {
	Search string
}

func prepareRequest() []string {
	errors := make([]string, 0)
	if query == nil {
		query = make(map[string]*sql.Stmt)
	}

	var e error

	query["AddTestName"], e = Link.Prepare(`INSERT INTO "title" ("name") VALUES ($1)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddDefBlockQuestion"], e = Link.Prepare(`INSERT INTO "question" ("test_id", "quest_id", "quest_name", "type") VALUES ($1, $2, $3, 'defBlock')`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddDefBlockValue"], e = Link.Prepare(`INSERT INTO "value_defBlock" ("test_id", "quest_id", "value_id", "value") VALUES($1, $2, $3, $4)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddDefBlockAnswer"], e = Link.Prepare(`INSERT INTO "answer_defBlock" ("test_id", "quest_id", "answer_id", "answer") VALUES ($1, $2, $3, $4)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddCheckBoxQuestion"], e = Link.Prepare(`INSERT INTO "question" ("test_id", "quest_id", "quest_name", "type") VALUES ($1, $2, $3, 'checkBox')`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddSelectBlockQuestion"], e = Link.Prepare(`INSERT INTO "question" ("test_id", "quest_id", "quest_name", "type") VALUES ($1, $2, $3, 'selectBlock')`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddSelectBlockValue"], e = Link.Prepare(`INSERT INTO "value_selectBlock" ("test_id", "quest_id", "value_id", "value") VALUES($1, $2, $3, $4)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddSelectBlockAnswer"], e = Link.Prepare(`INSERT INTO "answer_selectBlock" ("test_id", "quest_id", "answer_id", "answer") VALUES ($1, $2, $3, $4)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddCheckBoxValue"], e = Link.Prepare(`INSERT INTO "value_checkbox" ("test_id", "quest_id", "value_id", "value") VALUES($1, $2, $3, $4)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddCheckBoxAnswer"], e = Link.Prepare(`INSERT INTO "answer_checkbox" ("test_id", "quest_id", "answer_id", "answer") VALUES ($1, $2, $3, $4)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SelectTestId"], e = Link.Prepare(`SELECT "id" FROM "title" WHERE "name" = $1 ORDER BY "id"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetTestList"], e = Link.Prepare(`SELECT "id", "name" FROM "title" ORDER BY "id" DESC `)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetTestQuestion"], e = Link.Prepare(`SELECT "quest_id", "quest_name", "type" FROM "question" WHERE "test_id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetDefBlockValue"], e = Link.Prepare(`SELECT "value" FROM "value_defBlock" WHERE "test_id" = $1 AND "quest_id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetDefBlockAnswer"], e = Link.Prepare(`SELECT "answer" FROM "answer_defBlock" WHERE "test_id" = $1 AND "quest_id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetCheckBoxValue"], e = Link.Prepare(`SELECT "value" FROM "value_checkbox" WHERE "test_id" = $1 AND "quest_id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetSelectBlockValue"], e = Link.Prepare(`SELECT "value" FROM "value_selectBlock" WHERE "test_id" = $1 AND "quest_id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetCheckBoxAnswer"], e = Link.Prepare(`SELECT "answer" FROM "answer_checkbox" WHERE "test_id" = $1 AND "quest_id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetSelectBlockAnswer"], e = Link.Prepare(`SELECT "answer" FROM "answer_selectBlock" WHERE "test_id" = $1 AND "quest_id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetQuestionID"], e = Link.Prepare(`SELECT "quest_id", "type" FROM "question" WHERE "test_id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAnswerDefBlock"], e = Link.Prepare(`SELECT "answer" FROM "answer_defBlock" WHERE "test_id" = $1 AND "quest_id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAnswerCheckBox"], e = Link.Prepare(`SELECT "answer" FROM "answer_checkbox" WHERE "test_id" = $1 AND "quest_id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAnswerSelectBlock"], e = Link.Prepare(`SELECT "answer" FROM "answer_selectBlock" WHERE "test_id" = $1 AND "quest_id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddResult"], e = Link.Prepare(`INSERT INTO "result"("user_id", "test_id", "result") VALUES ($1, $2, $3)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetTestSearch"], e = Link.Prepare(`SELECT "id", "name" FROM "title" WHERE "name" ILIKE '%' || $1 || '%' ORDER BY "id" DESC `)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func SearchRes(value string) []TestList {
	var tests []TestList
	stmt, ok := query["GetTestSearch"]
	if !ok {
		return nil
	}

	rows, e := stmt.Query(value)
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}

	for rows.Next() {
		var test TestList
		e = rows.Scan(&test.Id, &test.Name)
		if e != nil {
			utils.Logger.Println(e)
			return nil
		}
		tests = append(tests, test)
	}
	return tests
}

func (test *Test) TestResult(UserID User) string {
	var trueAnswer []TrueAnswer
	var questionsID []TrueAnswer

	stmt, ok := query["GetQuestionID"]
	if !ok {
		return "Ошибка"
	}

	rows, e := stmt.Query(test.TestID)
	if e != nil {
		utils.Logger.Println(e)
		return "Ошибка"
	}

	defer rows.Close()

	questionCount := 0
	for rows.Next() {
		var questionTemp TrueAnswer
		e = rows.Scan(&questionTemp.QuestionID, &questionTemp.QuestionType)
		if e != nil {
			utils.Logger.Println(e)
			return "Ошибка"
		}
		questionsID = append(questionsID, questionTemp)
		questionCount++
	}

	var questionData TrueAnswer
	for _, id := range questionsID {
		questionData.QuestionID = id.QuestionID
		if id.QuestionType == "defBlock" {
			stmt, ok = query["GetAnswerDefBlock"]
			if !ok {
				return "Ошибка"
			}
		}
		if id.QuestionType == "checkBox" {
			stmt, ok = query["GetAnswerCheckBox"]
			if !ok {
				return "Ошибка"
			}
		}
		if id.QuestionType == "selectBlock" {
			stmt, ok = query["GetAnswerSelectBlock"]
			if !ok {
				return "Ошибка"
			}
		}

		rows, e = stmt.Query(test.TestID, id.QuestionID)
		if e != nil {
			utils.Logger.Println(e)
			return "Ошибка"
		}

		defer rows.Close()

		questionData.Answer = nil

		for rows.Next() {
			var temp string
			e = rows.Scan(&temp)
			if e != nil {
				utils.Logger.Println(e)
				return "Ошибка"
			}
			questionData.Answer = append(questionData.Answer, temp)
		}
		trueAnswer = append(trueAnswer, questionData)
	}

	trueAnswerCount := 0
	answerCountCheckBoxCounter := 0
	answerCountSelectBlockCounter := 0
	x := 0
	z := 0
	for i, id := range test.QuestionID {
		for _, el := range trueAnswer {
			if id == el.QuestionID {
				if test.QuestionType[i] == "defBlock" {
					if test.DefBlockAnswer[i] == el.Answer[0] {
						trueAnswerCount++
					}
				}
				if test.QuestionType[i] == "checkBox" {
					count := 0
					var countAnswerCheckBox int
					countAnswerCheckBox, e = strconv.Atoi(test.CheckBoxAnswerCount[answerCountCheckBoxCounter])
					if e != nil {
						utils.Logger.Println(e)
						return "Ошибка"
					}
					countAnswerCheckBox += x
					for j := x; j < countAnswerCheckBox; j++ {
						for _, item := range el.Answer {
							if test.CheckBoxAnswer[j] == item {
								count++
							}
						}
						x++
					}
					if countAnswerCheckBox == count {
						trueAnswerCount++
					}
					answerCountCheckBoxCounter++
				}
				if test.QuestionType[i] == "selectBlock" {
					countSelect := 0
					var countAnswerSelectBlock int
					countAnswerSelectBlock, e = strconv.Atoi(test.SelectBlockValeAnswerCount[answerCountSelectBlockCounter])
					if e != nil {
						utils.Logger.Println(e)
						return "Ошибка"
					}
					for _, item := range el.Answer {
						if test.SelectBlockAnswer[z] == item {
							countSelect++
						}
						z++
					}
					if countSelect == countAnswerSelectBlock {
						trueAnswerCount++
					}
					answerCountSelectBlockCounter++
				}
			}
		}
	}

	var result string
	percentageCompletion := 100 / questionCount * trueAnswerCount
	if percentageCompletion >= 90 {
		result = "Отлично"
	} else if percentageCompletion < 90 && percentageCompletion >= 75 {
		result = "Хорошо"
	} else if percentageCompletion < 75 && percentageCompletion >= 60 {
		result = "Удовлетворительно"
	} else {
		result = "Неудовлетворительно"
	}

	stmt, ok = query["AddResult"]
	if !ok {
		return "Ошибка"
	}

	_, e = stmt.Exec(UserID.Id, test.TestID, result)
	if e != nil {
		utils.Logger.Println(e)
		return "Ошибка"
	}

	return result
}

func TestLoad(id string) []GetTest {
	var test []GetTest

	stmt, ok := query["GetTestQuestion"]
	if !ok {
		return nil
	}

	rows, e := stmt.Query(id)
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var temp GetTest
		e = rows.Scan(&temp.QuestionID, &temp.Question, &temp.QuestionType)
		if e != nil {
			utils.Logger.Println(e)
			return nil
		}

		if temp.QuestionType == "defBlock" {
			stmt, ok = query["GetDefBlockValue"]
			if !ok {
				return nil
			}

			values, e := stmt.Query(id, temp.QuestionID)
			if e != nil {
				utils.Logger.Println(e)
				return nil
			}

			for values.Next() {
				var value string
				e = values.Scan(&value)
				if e != nil {
					utils.Logger.Println(e)
					return nil
				}
				temp.Value = append(temp.Value, value)
			}

			stmt, ok = query["GetDefBlockAnswer"]
			if !ok {
				return nil
			}

			answer := stmt.QueryRow(id, temp.QuestionID)

			var tempAnswer string
			e = answer.Scan(&tempAnswer)
			if e != nil {
				utils.Logger.Println(e)
				return nil
			}
			temp.Answer = append(temp.Answer, tempAnswer)
		}

		if temp.QuestionType == "checkBox" {
			stmt, ok = query["GetCheckBoxValue"]
			if !ok {
				return nil
			}

			values, e := stmt.Query(id, temp.QuestionID)
			if e != nil {
				utils.Logger.Println(e)
				return nil
			}

			for values.Next() {
				var value string
				e = values.Scan(&value)
				if e != nil {
					utils.Logger.Println(e)
					return nil
				}
				temp.Value = append(temp.Value, value)
			}

			stmt, ok = query["GetCheckBoxAnswer"]
			if !ok {
				return nil
			}

			answers, e := stmt.Query(id, temp.QuestionID)
			if e != nil {
				utils.Logger.Println(e)
				return nil
			}

			for answers.Next() {
				var tempAnswer string
				e = answers.Scan(&tempAnswer)
				if e != nil {
					utils.Logger.Println(e)
					return nil
				}
				temp.Answer = append(temp.Answer, tempAnswer)
			}
		}

		if temp.QuestionType == "selectBlock" {
			stmt, ok = query["GetSelectBlockValue"]
			if !ok {
				return nil
			}

			values, e := stmt.Query(id, temp.QuestionID)
			if e != nil {
				utils.Logger.Println(e)
				return nil
			}

			for values.Next() {
				var value string
				e = values.Scan(&value)
				if e != nil {
					utils.Logger.Println(e)
					return nil
				}
				temp.Value = append(temp.Value, value)
			}

			stmt, ok = query["GetSelectBlockAnswer"]
			if !ok {
				return nil
			}

			answers, e := stmt.Query(id, temp.QuestionID)
			if e != nil {
				utils.Logger.Println(e)
				return nil
			}

			for answers.Next() {
				var tempAnswer string
				e = answers.Scan(&tempAnswer)
				if e != nil {
					utils.Logger.Println(e)
					return nil
				}
				temp.Answer = append(temp.Answer, tempAnswer)
			}
		}

		test = append(test, temp)
	}

	return test
}

func (test *Test) SaveTest() bool {
	stmt, ok := query["AddTestName"]
	if !ok {
		return false
	}

	_, e := stmt.Exec(test.TestName)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	stmt, ok = query["SelectTestId"]
	if !ok {
		return false
	}

	row := stmt.QueryRow(test.TestName)
	e = row.Scan(&test.TestID)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	countQuestion := 1

	if len(test.DefBlockQuestion) > 0 {
		count := 1
		for _, q := range test.DefBlockQuestion {
			stmt, ok = query["AddDefBlockQuestion"]
			if !ok {
				return false
			}

			_, e = stmt.Exec(test.TestID, countQuestion, q)
			if e != nil {
				utils.Logger.Println(e)
				return false
			}

			stmt, ok = query["AddDefBlockAnswer"]
			if !ok {
				return false
			}

			_, e = stmt.Exec(test.TestID, countQuestion, countQuestion, test.DefBlockAnswer[countQuestion-1])
			if e != nil {
				utils.Logger.Println(e)
				return false
			}

			stmt, ok = query["AddDefBlockValue"]
			if !ok {
				return false
			}

			for j := 0; j < 4; j++ {
				_, e = stmt.Exec(test.TestID, countQuestion, count, test.DefBlockValue[count-1])
				if e != nil {
					utils.Logger.Println(e)
					return false
				}
				count++
			}
			countQuestion++
		}
	}

	countQuestionCheckBox := 0
	if len(test.CheckBoxQuestion) > 0 {
		countAnswer := 1
		countValue := 1
		x := 0
		for _, q := range test.CheckBoxQuestion {
			stmt, ok = query["AddCheckBoxQuestion"]
			if !ok {
				return false
			}

			_, e = stmt.Exec(test.TestID, countQuestion, q)
			if e != nil {
				utils.Logger.Println(e)
				return false
			}

			stmt, ok = query["AddCheckBoxAnswer"]
			if !ok {
				return false
			}

			limit, e := strconv.Atoi(test.CheckBoxAnswerCount[countQuestionCheckBox])
			if e != nil {
				utils.Logger.Println(e)
				return false
			}

			for j := 0; j < limit; j++ {
				_, e = stmt.Exec(test.TestID, countQuestion, countAnswer, test.CheckBoxAnswer[j])
				if e != nil {
					utils.Logger.Println(e)
					return false
				}
				countAnswer++
			}

			stmt, ok = query["AddCheckBoxValue"]
			if !ok {
				return false
			}

			limit, e = strconv.Atoi(test.CheckBoxValueCount[countQuestionCheckBox])
			if e != nil {
				utils.Logger.Println(e)
				return false
			}
			limit += x
			for j := x; j < limit; j++ {
				_, e = stmt.Exec(test.TestID, countQuestion, countValue, test.CheckBoxValue[j])
				if e != nil {
					utils.Logger.Println(e)
					return false
				}
				countValue++
				x++
			}
			countQuestion++
			countQuestionCheckBox++
		}
	}

	countSelectBlockQuestion := 0
	selectBlockCount, e := strconv.Atoi(test.SelectBlockQuestionCount[0])
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	if selectBlockCount > 0 {
		x := 0
		count := 1
		for i := 0; i < selectBlockCount; i++ {
			stmt, ok = query["AddSelectBlockQuestion"]
			if !ok {
				return false
			}

			_, e = stmt.Exec(test.TestID, countQuestion, "NULL")
			if e != nil {
				utils.Logger.Println(e)
				return false
			}

			stmt, ok = query["AddSelectBlockValue"]
			if !ok {
				return false
			}

			var stmt1 *sql.Stmt

			stmt1, ok = query["AddSelectBlockAnswer"]
			if !ok {
				return false
			}

			var limit int
			limit, e = strconv.Atoi(test.SelectBlockValeAnswerCount[countSelectBlockQuestion])
			if e != nil {
				utils.Logger.Println(e)
				return false
			}
			limit += x

			for j := x; j < limit; j++ {
				_, e = stmt.Exec(test.TestID, countQuestion, count, test.SelectBlockValue[j])
				if e != nil {
					utils.Logger.Println(e)
					return false
				}
				_, e = stmt1.Exec(test.TestID, countQuestion, count, test.SelectBlockAnswer[j])
				if e != nil {
					utils.Logger.Println(e)
					return false
				}
				count++
				x++
			}

			countSelectBlockQuestion++
			countQuestion++
		}
	}
	return true
}

func GetTestList() []TestList {
	tests := make([]TestList, 0)

	stmt, ok := query["GetTestList"]
	if !ok {
		return nil
	}

	rows, e := stmt.Query()
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var test TestList
		e = rows.Scan(&test.Id, &test.Name)
		if e != nil {
			utils.Logger.Println(e)
			return nil
		}

		tests = append(tests, test)
	}

	return tests
}
