package main

import (
	_ "github.com/PhysPeach/Tornament/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		// 開発環境での設定
		// ディレクトリにアクセスしたときファイルのリストが見えるようにする
		beego.BConfig.WebConfig.DirectoryIndex = true

		// 生成されたドキュメントにアクセスできるようにする
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "./swagger"
	}

	beego.Run()
}
