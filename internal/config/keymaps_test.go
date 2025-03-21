// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package config_test

import (
	"fmt"
	"testing"

	"github.com/derailed/k9s/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestKeymapsLoad(t *testing.T) {
	k := config.NewKeymaps()
	assert.NoError(t, k.LoadKeymaps("testdata/keymaps/keymaps.yaml"))
	fmt.Printf("keymaps: %#v\n", k)

	assert.Equal(t, 1, len(k.Keymaps))

	km, ok := k.Keymaps["view.details.back"]
	assert.True(t, ok)
	assert.Equal(t, "q", km.Key)
}
