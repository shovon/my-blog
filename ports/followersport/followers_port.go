package followersport

import "sus/nilable"

type FollowersManager interface {
	SaveFollower(followerID string) error
	GetFollowers(lastFollowerID nilable.Nilable[string]) ([]string, error)
}
