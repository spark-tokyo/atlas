package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/rs/cors"

	"github.com/spark-tokyo/atlas/api/resolver"
	"github.com/spark-tokyo/atlas/config"
	graphqlgen "github.com/spark-tokyo/atlas/graphql/generate"
)

type Router struct {
	Mux *chi.Mux
}

const (
	requestTimeOut = 240 * time.Second
)

func NewRouter(
	resolve *resolver.Resolver,
	config *config.Config,
) *Router {
	/* GraphQL のハンドリング設定 */

	// resolverの読み込み
	log.Println("GraphQL Resolver読み込み中")
	graphqlSchema := graphqlgen.NewExecutableSchema(
		graphqlgen.Config{
			Resolvers: resolve,
		},
	)

	// ハンドリング設定
	log.Println("ハンドリング設定追加中")
	srv := handler.New(graphqlSchema)

	// 通信時間をmax10秒に指定
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})

	// httpメソッドを追加できるように設定
	srv.AddTransport(transport.Options{})
	// GETを受付を可能に
	srv.AddTransport(transport.GET{})
	// POSTの受付を可能に
	srv.AddTransport(transport.POST{})
	// 大きなファイルのアップロードを可能にする設定 ⇨ https://qiita.com/s_yasunaga/items/61bcbae2aec5b03b6543
	srv.AddTransport(transport.MultipartForm{
		MaxUploadSize: 3 << 30, // 3gb
		MaxMemory:     3 << 30, // 3gb
	})

	// クエリのキャッシュ時間を指定
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	// serverにキャッシュがあれば、それを返却するように
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	/* webサーバーの構築 */
	log.Println("webサーバー構築中")
	mux := chi.NewRouter()

	// corsの設定
	mux.Use(cors.New(cors.Options{
		// オリジンからのリクエストを許可 ⇨ 前段のCDNでフィルタするのでserverでは全て許可する
		AllowedOrigins: []string{"*"},
		// すべてのヘッダーがリクエストで許可
		AllowedHeaders: []string{"*"},
		// 認証情報を含むリクエストが許可
		AllowCredentials: true,
	}).Handler)

	mux.Use(
		// ctxのセット
		SetContext(),
		// リクエストのタイムアウトを設定
		chiMiddleware.Timeout(requestTimeOut),
	)

	if IsOpenPlayGround(*config) {
		// プレイグラウンドの設定
		mux.Handle("/", playground.Handler("GraphQL PlayGround", "/query"))
		mux.Handle("/altair", playground.AltairHandler("GraphQL PlayGround", "/query"))
		mux.Handle("/apollo", playground.ApolloSandboxHandler("GraphQL PlayGround", "/query"))
		log.Println("ローカル: URL=http://127.0.0.1:8080/ && GraphQL PlayGround=/ or /altair or /apollo")
	} else {
		// プレイグランドにアクセスできないときは、スキーマも見れないようにする
		srv.AroundOperations(DisableIntrospection)
	}

	// アクセス
	mux.Handle("/query", srv)

	return &Router{
		Mux: mux,
	}
}

// リクエストした時間
type typeCtxClockKey struct{}

var ctxClockKey typeCtxClockKey

// リクエストID
type typeCtxRequestIdKey struct{}

var ctxRequestIdkey typeCtxRequestIdKey

// server起動時に以降の処理で必要な値を ctx に渡す
func SetContext() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			id, err := uuid.NewRandom()
			if err != nil {
				panic(err)
			}
			// 一意のリクエストIDを追加
			// ctx に渡すときは key をオリジナルの型を定義する必要がある
			ctx = context.WithValue(ctx, ctxRequestIdkey, id)

			// server内で同じ時間を使うため、middlewate 内で ctx に現在時刻を渡す
			ctx = context.WithValue(ctx, ctxClockKey, time.Now())
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// ステージ確認
func IsOpenPlayGround(stage config.Config) bool {
	return stage.IsLocal() || stage.IsDev()
}

// 無効化設定
func DisableIntrospection(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	graphql.GetOperationContext(ctx).DisableIntrospection = true
	return next(ctx)
}
