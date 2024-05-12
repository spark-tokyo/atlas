package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/assert"
)

func TestNewLoadConfig(t *testing.T) {
	// テストケースを定義
	tests := []struct {
		name        string
		envStage    string
		envFilePath string
		expectError bool
	}{
		{
			name:        "local environment",
			envStage:    appLocal,
			envFilePath: "../local.env",
			expectError: false,
		},
		{
			name:        "production environment",
			envStage:    appLocal,
			envFilePath: "../.env",
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 環境変数を設定
			require.NoError(t, os.Setenv("STAGE", tc.envStage))

			os.Setenv("STAGE", "local")
			cfg, err := NewLoadConfig()

			assert.Equal(t, cfg.Stage, tc.envStage)

			// エラー発生の期待値を検証
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// 環境変数をクリーンアップ
			require.NoError(t, os.Unsetenv("STAGE"))
		})
	}
}
