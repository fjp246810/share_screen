# Share Screen

Go语言简单实现屏幕分享功能 

# 技术点

- 本地开启Http Web Server
- 实时抓取屏幕图片
- 采用MJPEG传输屏幕图片
  
# 操作步骤

1 在项目根目录下执行 `go build` 生成可执行文件

2 执行生成的执行文件（windows生成的exe文件，需要以管理员的方式执行）

3 在浏览器中访问执行文件的`http://ip:8080`
