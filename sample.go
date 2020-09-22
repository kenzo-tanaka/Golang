// 問題: https://github.com/gopherdojo/dojo1/tree/master/kadai1

package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// @see https://news.mynavi.jp/article/gogogo-5/
// @see https://shiro-16.hatenablog.com/entry/2020/05/29/130508
func main() {
	flag.Parse()
	args := flag.Args()

	f, err := os.Open(args[0]) // 元画像読み込み
	if err != nil {
		fmt.Println("open:", err)
		return
	}
	defer f.Close()

	img, _, err := image.Decode(f) // 元画像デコード
	if err != nil {
		fmt.Println("decode:", err)
		return
	}

	fso, err := os.Create(args[1]) // 変換後画像作成
	if err != nil {
		fmt.Println("create:", err)
		return
	}
	defer fso.Close()

	slice := strings.Split(args[1], ".")

	switch slice[len(slice)-1] { // 出力画像の拡張子によってエンコードを変える
	case "jpeg", "jpg":
		jpeg.Encode(fso, img, &jpeg.Options{})
	case "png":
		png.Encode(fso, img)
	case "gif":
		gif.Encode(fso, img, nil)
	default:
	}
}
