
#选型
目前主要的go web框架主要有revel、beego、martini等。
revel比较老
beego大而全
martini小而精巧
目前暂时采用martini。
#扩展
应对http请求，只需要martini就可以了
要支持渲染模版，需要render扩展
要支持session，需要sessions扩展
要支持参数绑定，需要binding扩展
要支持数据压缩，需要gzip扩展

另外，还有节流阀、权限等扩展。
#系统设计
layout.html 公共结构
index.html  首页：左右介绍，右边登录
welcome.html欢迎页面：右上角有当前用户信息，退出按钮
