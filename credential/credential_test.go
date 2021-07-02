package credential

import (
	"fmt"
	"testing"
)

func TestUnmarshall(t *testing.T) {
    cred := Load()
    fmt.Printf(cred.SecurityKey)
    fmt.Printf(cred.ClientId)
    fmt.Printf(cred.Secret)
}
