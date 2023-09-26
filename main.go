package main

import (
	"fmt"
	"ry_go/src/global"
	"ry_go/src/route"
)

func main() {
	r := route.InitRoute()
	if err := r.Run(global.Port); err != nil {
		//panic(err)
		fmt.Println(fmt.Errorf("端口占用,err:%v\n", err))
	} //, "https/certificate.crt", "https/private.key"); err != nil {
	//	panic(err)
	//}
	fmt.Printf("serve run at %s", global.Port)
}
