package followersport

import (
	"sus/nilable"
	"time"
)

type QueryParameters struct {
	MaxID nilable.Nilable[int]
	MinID nilable.Nilable[int]
	Limit nilable.Nilable[int]
}

type FollowerMeta struct {
	ID            string    `json:"id"`
	FollowerIDIRI string    `json:"followerIdIri"`
	WhenFollowed  time.Time `json:"whenFollowed"`
}

type FollowersManager interface {
	SaveFollower(followerID string) error
	GetFollowers(
		lastFollowerID nilable.Nilable[string],
		queryParameters QueryParameters,
	) ([]FollowerMeta, error)
}
