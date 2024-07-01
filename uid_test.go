package uid

import "testing"

var (
	testOption = Option{
		Chars: []rune("0123456789ABCDEFHKLMNPQRSTUVWXYZabcdefghkmnpqrstuvwxyz"),
		Cnt:   10,
	}
	testUID *UID
)

func TestNewUID(t *testing.T) {
	uid, err := NewUID(testOption)
	if err != nil {
		t.Fatal(err)
	}
	testUID = uid
}

func TestGenerate(t *testing.T) {
	uid, err := testUID.Generate()
	if err != nil {
		t.Fatal(err)
	}
	if len(uid) != 10 {
		t.Fatalf("uid is not len %d", testOption.Cnt)
	}
}
