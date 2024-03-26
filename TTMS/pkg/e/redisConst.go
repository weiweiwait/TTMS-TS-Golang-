package e

import "time"

const (
	VerificationCodeKey    = "ttms:user:code:"
	VerificationCodeKeyTTL = time.Second * 60 * 3

	UserLoginInfo    = "ttms:user:token:"
	UserLoginInfoTTL = time.Hour * 24 * 30

	PetHotData    = "ttms:hotData:pet:"
	PetHotDataDDL = time.Second * 60
)
