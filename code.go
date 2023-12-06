package smtp_codes

import "fmt"

type SmtpCode struct {
	Code          int16
	EnhancedCodes [3]int16
}

var NoEnhancedCodes = [3]int16{0, 0, 0}

func (s *SmtpCode) GetCode() string {

	if s.EnhancedCodes != NoEnhancedCodes {
		return fmt.Sprintf("%d %d.%d.%d", s.Code, s.EnhancedCodes[0], s.EnhancedCodes[1], s.EnhancedCodes[2])

	}

	return fmt.Sprintf("%d", s.Code)
}
