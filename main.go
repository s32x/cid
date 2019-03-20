package main /* import "s32x.com/cid" */

import (
	"log"
	"os"

	"s32x.com/cid/cid"
)

func main() {
	cid.Start(
		getenv("USER_URL"),     // Where to retrieve the requested repository from
		getenv("DOMAIN"),       // The domain this service will be being hosted on
		getenv("PORT", "8080"), // The port this service will be hosted on
	)
}

// getenv attempts to retrieve and return a variable from the environment. If it
// fails it will either crash or failover to a passed default value
func getenv(key string, def ...string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	if len(def) == 0 {
		log.Fatalf("%s not defined in environment", key)
	}
	return def[0]
}
