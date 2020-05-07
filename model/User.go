package model

/*User refere-se a abstração de um usuário com acesso ao sistema*/
type User struct {
	ID            uint64 `json:"id"`
	NickName      string `json:"nick_name"`
	NumberOfPosts uint64 `json:"post_amount"`
	Avatar        string `json:"avatar"`
	Role          string `json:"role"`
	Mask          string `json:"mask"`
	Signature     string `json:"signature"`
	Password      string `json:"pwd"`
}
