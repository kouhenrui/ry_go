package reqDto

type AdminLogin struct {
	Phone    string `json:"phone"`
	UserName string `json:"user_name"`
	Password string `json:"password" binding:"required"`
	Method   string `json:"method" binding:"required,oneof=user_name phone" `
	Revoke   bool   `json:"revoke" validate:"required"`
}
type UpdateAdmin struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
type AddAdmin struct {
	Name     string `json:"name,omitempty"`
	Account  string `json:"account"  binding:"required" validate:"required"`
	Password string `json:"password,omitempty"`
	Role     int    `json:"role"`
	Salt     string `json:"salt,omitempty"`
}
type AdminList struct {
	Take int    `json:"take,omitempty" binding:"required"`
	Skip uint   `json:"skip,omitempty"`
	Name string `json:"name,omitempty"`
}
