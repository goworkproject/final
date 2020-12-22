# 服务计算第九次作业
### 小组分工：

+ 伍德翔：大部分后台，一部分前端的demo

+ 王裕文：部分后台，包括注册、登录、退出，大部分前端

### 下载文件：

```
go get https://github.com/goworkproject/final.git
```

### 博客内容：

>  使用语言:go
>
> 使用goWeb框架:gin 学习网站：https://www.bookstack.cn/read/gin-doc/controller.md
>
> 使用的前端框架: vue.js 
>
> 使用的前端UI框架: element UI
>
> Ubuntu下安装环境的教程，以及遇到的难题，请参考博客

 

+ 快速下载Golang框架gin

  https://www.cnblogs.com/woodx/p/14130891.html

+ 用官网包安装mongodb

  https://www.cnblogs.com/woodx/p/14134885.html

+ Go下Mongodb操作入门

  https://www.cnblogs.com/woodx/p/14135703.html

+ Ubuntu下Mongodb操作入门

  https://www.cnblogs.com/woodx/p/14140943.html

+ 使用Vue.js播放视频

  https://www.cnblogs.com/woodx/p/14166183.html 

+ gin获取post的body里的参数

  https://i.cnblogs.com/posts/edit-done;postId=14176115

#### API 接口语言(JavaScript)

#### 登录

```go
this.$http.post('http://172.26.65.118:8080/api/signIn',{ 

        nickname: this.name, 

        password: this.password, 

  

      },{ 

       // emulateJSON: true 

       }).then(function(res){ 

        console.log(res); 

        //tiaozhuan  

  

        var url = "http://172.26.65.118:8080/html/index.html?id=" + res.body.data; 

        window.location.href = url; 
 });
```



##### 注册

```go
this.$http.post('http://172.26.65.118:8080/api/signUp',{ 

      nickname: this.rname, 

      password:  this.rpassword, 

       email: this.email, 

       }, 

       { 

       // emulateJSON: true 

       }).then(function(res){ 

        console.log(res) 

 }); 
```

退出API （服务器端实现，未写出）

 

##### 获取所有视频API

```go
 this.$http.get('http://172.26.65.118:8080/api/getVideoInfo',{ 

      params: { 

        

      }, 

     },{ 

      // emulateJSON: true 

      }).then(function(res){ 

      this.TVinfos = res.body.data; 

       this.show_TVinfos = this.TVinfos; 

       console.log(*this.TVinfos); 

     } 

 ); 
```



##### 根据用户id获取个人信息

```go
this.$http.get('http://172.26.65.118:8080/api/getSingleUserInfo',{ 

       params: { 

        tvid: newUserId 

      }, 

       },{ 

        // emulateJSON: true 

       }).then(function(res){ 

        this.userInfo = res.body.data; 

        console.log(*this.userInfo); 

      }); 
```

##### 根据视频id获取视频信息

```go
this.$http.get('http://172.26.65.118:8080/api/getSingleTVInfo',{ 

       params: { 

        tvid: newsid 

       }, 

       },{ 

        // emulateJSON: true 

       }).then(function(res){ 

        this.TVInfo = res.body.data; 

        this.description = *this.TVInfo.Name; 

        // console.log(this.TVInfo); 

 }); 
```



##### 根据视频id获取播放信息

 ```go
this.$http.get('http://172.26.65.118:8080/api/getPlayInfo',{ 

       params: { 

        tvid: newsid 

      }, 

       },{ 

        // emulateJSON: true 

       }).then(function(res){ 

        this.PlayInfo = res.body.data; 

        this.show_time = this.PlayInfo.ViewTime; 

       this.$refs.video.src = this.PlayInfo.VideoUrl; 

        this.detail = this.PlayInfo.Brief; 

        if (this.PlayInfo.RatePeopleNum == 0) { 

         this.value2 = 5; 

        } else { 

         this.value2 = (this.PlayInfo.AllRateNum / this.PlayInfo.RatePeopleNum); 

        } 
      }); 
 ```

##### 根据视频id获取留言信息

```go
this.$http.get('http://172.26.65.118:8080/api/getMessage',{ 

       params: { 

        tvid: newsid 

       }, 

       },{ 

        // emulateJSON: true 

       }).then(function(res){ 

        this.Messages = res.body.data.Messages; 

    	console.log(this.Messages); 

      }); 
```

##### 根据视频id发表评论

```go
this.$http.get('http://172.26.65.118:8080/api/postMessage',{ 

      params: { 

       tvid: newsid, 

       sender: this.userInfo.Nickname, 

       receiver: "", 

       content: this.chat_message 

      }, 

     },{ 

       // emulateJSON: true 

      }).then(function(res){ 

        

     }); 
```

##### 根据视频id给视频评分

```go
this.$http.post('http://172.26.65.118:8080/api/thumbUp',{ 

       tvid: this.PlayInfo.TVID, 

       rate: ""+this.value2, 

       },{ 

        // emulateJSON: true 

       }).then(function(res){ 
		console.log(res);  
	}); 
```



