package auth

import "testing"

func TestSimpleAuth(t *testing.T) {
	a := NewSimpleAuth()
	token, err := a.Login("user", "pass")
	if err != nil {
		t.Fatal(err)
	}
	ok, err := a.Verify(token.AccessToken)
	if err != nil || !ok {
		t.Fatalf("verify failed: %v %v", ok, err)
	}
	oldToken := token.AccessToken
	newToken, err := a.Refresh(token.RefreshToken)
	if err != nil {
		t.Fatal(err)
	}
	if newToken.AccessToken == oldToken {
		t.Fatalf("expected new access token")
	}
}
