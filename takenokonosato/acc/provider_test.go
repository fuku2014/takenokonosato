package acc

import (
	"testing"

	"github.com/fuku2014/terraform-provider-takenokonosato/takenokonosato"
)

func TestProvider(t *testing.T) {
	if err := takenokonosato.Provider().InternalValidate(); err != nil {
		t.Fatal(err.Error())
	}
}
