package repository

import (
	"database/sql"
	"fmt"
	"github.com/gchaincl/dotsql"
	_ "github.com/lib/pq"
	"github.com/mfaulther/avito-tradex-test-task/internal/app/model"
	"log"
)

type StatRepository struct {
	db *sql.DB
}

func New(DatabaseURL string) (*StatRepository, error) {

	db, err := sql.Open("postgres", DatabaseURL)

	if err != nil {
		return nil, err
	}

	dot, err := dotsql.LoadFromFile("./create-tables.sql")

	res, err := dot.Exec(db, "create-statistics-table")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	return &StatRepository{
		db: db,
	}, nil

}

func (r *StatRepository) GetStatistics() []model.Statistics {

	res := make([]model.Statistics, 0)

	rows, err := r.db.Query("SELECT stat_date, SUM(views), SUM(clicks), SUM(cost) FROM statistics GROUP BY stat_date ORDER BY stat_date;")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var stat model.Statistics
		err := rows.Scan(&stat.Time.Time, &stat.Views, &stat.Clicks, &stat.Cost)
		if err != nil {
			log.Fatal(err)
		}
		stat.Cpc = stat.Cost / float32(stat.Clicks)
		stat.Cpm = stat.Cost / float32(stat.Views)
		res = append(res, stat)
	}

	return res
}

func (r *StatRepository) AddStatistics(stat *model.Statistics) {

	query := "INSERT INTO statistics (stat_date, views, clicks, cost, cpc, cpm) " +
		"VALUES ($1, $2, $3, $4, $5, $6)"

	stat.Cpc = stat.Cost / float32(stat.Clicks)
	stat.Cpm = stat.Cost / float32(stat.Views)

	stmt, err := r.db.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(stat.Time.Time, stat.Views, stat.Clicks, stat.Cost, stat.Cpc, stat.Cpm)

	if err != nil {
		log.Fatal(err)
	}

}
