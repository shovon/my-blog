package followersadapter

import (
	"database/sql"
	"errors"
	"sus/nilable"
	"sus/ports/followersport"
)

type FollowersManagerSQLite struct {
	db *sql.DB
}

func NewFollowersManagerSQLite(db *sql.DB) FollowersManagerSQLite {
	return FollowersManagerSQLite{db: db}
}

var _ followersport.FollowersManager = FollowersManagerSQLite{}

func (f FollowersManagerSQLite) SaveFollower(followerID string) error {
	return errors.New("not yet implemented")
}

func (f FollowersManagerSQLite) GetFollowers(lastFollowerID nilable.Nilable[string]) ([]string, error) {
	return nil, errors.New("not yet implemented")
}
