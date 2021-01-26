package gmp4

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_NewRemote(t *testing.T) {
	remote,err :=NewRemote("https://www.runoob.com/try/demo_source/movie.mp4")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("remote demo:%ds\n",GetDuration(remote))
}

func TestRemoteVideo_SetHeader(t *testing.T) {
	remote,err:=NewRemote("https://www.runoob.com/try/demo_source/movie.mp4")
	if err!=nil{
		fmt.Println(err)
		return
	}
	remote.SetHeader(http.Header{
		"User-Agent":{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.19041"},
		"Cookie": {"has_js=1","login_flag=0;", "uc_local_login=2;"},
	})
	fmt.Printf("SetHeader demo:%ds\n",GetDuration(remote))
}

func Test_GetDuration(t *testing.T) {
	remote,err :=NewRemote("https://www.runoob.com/try/demo_source/movie.mp4")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("GetDuration demo:%ds\n",GetDuration(remote))
}


func Test_NewLocal(t *testing.T) {
	local,err:=NewLocal("demo.mp4")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("Local demo:%ds\n",GetDuration(local))
}

