package options

import (
	"fmt"
	"log"

	"github.com/diogo-costa98/GoGraphQL/internal/db/sqlite"
	"github.com/diogo-costa98/GoGraphQL/internal/questions"
)

type Option struct {
	ID       string
	Body     string
	Correct  bool
	Question *questions.Question
}

func (option Option) Save() int64 {

	var id int64

	query := fmt.Sprintf("INSERT INTO options(body,correct,question_id) VALUES ('%v','%v','%v') RETURNING id", option.Body, option.Correct, option.Question.ID)

	err := sqlite.Db.QueryRow(query).Scan(&id)
	if err != nil {
		log.Fatal("OptionSave QueryRow Err: ", err)
	}

	log.Printf("Option %d inserted.", id)
	return id
}

func GetByQuestionId(questionID string) []Option {
	stmt, err := sqlite.Db.Prepare("SELECT O.id, O.body, O.correct FROM options O WHERE O.question_id = ?")
	if err != nil {
		log.Fatal("OptionsGetByQuestionId Preparation Err: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(questionID)
	if err != nil {
		log.Fatal("OptionsGetByQuestionId Query Err: ", err)
	}

	defer rows.Close()
	var options []Option

	for rows.Next() {
		var option Option
		err := rows.Scan(&option.ID, &option.Body, &option.Correct)
		if err != nil {
			log.Fatal(err)
		}

		options = append(options, option)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return options
}

func Delete(optionsList []questions.Option) {

	if len(optionsList) == 0 {
		log.Panic("OptionsDelete len(optionsList) Err: trying to delete 0 options")
	}

	query := fmt.Sprintf("DELETE FROM options WHERE")
	for i, option := range optionsList {
		if i != 0 {
			query = fmt.Sprint(query, " OR")
		}
		query = fmt.Sprint(query, " id = ", option.ID)
	}
	_, err := sqlite.Db.Exec(query)
	if err != nil {
		log.Fatal("OptionsDelete Execution Err: ", err)
	}

	return
}
