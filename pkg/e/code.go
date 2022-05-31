package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	ErrorExistUser      = 10001
	ErrorNotExistUser   = 10002
	ErrorFailEncryption = 10003
	ErrorNotCompare     = 10004

	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuthAuth              = 20004

	ErrorDatabase = 30001
)
