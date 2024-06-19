package followersport

import "sus/nilable"

type QueryParameters struct {
	MaxID nilable.Nilable[int]
	MinID nilable.Nilable[int]
	Limit nilable.Nilable[int]
}

type FollowersManager interface {
	SaveFollower(followerID string) error
	GetFollowers(
		lastFollowerID nilable.Nilable[string],
		queryParameters QueryParameters,
	) ([]string, error)
}
