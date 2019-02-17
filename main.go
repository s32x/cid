package main /* import "s32x.com/cid" */

import (
	"log"
	"os"

	"s32x.com/cid/cid"
)

func main() {
	cid.Start(
		getenv("USER_URL"), // Where to retrieve the requested repository from
		getenv("DOMAIN"),   // The domain this service will be being hosted on
		getenv("PORT"),     // The port this service will be hosted on
	)
}

// getenv attempts to retrieve an variable from the environment, throwing an
// error if the key doesn't exist
func getenv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("%s not found in environment", key)
	}
	return value
}
