# goenv
设置GOPATH，一键安装go-tools

# 安装
go get github.com/freeonsky/goenv

安装好之后，先运行 goenv 程序，当其实是否将 goenv 拷贝到 $GOROOT/bin 时，请输入y，这样就能在cmd/powershell下执行 goenv

# 使用
goenv get url
例如：
goenv get get github.com/nsf/gocode
会将 gocode 安装到当前目录下，并设置 GOPATH 为当前目录