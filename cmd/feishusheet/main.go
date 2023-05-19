// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kubecub/feishu-sheet-parser/internal/controllers"
)

const helpTips = `欢迎使用分享到飞书表格
提示：本程序的所有的数据都存储在内容中。
访问“http://IP或域名:8000”检测是否开启成功。

使用:
在发送通知之前请设置飞书配置。
1.设置飞书配置：
- 访问“http://你的IP或域名/set/表格Id/应用Id/应用Secret”。
2.获取表格工作表Id：
- 访问“http://你的IP或域名/sheets/密码”。
3.发送通知：
- 访问“http://你的IP或域名/send/密码”。

进阶:
1.达到设置存储的通知条数后再同步的飞书表格：
+ 当内存中存储了10条通知就会将这10条通知一次性发送到飞书表格。如果没有就继续等待满足了后再发送。
- 访问“http://你的IP或域名/sends/工作表Id/密码”。

QA:
1.密码有什么用？
- 密码用于防止其他用户恶意插入数据或者造成其他损失。
2.如何获取密码？
- 访问“http://你的IP或域名/set/表格Id/应用Id/应用Secret”地址会返回128位随机字符串。
3.忘记密码该怎么办？
- 重新启动该服务。再次访问“http://你的IP或域名/set/表格Id/应用Id/应用Secret”。
4.我想修改飞书配置该如何做？
- 访问“http://你的IP或域名/update/表格Id/应用Id/应用Secret/密码”。`

func main() {
	// 配置飞书
	http.HandleFunc("/set/", controllers.SetFeiShuConfig)
	http.HandleFunc("/update/", controllers.UpdateFeiShuConfig)

	// 表格
	http.HandleFunc("/sheets/", controllers.QuerySheetsInfo)

	// 发送通知
	http.HandleFunc("/send/", controllers.SendDataHandle)
	http.HandleFunc("/sends/", controllers.SendDatasHandle)

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, helpTips)
	}
	http.HandleFunc("/", helloHandler)

	fmt.Println(helpTips)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
