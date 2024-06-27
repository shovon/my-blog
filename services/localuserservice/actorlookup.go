package actorlookupservice

import (
	"fmt"
	"strings"
	"sus/services/originservice"
)

func DoesActorIDByLocalActorExists(id string) bool {
	id = strings.TrimSpace(id)
	return id == originservice.GetHTTPOrigin() || id == fmt.Sprintf("%s/", originservice.GetHTTPOrigin())
}
