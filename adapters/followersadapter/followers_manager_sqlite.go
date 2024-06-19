package followersadapter

import (
	"database/sql"
	"errors"
	"sus/nilable"
	"sus/ports/followersport"
	"sus/services/sqliteservice"
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
	return errors.New("not yet implemented")
}

func (f FollowersManagerSQLite) GetFollowers(lastFollowerID nilable.Nilable[string]) ([]string, error) {
	return nil, errors.New("not yet implemented")
}
