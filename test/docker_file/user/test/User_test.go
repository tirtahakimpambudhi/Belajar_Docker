package test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWhoami(t *testing.T)  {
	cmd := exec.Command("whoami")
	user, err := cmd.Output()
	require.NoError(t,err)
	require.Equal(t,"tirtahakim\n",string(user))
}