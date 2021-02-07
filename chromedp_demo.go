package main

//https://studygolang.com/topics/12596
import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	goQrcode "github.com/skip2/go-qrcode"
)

const loginURL = "https://account.wps.cn/"
// 常用选择器
//chromedp.BySearch // 如果不写，默认会使用这个选择器，类似devtools ctrl+f 搜索
//chromedp.ByID // 只id来选择元素
//chromedp.ByQuery // 根据document.querySelector的规则选择元素，返回单个节点
//chromedp.ByQueryAll // 根据document.querySelectorAll返回所有匹配的节点
//chromedp.ByNodeIP // 检索特定节点(必须先有分配的节点IP)，这个暂时没用过也没看到过例子，如果有例子可以发给我看下

// 常用api
//chromedp.Navigate("https://xxxx") // 设置url
//chromedp.WaitVisible(`#username`, chromedp.ByID), //  使用chromedp.ByID选择器。所以就是等待id=username的标签元素加载完。
//chromedp.SendKeys(`#username`, "username", chromedp.ByID), // 使用chromedp.ByID选择器。向id=username的标签输入username。
//chromedp.Value(`#input1`, val1, chromedp.ByID), // 获取id=input1的值，并把值传给val1
//chromedp.Click("btn-submit",chromedp.Bysearch), // 触发点击事件，
//chromedp.Screenshot(`#row`, &buf, chromedp.ByID), // 截图id=row的标签，把值传入buf 需要事先定义var buf []byte
//chromedp.ActionFunc(func(context.Context, cdp.Executor) error { // 将图片写入文件
//	return ioutil.WriteFile("test.png", buf, 0644)
//}),

func main() {
	// chromdp依赖context上限传递参数
	ctx, _ := chromedp.NewExecAllocator(context.Background(),
		// 以默认配置的数组为基础，覆写headless参数
		// 当然也可以根据自己的需要进行修改，这个flag是浏览器的设置
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
		)...,
	)

	// 创建新的chromedp上下文对象，超时时间的设置不分先后
	// 注意第二个返回的参数是cancel()，只是我省略了
	ctx, _ = context.WithTimeout(ctx, 30*time.Second)
	ctx, _ = chromedp.NewContext(
		ctx,
		// 设置日志方法
		chromedp.WithLogf(log.Printf),
	)
	// 通常可以使用 defer cancel() 去取消
	// 但是在Windows环境下，我们希望程序能顺带关闭掉浏览器
	// 如果不希望浏览器关闭，使用cancel()方法即可
	// defer cancel()
	// defer chromedp.Cancel(ctx)

	if err := chromedp.Run(ctx, myTasks()); err != nil {
		log.Fatal(err)
		return
	}

}

func myTasks() chromedp.Tasks {
	return chromedp.Tasks{
		// 1. 打开金山文档的登陆界面
		chromedp.Navigate(loginURL),
		// 2. 点击微信登陆按钮
		// #wechat > span:nth-child(2)
		chromedp.Click(`#wechat > span:nth-child(2)`),
		// 判断一下是否已经登陆  <-- 变动
		checkLoginStatus(),
		// 3. 点击确认按钮
		// #dialog > div.dialog-wrapper > div > div.dialog-footer > div.dialog-footer-ok
		chromedp.Click(`#dialog > div.dialog-wrapper > div > div.dialog-footer > div.dialog-footer-ok`),
		// 4. 获取二维码
		// #wximport
		getCode(),

		// 5. 若二维码登录后，浏览器会自动跳转到用户信息页面  <-- 变动
		saveCookies(),
	}

}

func getCode() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// 1. 用于存储图片的字节切片
		var code []byte

		// 2. 截图
		// 注意这里需要注明直接使用ID选择器来获取元素（chromedp.ByID）
		if err = chromedp.Screenshot(`#wximport`, &code, chromedp.ByID).Do(ctx); err != nil {
			return
		}

		//// 3. 保存文件
		//if err = ioutil.WriteFile("code.png", code, 0755); err != nil {
		//	return
		//}
		// 3. 把二维码输出到标准输出流
		if err = printQRCode(code); err != nil {
			return err
		}
		return
	}
}

func printQRCode(code []byte) (err error) {
	// 1. 因为我们的字节流是图像，所以我们需要先解码字节流
	img, _, err := image.Decode(bytes.NewReader(code))
	if err != nil {
		return
	}

	// 2. 然后使用gozxing库解码图片获取二进制位图
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return
	}

	// 3. 用二进制位图解码获取gozxing的二维码对象
	res, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
	if err != nil {
		return
	}
	// 4. 用结果来获取go-qrcode对象（注意这里我用了库的别名）
	qr, err := goQrcode.New(res.String(), goQrcode.High)
	if err != nil {
		return
	}

	// 5. 输出到标准输出流
	fmt.Println(qr.ToSmallString(false))
	time.Sleep(10 * time.Second)
	return
}

// 保存Cookies
func saveCookies() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// 等待二维码登陆
		if err = chromedp.WaitVisible(`#app`, chromedp.ByID).Do(ctx); err != nil {
			return
		}

		// cookies的获取对应是在devTools的network面板中
		// 1. 获取cookies
		cookies, err := network.GetAllCookies().Do(ctx)
		if err != nil {
			return
		}

		// 2. 序列化
		cookiesData, err := network.GetAllCookiesReturns{Cookies: cookies}.MarshalJSON()
		if err != nil {
			return
		}

		// 3. 存储到临时文件
		if err = ioutil.WriteFile("cookies.tmp", cookiesData, 0755); err != nil {
			return
		}
		return
	}
}

// 加载Cookies
func loadCookies() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		// 如果cookies临时文件不存在则直接跳过
		if _, _err := os.Stat("cookies.tmp"); os.IsNotExist(_err) {
			return
		}

		// 如果存在则读取cookies的数据
		cookiesData, err := ioutil.ReadFile("cookies.tmp")
		if err != nil {
			return
		}

		// 反序列化
		cookiesParams := network.SetCookiesParams{}
		if err = cookiesParams.UnmarshalJSON(cookiesData); err != nil {
			return
		}

		// 设置cookies
		return network.SetCookies(cookiesParams.Cookies).Do(ctx)
	}
}

// 检查是否登陆
func checkLoginStatus() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		var url string
		if err = chromedp.Evaluate(`window.location.href`, &url).Do(ctx); err != nil {
			return
		}
		if strings.Contains(url, "https://account.wps.cn/usercenter/apps") {
			log.Println("已经使用cookies登陆")
			chromedp.Stop()
		}
		return
	}
}
