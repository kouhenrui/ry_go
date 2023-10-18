package reqDto

type AdminLogin struct {
	Phone    string `json:"phone" ` //binding:"len=12"
	UserName string `json:"user_name"`
	Password string `json:"password" binding:"required"`
	Method   string `json:"method" binding:"oneof=name phone,required" `
	Revoke   bool   `json:"revoke" ` //binding:"required"
	Code     string `json:"code" binding:"required,len=6"`
	Uuid     string `json:"uuid" binding:"required"`
}
type UpdateAdmin struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
type AddAdmin struct {
	UserName string `json:"user_name" binding:""`
	NickName string `json:"nick_name" binding:""`
	Password string `json:"password" binding:"len=6,omitempty"`
	Phone    string `json:"phone" ` //binding:"len=12,omitempty,-"
	Sex      int    `json:"sex" `   //binding:"len=1,omitempty,-"
	Avatar   string `json:"avatar" binding:""`
	Email    string `json:"email" ` //binding:"email,omitempty,-"
	Role     []int  `json:"role" binding:""`
}
type AdminList struct {
	Take int    `json:"take,omitempty" binding:"required"`
	Skip uint   `json:"skip,omitempty"`
	Name string `json:"name,omitempty"`
}
