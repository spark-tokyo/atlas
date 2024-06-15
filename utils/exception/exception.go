package exception

import (
	"fmt"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ErrorCode はカスタムエラーコードを表します
type ErrorCode string

const (
	// エラーコードの定義
	// バリデーションエラーなど
	BadRequest ErrorCode = "BAD_REQUEST"
	// オブジェクトが存在しない
	NotFound ErrorCode = "NOT_FOUND"
	// 認証エラー
	Unauthorized ErrorCode = "UNAUTHORIZED"
	// 想定外のエラーなど
	Internal ErrorCode = "INTERNAL"
)

// CustomError はカスタムエラーの構造体です
type CustomError struct {
	Code    ErrorCode
	Message string
	Values  map[string]any
}

// New は新しいカスタムエラーを作成します
func New(code ErrorCode, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

// Error は error インターフェースの実装です
func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// エラーと一緒に返したい値
func (e *CustomError) WithValues(values map[string]any) {
	e.Values = values
}

// ToGraphQLError はカスタムエラーを GraphQL のエラー形式に変換します
func (e *CustomError) ToGraphQLError() *gqlerror.Error {
	return &gqlerror.Error{
		Message: e.Message,
		Extensions: map[string]interface{}{
			"code": e.Code,
		},
	}
}
