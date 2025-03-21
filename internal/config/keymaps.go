// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package config

import (
	"errors"
	"io/fs"
	"log/slog"
	"os"

	"github.com/derailed/k9s/internal/config/data"
	"github.com/derailed/k9s/internal/config/json"
	"github.com/derailed/k9s/internal/slogs"
	"gopkg.in/yaml.v3"
)

// Keymaps represents a collection of keymaps.
type Keymaps struct {
	Keymaps map[string]Keymap `yaml:"keymaps"`
}

// Keymap describes a key to action mapping.
type Keymap struct {
	Key string `yaml:"key"`
}

// NewKeymaps returns a new keymapping configuration.
func NewKeymaps() Keymaps {
	return Keymaps{
		Keymaps: make(map[string]Keymap),
	}
}

// Load keymaps.
func (k Keymaps) Load(path string) error {
	if err := k.LoadKeymaps(AppHotKeysFile); err != nil {
		return err
	}
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return nil
	}

	return k.LoadKeymaps(path)
}

// LoadKeymaps loads keymaps from a given file.
func (k Keymaps) LoadKeymaps(path string) error {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	bb, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := data.JSONValidator.Validate(json.KeymapsSchema, bb); err != nil {
		slog.Warn("Validation failed. Please update your config and restart.",
			slogs.Path, path,
			slogs.Error, err,
		)
	}

	var kk Keymaps
	if err := yaml.Unmarshal(bb, &kk); err != nil {
		return err
	}

	for action, v := range kk.Keymaps {
		k.Keymaps[action] = v
	}

	return nil
}
