package smtp_codes

type SmtpCode struct {
	Code          int16
	EnhancedCodes [3]int16
}

var NoEnhancedCodes = [3]int16{0, 0, 0}
