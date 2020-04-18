package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"server/models"
)

type PgRepository struct {
	database *sql.DB
}

func NewPgRepository(url string) (*PgRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if _, err = db.Exec(`create table if not exists reviewers(id serial primary key, info json, name text)`); err != nil {
		return nil, err
	}

	repository := &PgRepository{
		database: db,
	}

	return repository, nil
}

func (this *PgRepository) Add(reviewer *models.Reviewer) error {
	jsonInfo, err := json.Marshal(reviewer.Info)
	if err != nil {
		return err
	}

	_, err = this.database.Exec(fmt.Sprintf("insert into reviewers (info, name) values ('%s', '%s')", string(jsonInfo), reviewer.Name))
	return err
}

func (this *PgRepository) Remove(reviewer models.Reviewer) error {
	jsonInfo, err := json.Marshal(reviewer.Info)
	if err != nil {
		return err
	}

	_, err = this.database.Exec(fmt.Sprintf("delete from reviewers where info='%s'", jsonInfo))
	return err
}

func (this *PgRepository) Check(reviewer models.Reviewer) error {
	jsonInfo, err := json.Marshal(reviewer.Info)
	if err != nil {
		return err
	}

	row := this.database.QueryRow("select id, name from reviewers where info='%s'", jsonInfo)
	return row.Scan(reviewer.Id, reviewer.Name)
}

func (this *PgRepository) All(mapper models.InfoMapper) (result []models.Reviewer, err error) {
	rows, err := this.database.Query("select * from reviewers")
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		reviewer := &models.Reviewer{}

		var info []byte
		err = rows.Scan(&reviewer.Id, &info, &reviewer.Name)
		if err != nil {
			return
		}

		reviewer.Info = mapper(info)
		result = append(result, *reviewer)
	}

	return
}

func (this *PgRepository) Close() error {
	return this.database.Close()
}
