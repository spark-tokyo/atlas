package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"atlas/api/app/di"
	"atlas/utils"
)

/*
init関数は特殊な関数で、パッケージの初期化に使われます。
mainパッケージに書くとmain関数より先に実行されます。
*/
func init() {
	log.Println("start server")
}

func main() {
	app, clean, err := di.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	defer clean()
	server := &http.Server{
		// ポート8080で受付
		Addr: "0.0.0.0:8080",
		Handler: h2c.NewHandler(
			app.Router.Mux,
			&http2.Server{
				IdleTimeout: 120 * time.Second,
			},
		),
		ReadTimeout:       0,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	// チャンネルの作成 idleConnsClosed は、シャットダウンプロセスが完了したことを通知するためのチャネルとして使用されます。
	idleConnsClosed := make(chan struct{})
	ctx := context.Background()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		log.Println("start server")
		<-sigCh

		shoutDownCtx, cancel := context.WithTimeout(ctx, utils.RequestTimeout+10*time.Second)
		//  Canceledは、コンテキストがキャンセルされたときに[Context.Err]が返すエラーです。
		// var Canceled = errors.New("context canceled")
		defer cancel()
		err := server.Shutdown(shoutDownCtx)
		if err != nil {
			log.Fatal(errors.New("server is already listening"))
		}
		// close組み込み関数は、双方向か送信専用でなければならないチャンネルを閉じる。
		close(idleConnsClosed)
	}()

	// サーバーが起動し、リクエストを待ち受けます。
	// ListenAndServe が http.ErrServerClosed エラー以外のエラーを返した場合、
	// プログラムはエラーとしてログに記録して終了します。
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("%+v\n", err)
		}
	}
	// idleConnsClosed チャンネルから値を受け取るまでブロックされます。このため、シャットダウンが完了するまでプログラムは終了しません。
	<-idleConnsClosed
	fmt.Printf("end")
}

// todo: wire=done、main=done、docker=done、config=実装、db=done、transaction=未実装, error=未実装
