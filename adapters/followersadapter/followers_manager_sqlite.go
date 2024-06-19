package followersadapter

import (
	"database/sql"
	"errors"
	"sus/nilable"
	"sus/ports/followersport"
	"sus/services/sqliteservice"

	"github.com/mattn/go-sqlite3"
)

type FollowersManagerSQLite struct {
	writer *sql.DB
	reader *sql.DB
}

func NewFollowersManagerSQLite(
	writer sqliteservice.Writer,
	reader sqliteservice.Reader,
) FollowersManagerSQLite {
	return FollowersManagerSQLite{writer: writer, reader: reader}
}

var _ followersport.FollowersManager = FollowersManagerSQLite{}

func (f FollowersManagerSQLite) SaveFollower(followerID string) error {
	r, err := f.writer.Exec("insert into followers (follower_id_iri) values ($1);", followerID)
	if err != nil {
		return err
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		sqliteErr, ok := err.(sqlite3.Error)
		if ok && sqliteErr.Code == sqlite3.ErrConstraint && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return nil
		}
		return err
	}
	if rowsAffected == 0 {
		return errors.New("faled to save follower")
	}
	return nil
}

func (f FollowersManagerSQLite) GetFollowers(lastFollowerID nilable.Nilable[string]) ([]string, error) {
	return nil, errors.New("not yet implemented")
}
