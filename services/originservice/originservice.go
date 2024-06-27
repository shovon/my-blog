package originservice

import (
	"fmt"
	"sus/config"
)

func GetHTTPOrigin() string {
	var protocol string
	if config.IsSecure() {
		protocol = "https"
	} else {
		protocol = "http"
	}
	return fmt.Sprintf("%s://%s", protocol, config.Host())
}
