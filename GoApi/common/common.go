package common

const (
	SuccessCode = "S"
	ErrorCode   = "E"
	Pending     = "Pending"
	Approved    = "Approved"
	Rejected    = "Rejected"
)

func ValidateUserRole(lRole string) bool {
	if lRole == "" || lRole != "BO" || lRole != "B" || lRole != "APPR" {
		return false
	}
	return true
}
