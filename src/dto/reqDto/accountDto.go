package reqDto

type AccountLogin struct {
	Phone    string `json:"phone" ` //binding:"len=12"
	UserName string `json:"user_name"`
	Password string `json:"password" binding:"required"`
	Method   string `json:"method" validate:"oneof=name phone,required"  example:"name"`
	Revoke   bool   `json:"revoke" example:"false"` //binding:"required"
	Code     string `json:"code" binding:"required"`
	Uuid     string `json:"uuid" binding:"required"`
}
type UpdateAccount struct {
	Id       uint   `json:"id" binding:"required" example:"2"`
	UserName string `json:"user_name" `
	NickName string `json:"nick_name" `
	Phone    string `json:"phone" `
	Avatar   string `json:"avatar" `
	Email    string `json:"email" `
	Role     []int  `json:"role" `
	Password string `json:"password"  example:"123456"`
	Pwd2     string `json:"pwd2" validate:"eqfield=Password"`
	Class    string `json:"class"`
}
type AddAccount struct {
	UserName string `json:"user_name" binding:""`
	NickName string `json:"nick_name" binding:""`
	Password string `json:"password" binding:"len=6,omitempty"`
	Phone    string `json:"phone" ` //binding:"len=12,omitempty,-"
	Sex      int    `json:"sex" `   //binding:"len=1,omitempty,-"
	Avatar   string `json:"avatar" binding:""`
	Email    string `json:"email" ` //binding:"email,omitempty,-"
	Role     []int  `json:"role" binding:""`
	Class    string `json:"class"`
}
type AccountList struct {
	Take int    `json:"take,omitempty" binding:"required"`
	Skip uint   `json:"skip,omitempty"`
	Name string `json:"name,omitempty"`
}
