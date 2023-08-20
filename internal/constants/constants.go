package constants

import "time"

const JWTExpireDuration = 1 * time.Hour
const PasswordHashCost = 14

const AuthorizationHeaderKey = "Authorization"
const AuthorizationTypeBearer = "bearer"
const AuthorizationPayload = "AuthorizationPayload"
const DefaultTimeZone = "UTC"
