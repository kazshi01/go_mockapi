package router

import (
	"fmt"
	"mockapi/web"
	"net/http"
)

type UserRouter struct {
	uh web.IUserHandler
}

func NewUserRouter(uh web.IUserHandler) *UserRouter {
	return &UserRouter{uh}
}

// ルーティングしていないパス以外のエラー処理
func (rt *UserRouter) HandleRoot(w http.ResponseWriter, r *http.Request) {
	// URLパスが "/" 以外の場合は 404 を返す
	if r.URL.Path != "/" {
		// JSONでメッセージを返す設定
		w.Header().Set("Content-Type", "application/json")
		// ステータスコードを404に設定
		w.WriteHeader(http.StatusNotFound)
		// エラーメッセージを返却
		w.Write([]byte(`{"error": "Not Found", "message": "The requested resource does not exist"}`))
		return
	}

	// ルートパス("/")へのアクセスの場合は別の処理
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Welcome to the API", "endpoints": "/sayにアクセスしてください"}`))
}

func (rt *UserRouter) SetupRoutes() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/say", rt.uh.SayHandler)

	mux.HandleFunc("/", rt.HandleRoot)

	fmt.Println("Server Running")

	// サーバーの設定
	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return server
}
