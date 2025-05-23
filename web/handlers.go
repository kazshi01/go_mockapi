package web

import (
	"encoding/json"
	"fmt"
	"mockapi/service"
	"net/http"
	"strconv"
	"strings"
)

type IUserHandler interface {
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	GetUserByIdHandler(w http.ResponseWriter, r *http.Request)
	HandleRoot(w http.ResponseWriter, r *http.Request)
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	us service.IUseServeice
}

func NewUserHandler(us service.IUseServeice) IUserHandler {
	return &UserHandler{us}
}

// ユーザー一覧を取得
func (uh *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users := uh.us.GetUsers()

	// レスポンスのContent-Typeを設定
	w.Header().Set("Content-Type", "application/json")

	// 実際のusersデータをJSONとしてエンコードして返却
	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		// エンコードエラーが発生した場合のエラーハンドリング
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to encode response"}`))
		fmt.Println("Error encoding JSON:", err)
	}
}

// ユーザーを検索
func (uh *UserHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	// URL パスから ID を抽出 ("/getuser/1" → "1")
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/getuser/")

	// 数値に変換
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// 不正な ID 形式
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid ID format"}`))
		return
	}

	// serviceレイヤーから Userを取得
	user, err := uh.us.GetUserById(id)
	if err != nil {
		// ユーザーが見つからない場合
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "User not found"}`))
		return
	}

	// 成功した場合は User を JSON で返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// ルーティング設定していないパス以外のエラー処理
func (uh *UserHandler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	// URLパスが "/" 以外の場合は 404 を返す
	if r.URL.Path != "/" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Not Found", "message": "The requested resource does not exist"}`))
		return
	}

	// ルートパス("/")へのアクセスの場合は別の処理
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Welcome to the API", "endpoints": "Please access '/getusers'"}`))
}

func (uh *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// POSTメソッドのみを許可
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not allowed"}`))
		return
	}

	// リクエストボディを格納する構造体
	var requestBody struct {
		Username string `json:"username"`
		Password int    `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request body"}`))
		return
	}

	err := uh.us.Register(requestBody.Username, requestBody.Password)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Registration successful"}`))
}

func (uh *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// POSTメソッドのみを許可
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not allowed"}`))
		return
	}

	// リクエストボディを格納する構造体
	var requestBody struct {
		Username string `json:"username"`
		Password int    `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request body"}`))
		return
	}

	err := uh.us.Login(requestBody.Username, requestBody.Password)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Login successful"}`))
}
