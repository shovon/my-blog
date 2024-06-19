package sqliteservice

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Writer *sql.DB
type Reader *sql.DB

func migrations() []string {
	return []string{
		`
		create table followers (
			id integer primary key autoincrement,
			follower_id_iri text not null,
			when_followed timestamp not null default current_timestamp
		);
		`,
		`
		create table followings (
			id integer primary key autoincrement,
			following_id_iri text not null,
			when_followed timestamp not null default current_timestamp
		);
		`,
	}
}

func sqliteLocation() string {
	v, ok := os.LookupEnv("SQLITE_PATH")
	if ok {
		return v
	}

	return "database.db"
}

var writeable *sql.DB
var getWriteableDBLock sync.Mutex

var readable *sql.DB
var getReadableDBLock sync.Mutex

func GetWriteableDB() Writer {
	getWriteableDBLock.Lock()
	defer getWriteableDBLock.Unlock()

	if writeable == nil || writeable.Stats().MaxOpenConnections == 0 {
		// Opens a read-write connection
		w, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_mode=rw&_journal=WAL&_timeout=5000", sqliteLocation()))

		w.SetMaxOpenConns(1)
		w.SetMaxIdleConns(1)
		if err != nil {
			panic(err)
		}
		writeable = w
	}

	return writeable
}

func GetReadableDB() Reader {
	getReadableDBLock.Lock()
	defer getReadableDBLock.Unlock()

	if readable == nil || readable.Stats().MaxOpenConnections == 0 {
		r, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_mode=ro&_journal=WAL", sqliteLocation()))
		r.SetMaxOpenConns(64)
		r.SetMaxIdleConns(64)
		r.SetConnMaxIdleTime(0)
		if err != nil {
			panic(err)
		}
		readable = r
	}

	return readable
}

func createSqliteDatabase() {
	filename := sqliteLocation()
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		file.Close()
	}
}

func init() {

	createSqliteDatabase()

	w := (*sql.DB)(GetWriteableDB())
	r := (*sql.DB)(GetReadableDB())

	_, err := w.Exec(`
	create table if not exists migrations (
		id integer primary key autoincrement,
		when_run timestamp not null default current_timestamp
	);
	`)
	fmt.Println("Idempotently created a new table")
	if err != nil {
		panic(err)
	}

	var lastMigrationId int

	row := r.QueryRow("select id from migrations order by id desc;")
	if err := row.Scan(&lastMigrationId); err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	fmt.Printf("Last migration ID: %d\n", lastMigrationId)

	m := migrations()
	for i := lastMigrationId; i < len(m); i++ {
		tx, err := w.Begin()
		if err != nil {
			panic(err)
		}
		_, err = tx.Exec(m[i])
		if err != nil {
			panic(err)
		}

		tx.Exec("insert into migrations default values;")
		tx.Commit()
	}
}
