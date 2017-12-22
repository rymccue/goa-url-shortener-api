package repositories

import (
	"database/sql"
	"github.com/rymccue/goa-url-shortener-api/app"
)

func CreateURL(db *sql.DB, url, path string) error {
	const query = `
		insert into urls (
			url,
			path
		) values (
			$1,
			$2
		)
	`
	_, err := db.Exec(query, url, path)
	return err
}

func GetURL(db *sql.DB, path string) (*app.URL, error) {
	const query = `
		select
			id,
			url,
			path
		from urls
		where
			path = $1
	`
	var url app.URL
	err := db.QueryRow(query, path).Scan(&url.ID, &url.URL, &url.Path)
	return &url, err
}

func UpdateURL(db *sql.DB, url, path string) error {
	const query = `
		update urls set
			url = $1,
			path = $2
		where path = $2
	`
	_, err := db.Exec(query, url, path)
	return err
}

func DeleteURL(db *sql.DB, path string) error {
	const query = `delete from urls where path = $1`
	_, err := db.Exec(query, path)
	return err
}

func AddUrlHit(db *sql.DB, path string) error {
	const query = `
		insert into analytics (
			url_id
		) values (
			(select id from urls where path = $1)
		)
	`
	_, err := db.Exec(query, path)
	return err
}

func GetAnalytics(db *sql.DB, path string) (*app.Analytics, error) {
	const query = `
		select
			count(analytics.*)
		from analytics
		join urls on urls.id = analytics.url_id
		where urls.path = $1
	`
	var analytics app.Analytics
	err := db.QueryRow(query, path).Scan(&analytics.Hits)
	return &analytics, err
}
