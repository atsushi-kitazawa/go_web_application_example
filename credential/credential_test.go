package credential

import (
	"fmt"
	"testing"
)

func TestUnmarshall(t *testing.T) {
    cred := Load()
    fmt.Println(cred.SecurityKey)
    fmt.Println(cred.ClientId)
    fmt.Println(cred.Secret)
}
