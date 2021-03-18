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

func New(config Config) (*StatRepository, error) {

	DatabaseURL := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Passw, config.DatabaseName)

	db, err := sql.Open("postgres", DatabaseURL)

	if err != nil {
		return nil, err
	}

	dot, err := dotsql.LoadFromFile("./create-tables.sql")

	_, err = dot.Exec(db, "create-statistics-table")

	if err != nil {
		log.Fatal(err)
	}

	return &StatRepository{
		db: db,
	}, nil

}

func (r *StatRepository) GetStatistics(from string, to string, s string) []model.Statistics {

	res := make([]model.Statistics, 0)

	if len(s) == 0 {
		s = "stat_date"
	}

	query := "SELECT * FROM statistics WHERE stat_date BETWEEN $1 AND $2 ORDER BY stat_date"

	rows, err := r.db.Query(query, from, to)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var stat model.Statistics
		err := rows.Scan(&stat.Date, &stat.Views, &stat.Clicks, &stat.Cost, &stat.Cpc, &stat.Cpm)
		if err != nil {
			log.Fatal(err)
		}
		stat.Date = stat.Date[:10]
		if stat.Clicks > 0 {
			stat.Cpc = stat.Cost / float32(stat.Clicks)
		}
		if stat.Views > 0 {
			stat.Cpm = stat.Cost / (float32(stat.Views) * 1000)
		}
		res = append(res, stat)
	}

	return res
}

/*
func (r *StatRepository) AddStatistics(stat *model.Statistics) {

	query := "INSERT INTO statistics (stat_date, views, clicks, cost, cpc, cpm) " +
		"VALUES ($1, $2, $3, $4, $5, $6)"

	if stat.Clicks > 0 {
		stat.Cpc = stat.Cost / float32(stat.Clicks)
	}
	if stat.Views > 0 {
		stat.Cpm = stat.Cost / (1000 * float32(stat.Views))
	}

	stmt, err := r.db.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(stat.Date, stat.Views, stat.Clicks, stat.Cost, stat.Cpc, stat.Cpm)

	if err != nil {
		log.Fatal(err)
	}

} */

func (r *StatRepository) AddStatistics(newStat *model.Statistics) {

	insertQuery := "INSERT INTO statistics (stat_date, views, clicks, cost, cpc, cpm) " +
		"VALUES ($1, $2, $3, $4, $5, $6)"

	updateQuery := "UPDATE statistics SET views=$1, clicks=$2, cost=$3, cpc=$4, cpm=$5 " +
		"WHERE stat_date=$6"

	updStmt, err := r.db.Prepare(updateQuery)
	if err != nil {
		log.Fatal(err)
	}
	insStmt, err := r.db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err)
	}

	//Смотрим есть ли в таблице строка с данной датой

	var isExisted int
	row := r.db.QueryRow("SELECT 1 FROM statistics WHERE stat_date=$1", newStat.Date)
	row.Scan(&isExisted)

	//Если есть, то делаем UPDATE, иначе INSERT

	if isExisted > 0 {

		var oldStat model.Statistics

		oldRow := r.db.QueryRow("SELECT * FROM statistics WHERE stat_date=$1", newStat.Date)
		oldRow.Scan(&oldStat.Date, &oldStat.Views, &oldStat.Clicks, &oldStat.Cost, &oldStat.Cpc, &oldStat.Cpm)

		newStat.Views += oldStat.Views
		newStat.Clicks += oldStat.Clicks
		newStat.Cost += oldStat.Cost
		if newStat.Clicks > 0 {
			newStat.Cpc = newStat.Cost / float32(newStat.Clicks)
		}
		if newStat.Views > 0 {
			newStat.Cpm = newStat.Cost / (float32(newStat.Views) * 1000)
		}

		_, err := updStmt.Exec(newStat.Views, newStat.Clicks, newStat.Cost, newStat.Cpc, newStat.Cpm, newStat.Date)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		_, err = insStmt.Exec(newStat.Date, newStat.Views, newStat.Clicks, newStat.Cost, newStat.Cpc, newStat.Cpm)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func (r *StatRepository) DeleteStatistics() {

	res, err := r.db.Exec("TRUNCATE statistics")
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}

}
