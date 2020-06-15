package service

import (
	"fmt"
	"pm5-emulator/config"
)

//getFullUUID returns 128bit UUID
func getFullUUID(uuid string) string {
	return fmt.Sprintf("%s%s%s", config.UUID_PREFIX, uuid, config.UUID_SUFFIX)
}
