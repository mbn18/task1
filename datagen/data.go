package datagen

import (
	"github.com/mbn18/dream/internal/entity"
	"log"
	"math/rand/v2"
	"time"
)

const (
	david  = "David"
	miriam = "Miriam"
	admin  = "Admin"
)

var usernames = []string{david, miriam, admin}

var users = map[string]entity.User{
	david: {
		Meta: map[string]any{
			"name":       david,
			"occupation": "Student",
			"age":        "23",
		},
	},
	miriam: {
		Meta: map[string]any{
			"name":       miriam,
			"occupation": "Student",
			"age":        "23",
		},
	},
	admin: {
		Meta: map[string]any{
			"name":       admin,
			"occupation": "Maintainer",
			"age":        "33",
		},
	},
}

var hosts = []*entity.Host{
	{
		ID: 1,
		OS: entity.Linux,
		Meta: map[string]any{
			"owner":           "University ABC",
			"department":      "physics",
			"last_maintained": time.Now().Add(time.Hour * 24 * 30 * 3),
			"manufacturer":    "Dell",
		},
	},
	{
		ID: 2,
		OS: entity.Darwin,
		Meta: map[string]any{
			"owner":        david,
			"manufacturer": "Lenovo",
		},
	},
	{
		ID: 3,
		OS: entity.Linux,
		Meta: map[string]any{
			"owner":        miriam,
			"manufacturer": "Lenovo",
		},
	},
}

func Generate(days int) (host *entity.Host) {

	user := usernames[rand.IntN(len(usernames))]
	switch user {
	case david, miriam:
		host = genHost(user)
	case admin:
		host = hosts[0]
		host.User = users[admin]
	default:
		log.Fatal("switch usernames failed")
	}

	executedAt := time.Now().Add(-time.Hour * 24 * time.Duration(days))
	host.Processes = genProcessList(executedAt)

	log.Printf("Generated %d processors for host %d of %s type by user %s", len(host.Processes.Processes), host.ID, host.OS, host.User.Meta["name"])
	return
}

func genHost(student string) (host *entity.Host) {
	if rand.IntN(2) == 1 {
		host = hosts[0]
	} else {
		host = hosts[2]
	}
	host.User = users[student]
	return host
}
