package model

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 戻り値に使用する型を定義
type AllUsers []*User

// パッケージレベルの変数は var を使用する必要がある（メソッドの中ならvar不要）
var Users = []*User{
	{Id: 1, Name: "shiga"},
	{Id: 2, Name: "arisa"},
}
