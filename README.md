# goenv
将 package 安装到当前目录下，并设置 GOPATH 为当前目录

# 安装
go get github.com/freeonsky/goenv

安装好之后，先运行 goenv 程序，当提示是否将 goenv 拷贝到 $GOROOT/bin 时，请输入y，这样就会把goenv 程序拷贝到$GOROOT/bin，一般都会将$GOROOT/bin添加到PATH环境变量，这样能够保证在cmd/powershell下执行 goenv 命令

# 使用
执行 `goenv get package` ,会将 package 安装到当前目录下，并设置 GOPATH 为当前目录
例如安装gocode，执行如下命令：
`goenv get get github.com/nsf/gocode`

# vscode中有很多go的相关插件，非常好用如下：
gocode
gopkgs
go-outline
go-symbols
guru
gorename
gomodifytags
goplay
impl
godef
goreturns
golint
gotests
dlv