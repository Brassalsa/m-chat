package pkg

import "testing"

func TestTokens(t *testing.T) {
	type testToken struct {
		Id string `json:"id"`
	}

	tstT := testToken{
		Id: "random_test_id",
	}
	t.Log("testing token generate")
	token, err := GenerateJWT(tstT)
	if err != nil {
		t.Error(err)
	}

	tstN := testToken{}
	t.Log("testing token validation")
	if err := ValidateToken(token, &tstN); err != nil {
		t.Error(err)
	}

	if tstT != tstN {
		t.Errorf("have %s want %s", tstN.Id, tstN.Id)
	}

}
