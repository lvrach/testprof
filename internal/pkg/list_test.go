package pkg_test

import (
	"testing"

	"github.com/lvrach/testprof/internal/pkg"
	"github.com/stretchr/testify/require"
)

func Test_List(t *testing.T) {
	pkgs, err := pkg.List()
	require.NoError(t, err)
	require.Equal(t, []string{
		"github.com/lvrach/testprof/internal/pkg",
	}, pkgs)
}
