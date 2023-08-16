package constants

import "time"

const JWTExpireDuration = 1 * time.Hour
const PasswordHashCost = 14

const UserClaims = "userClaims"
const DefaultTimeZone = "UTC"
