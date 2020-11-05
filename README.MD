使用gin和gorm开发的问答系统

项目代码清晰，非常适合初学者学习，本代码已经被我放入我的新书《Go微服务实战》中，作为基础篇的一个例子，是Go语言Web开发一个比较好的例子，欢迎大家Star！

##项目的结构图
![image](https://github.com/ScottAI/questions/blob/master/pictures/architecture.png)
##Go Web项目常用的一些分包：
####	--Project
####	----api
####	----cmd
####	----internal
####	----pkg
####	----configs
####	----build
####	----docs
####	----examples
####	----assets
####	----templates

## 包介绍
	大家在阅读很多开源项目（github上有很多）时，会发现很多项目都有上述的一些包，不过不会完全一样，有共性的基本时上面列出来的这些。我们在案例的实现时，也会参考这些布局方式，也会灵活的按照需求添加一些新的分包方式。
	先来介绍一下上面各个目录的主要意义：
### 	api：读者可以参考Go语言著名开源项目kubernetes下的api路径，主要放Openapi/Swagger的spec、JSON的protocol或schema定义。
### 	cmd：如果我们开发的不是一个简单函数包，而是一个比较完整和复杂的应用程序，那么可以把main函数放到cmd下。如果有多个main的话可以在cmd下再分成几个子目录。不过要注意，不要在cmd下放太多代码，主要业务代码可以放到pkg或internal下。
### 	internal：私有的内部包，只能被包内或者其直接父级目录引用，不能被外部引用。这样可以防止接口的无序扩散。
### 	pkg：可以被项目内外引用的代码库。
### 	configs：配置文件及配置信息解析代码。
### 	build：打包和ci相关的文件，在下一篇会介绍的docker等都放在这里。
### 	docs：设计文档、用户文档的存放处。
### 	examples：示例代码存放处。
### 	assets：css、图片等web用到的资源存放处。
### 	templates：页面模板存放处。

## 最终运行效果

![](https://github.com/ScottAI/questions/blob/master/pictures/view.png)

##运行
git clone
