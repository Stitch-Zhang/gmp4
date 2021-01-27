# 说明
## gmp4
gmp4是一个能够通过视频的url\本地路径获取其时长的轻量级库，目前仅支持MP4格式的视频
## 安装
```
go get -u github.com/Stitch-Zhang/gmp4
```
## 使用
### 获取网络资源中MP4视频时长
*https://www.runoob.com/try/demo_source/movie.mp4*

#### Example
```go
package main

import (
	"fmt"
	"github.com/Stitch-Zhang/gmp4"
)

func main() {
	video, err := gmp4.NewRemote("https://www.runoob.com/try/demo_source/movie.mp4")
	if err != nil {
		fmt.Println("获取失败:", err)
		return
	}
	videoTime := gmp4.GetDuration(video)
	fmt.Printf("获取成功，时长=%ds", videoTime)
}

///获取成功，时长=:12s
```
#### 设置请求头
```go
···
    video, err := gmp4.NewRemote("https://www.runoob.com/try/demo_source/movie.mp4")
	if err != nil {
		fmt.Println("获取失败:", err)
		return
    }
    header :=http.Header{
	    "User-Agent": {"Mozilla/5.0 (Windows NT 10.0; x64; rv:85.0) Gecko/20100101 Firefox/85.0"},
		///Key,Value 有两种方式存放
		///方法1
		"Cookie": {"cookie1_key=cookie1_value","cookie2_key=cookie2_value"},
		///方法2
		"Cookie": {"cookie1_key=cookie1_value;cookie2_key=cookie2_value"}
	}
	video.SetHeader(header)
···
```
#### 设置请求超时
```go    
···
    video, err := gmp4.NewRemote("https://www.runoob.com/try/demo_source/movie.mp4")
	if err != nil {
		fmt.Println("获取失败:", err)
		return
    }
    video.SetTimeout(3*time.Second)

···
```
### 获取本地MP4视频时长
文件路径:C:\demo.mp4
```go
package main

import (
	"fmt"
	"github.com/Stitch-Zhang/gmp4"
)

func main() {
	video, err := gmp4.NewLocal(`C:\demo.mp4`)
	if err != nil {
		fmt.Println("获取失败:", err)
		return
	}
	videoTime := gmp4.GetDuration(video)
	fmt.Printf("获取成功，时长=%ds", videoTime)
}

///获取成功，时长=:2233s
```
#### 🎉🎉🎉
如有Bug或疑难，欢迎提交PR

# Readme
## gmp4
## Install
```
go get -u github.com/Stitch-Zhang/gmp4
```
## Usage
### Get a web video duration with url 
*https://www.runoob.com/try/demo_source/movie.mp4*

#### Example
```go
package main

import (
	"fmt"
	"github.com/Stitch-Zhang/gmp4"
)

func main() {
	video, err := gmp4.NewRemote("https://www.runoob.com/try/demo_source/movie.mp4")
	if err != nil {
		fmt.Println("Failed to parse", err)
		return
	}
	videoTime := gmp4.GetDuration(video)
	fmt.Printf("Succeess!，Duration=%ds", videoTime)
}

///Succeess!，Duration=:2233s
```
#### Set request header
```go
···
    video, err := gmp4.NewRemote("https://www.runoob.com/try/demo_source/movie.mp4")
	if err != nil {
		fmt.Println(":", err)
		return
    }
    header :=http.Header{
	    "User-Agent": {"Mozilla/5.0 (Windows NT 10.0; x64; rv:85.0) Gecko/20100101 Firefox/85.0"},
		///There are two ways to store key,value
		///Way1
		"Cookie": {"cookie1_key=cookie1_value","cookie2_key=cookie2_value"},
		///Way2
		"Cookie": {"cookie1_key=cookie1_value;cookie2_key=cookie2_value"}
	}
	video.SetHeader(header)
···
```
#### Set request timeout
```go    
···
    video, err := gmp4.NewRemote("https://www.runoob.com/try/demo_source/movie.mp4")
	if err != nil {
		fmt.Println("Failed to parse:", err)
		return
    }
    video.SetTimeout(3*time.Second)

···
```
### Get local file mp4 duration
File path:C:\demo.mp4
```go
package main

import (
	"fmt"
	"github.com/Stitch-Zhang/gmp4"
)

func main() {
	video, err := gmp4.NewLocal(`C:\demo.mp4`)
	if err != nil {
		fmt.Println("Failed to parse:", err)
		return
	}
	videoTime := gmp4.GetDuration(video)
	fmt.Printf("Succeess!,Duration=%ds", videoTime)
}

///Succeess!,Duration=:2233s
```
### 🎉🎉🎉
**It is welcome if you can make this project better in Pull Request**