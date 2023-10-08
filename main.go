package main

import (
	"fmt"
	"ry_go/src/global"
	"ry_go/src/route"
)

func main() {
	r := route.InitRoute()
	//fmt.Println("路由挂载在main程序")
	if err := r.Run(global.Port); err != nil {
		//panic(err)
		fmt.Println(fmt.Errorf("端口占用,err:%v\n", err))
	} //, "https/certificate.crt", "https/private.key"); err != nil {
	//	panic(err)
	//}
	art :=
		`
   ###      ##     #          ##     #    #   ####               ####    #    #     ###      ###    ######    ####     ####
  #   #    #  #    #         #  #    ##   #    #  #             #    #   #    #    #   #    #   #   #        #    #   #    #
 #        #    #   #        #    #   # #  #    #   #            #        #    #   #        #        #        #        #
 #  ###   #    #   #        ######   #  # #    #   #             ####    #    #   #        #        ####      ####     ####
 #    #   #    #   #        #    #   #   ##    #   #                 #   #    #   #        #        #             #        #
  #   #    #  #    #        #    #   #    #    #  #             #    #   #    #    #   #    #   #   #        #    #   #    #
   ###      ##     ######   #    #   #    #   ####               ####     ####      ###      ###    ######    ####     ####
`
	//	fmt.Println(`
	//   ###      ##     #          ##     #    #   ####               ####    #    #     ###      ###    ######    ####     ####
	//  #   #    #  #    #         #  #    ##   #    #  #             #    #   #    #    #   #    #   #   #        #    #   #    #
	// #        #    #   #        #    #   # #  #    #   #            #        #    #   #        #        #        #        #
	// #  ###   #    #   #        ######   #  # #    #   #             ####    #    #   #        #        ####      ####     ####
	// #    #   #    #   #        #    #   #   ##    #   #                 #   #    #   #        #        #             #        #
	//  #   #    #  #    #        #    #   #    #    #  #             #    #   #    #    #   #    #   #   #        #    #   #    #
	//   ###      ##     ######   #    #   #    #   ####               ####     ####      ###      ###    ######    ####     ####
	//`)
	fmt.Println(art)

	//fmt.Println("server at ", &global.Port)
	//addrs, err := net.InterfaceAddrs() //Dial("udp", "8.8.8.8:80")
	//if err != nil {
	//	fmt.Println("获取id错误", err)
	//}
	//var ip string
	//for _, addr := range addrs {
	//	// 检查IP地址的类型，排除IPv6地址和回环地址
	//	if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//		if ipnet.IP.To4() != nil {
	//			ip = ipnet.IP.String()
	//			fmt.Println("IPv4 Address:", ipnet.IP.String())
	//		}
	//	}
	//}
	//ip := conn.LocalAddr().(*net.UDPAddr).IP.String()
	//fmt.Println("获取id", )
	//fmt.Println("serve run at %s", global.Port)
	//fmt.Printf("serve run at %s%s", ip, global.Port)
}
