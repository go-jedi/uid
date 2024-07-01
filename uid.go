package uid

import (
	"crypto/rand"
	"math/big"
)

var (
	defChars = []rune("0123456789ABCDEFHKLMNPQRSTUVWXYZabcdefghkmnpqrstuvwxyz")
	defCnt   = 15
)

type UID struct {
	chars []rune
	cnt   int
}

func (u *UID) init() error {
	if len(u.chars) == 0 {
		u.chars = defChars
	}

	if u.cnt <= 0 {
		u.cnt = defCnt
	}

	return nil
}

func NewUID(opt Option) (*UID, error) {
	u := &UID{
		chars: opt.Chars,
		cnt:   opt.Cnt,
	}

	if err := u.init(); err != nil {
		return nil, err
	}

	return u, nil
}

// Generate unique uid
func (u *UID) Generate() (string, error) {
	// characters used in UID
	uid := make([]rune, u.cnt)

	for i := range uid {
		// generate a random index and select a character from chars
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(u.chars))))
		if err != nil {
			return "", err
		}

		uid[i] = u.chars[num.Int64()]
	}

	return string(uid), nil
}
