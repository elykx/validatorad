package validatorad

import "errors"

var (
	errorEmptyTitle = errors.New("empty title")
	errorEmptyBody  = errors.New("empty body")
	errorLongTitle  = errors.New("too long title")
	errorLongBody   = errors.New("too long body")
)

func ValidateAd(title, body string, maxTitle, maxBody int) error {
	if title == "" {
		return errorEmptyTitle
	}

	if body == "" {
		return errorEmptyBody
	}

	if len(title) > maxTitle {
		return errorLongTitle
	}

	if len(body) > maxBody {
		return errorLongBody
	}

	return nil
}
