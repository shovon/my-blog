package followersadapter

import (
	"database/sql"
	"errors"
	"sus/nilable"
	"sus/ports/followersport"
	"sus/services/sqliteservice"
	"time"

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

func (f FollowersManagerSQLite) GetFollowers(
	lastFollowerID nilable.Nilable[string],
	queryParameters followersport.QueryParameters,
) ([]followersport.FollowerMeta, error) {
	rows, err := f.reader.Query(`
		select
			id,
			follower_id_iri,
			when_followed
		from followers
		where $2 < id and id < $3
		order by id desc limit $1;
	`,
		queryParameters.Limit.ValueOrDefault(40),
		queryParameters.Limit.ValueOrDefault(0),
		queryParameters.Limit.ValueOrDefault(9223372036854775807))
	if err != nil {
		return nil, err
	}

	followers := []followersport.FollowerMeta{}

	for rows.Next() {
		var id string
		var followerIDIRI string
		var whenFollowed time.Time
		if err := rows.Scan(&id, &followerIDIRI, &whenFollowed); err != nil {
			return nil, err
		}

		followers = append(followers, followersport.FollowerMeta{
			ID:            id,
			FollowerIDIRI: followerIDIRI,
			WhenFollowed:  whenFollowed,
		})
	}

	return followers, nil
}
