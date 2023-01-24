package questions

import (
	"fmt"
	"log"
	"strconv"

	sqlite "github.com/diogo-costa98/GoGraphQL/internal/db/sqlite"
	"github.com/diogo-costa98/GoGraphQL/internal/users"
)

type Question struct {
	ID   string
	Body string
	User *users.User
}

type Option struct {
	ID       string
	Body     string
	Correct  bool
	Question *Question
}

func (question Question) Save() int64 {

	var id int64
	query := fmt.Sprintf("INSERT INTO questions(body,user_id) VALUES ('%v','%v') RETURNING id", question.Body, question.User.ID)
	err := sqlite.Db.QueryRow(query).Scan(&id)
	if err != nil {
		log.Fatal("QuestionSave QueryRow Err: ", err)
	}
	log.Printf("Question %d inserted.", id)
	return id
}

func (question Question) Update() int64 {

	var id int64
	query := fmt.Sprintf("UPDATE questions SET body = '%v' WHERE id = '%v' RETURNING id", question.Body, question.ID)
	err := sqlite.Db.QueryRow(query).Scan(&id)
	if err != nil {
		log.Fatal("QuestionUpdate QueryRow Err: ", err)
	}
	log.Printf("Question %d updated.", id)
	return id
}

func (question Question) Delete() int64 {

	var id int64
	query := fmt.Sprintf("DELETE FROM questions WHERE id = '%v' RETURNING id", question.ID)
	err := sqlite.Db.QueryRow(query).Scan(&id)
	if err != nil {
		log.Fatal("QuestionDelete QueryRow Err: ", err)
	}
	log.Printf("Question %d deleted.", id)
	return id
}

func GetAll(page *string, pageSize *string, userId *string) (map[string]Question, map[string]Option) {
	var firstId int64
	var paginationSize int64
	var err error

	//Calculate page size
	if pageSize == nil || *pageSize == "" {
		paginationSize = 10
	} else {
		paginationSize, err = strconv.ParseInt(*pageSize, 10, 64)
		if err != nil {
			log.Fatal("Questions ParseInt Err: ", err)
		}
	}

	//Determine page key-set
	if page == nil || *page == "" {
		firstId = 0
	} else {
		pageN, err := strconv.ParseInt(*page, 10, 64)
		if err != nil {
			log.Fatal("Questions ParseInt Err: ", err)
		}
		firstId = pageN * paginationSize
	}

	stmt, err := sqlite.Db.Prepare("SELECT Q.id, Q.body, O.id, O.body, O.correct FROM questions Q INNER JOIN options O ON Q.id = O.question_id WHERE Q.user_id = ? ORDER BY Q.id ASC LIMIT ?,?")
	if err != nil {
		log.Fatal("Questions Preparation Err: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, firstId, paginationSize)
	if err != nil {
		log.Fatal("QuestionsGetAll Query Err: ", err)
	}
	defer rows.Close()

	questions := make(map[string]Question)
	options := make(map[string]Option)
	for rows.Next() {
		var question Question
		var option Option

		err := rows.Scan(&question.ID, &question.Body, &option.ID, &option.Body, &option.Correct)
		if err != nil {
			log.Fatal(err)
		}
		option.Question = &question

		questions[question.ID] = question
		options[option.ID] = option
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return questions, options
}

// GetQuestionById gets a question using a specific id
func GetQuestionById(questionId string) (Question, []Option) {
	statement, err := sqlite.Db.Prepare("SELECT Q.id, Q.body, Q.user_id, O.id, O.body, O.correct FROM questions Q INNER JOIN options O ON Q.id = O.question_id AND Q.id = ? ORDER BY O.id")
	if err != nil {
		log.Fatal("GetQuestionById Prepare Err: ", err)
	}

	rows, err := statement.Query(questionId)
	if err != nil {
		log.Fatal("GetQuestionById Query Err: ", err)
	}

	defer rows.Close()
	var question Question
	question.User = &users.User{}
	var options []Option

	for rows.Next() {
		var option Option
		err := rows.Scan(&question.ID, &question.Body, &question.User.ID, &option.ID, &option.Body, &option.Correct)
		if err != nil {
			log.Fatal(err)
		}
		options = append(options, option)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return question, options
}
