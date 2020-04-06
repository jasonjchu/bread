package tag

import (
	"github.com/jasonjchu/bread/app/db"
	"github.com/jmoiron/sqlx"
)

type Id int
type Name string
type Description string

type Tag struct {
	Id          Id          `db:"_id"`
	Name        Name        `db:"tag_name"`
	Description Description `db:"tag_description"`
}

type Tags[] *Tag

func GetAllJobTags() (Tags, error) {
	pool := db.Pool
	row, err := pool.Queryx("SELECT * FROM jobTags")
	if err != nil {
		return nil, err
	}

	tags, err := scanTagsFromRows(row)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func scanTagsFromRows(rows *sqlx.Rows) (Tags, error) {
	var tags Tags
	var err error
	for rows.Next() {
		tag := Tag{}
		err = rows.StructScan(&tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}
	return tags, err
}