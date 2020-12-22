package main

import (
	"github.com/gin-gonic/gin"
    "fmt"
    "log"
    "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"path"
	"time"
	"encoding/json"
	"strconv"
)

//用户个人信息
type UserInfo struct {
	ID	bson.ObjectId `bson:"_id,omitempty"` //类型是bson.ObjectId
   	Nickname string
	Password string
	Email string
}

//TV信息，生成的ID主键与PlayInfo、LeavingMessage的ID主键是一样的
type TV struct {
	TVID	bson.ObjectId `bson:"_id,omitempty"` //类型是bson.ObjectId
	Name string
	ImageUrl string
}

//播放页面信息
type PlayInfo struct {
	TVID	bson.ObjectId `bson:"_id,omitempty"` //类型是bson.ObjectId
	UpName	string
	VideoUrl string //shi pin dizhi
	Brief string 	//jian jie
	AllRateNum int	//RateNum
	RatePeopleNum int 	//RatePeopleRateNum
	ViewTime int	//guankan cishu
	SetNum int		//duoshao ji
}

//留言的单元，储存留言的信息
type Message struct {
	Sender string
	Receiver string
	Content string
	Time time.Time
}

//TV的留言
type LeavingMessage struct {
	TVID	bson.ObjectId `bson:"_id,omitempty"` //类型是bson.ObjectId
	Messages []Message
}
 

var signInUserMap map[string]UserInfo

func main() {
	signInUserMap = make(map[string]UserInfo)

	router := gin.Default()

	//可本地可远程，不指定协议时默认为http协议访问，此时需要设置 mongodb 的nohttpinterface=false来打开httpinterface。
	//也可以指定mongodb协议，如 "mongodb://127.0.0.1:27017"
	var MOGODB_URI = "127.0.0.1:27017"
	//连接
	session, err := mgo.Dial(MOGODB_URI)
	//连接失败时终止
	if err != nil {
		panic(err)
	}
	//延迟关闭，释放资源
	defer session.Close()
	//设置模式
	session.SetMode(mgo.Monotonic, true)
	//选择数据库与集合
	c := session.DB("meijuAPI").C("userInformation")
	

	//插入文档
	
	if err != nil {
		log.Fatal(err)
	}

	router.Static("/html", "./html")

	router.POST("/uploadfiles", func(context *gin.Context) {
		t_name := context.PostForm("name")
		t_brief := context.PostForm("brief")
		// t_id := context.PostForm("id")
		fmt.Println(t_name)
		fmt.Println(t_brief)

		//can't repeat name 
		//查询文档
		result := TV{}
		//注意mongodb存储后的字段大小写问题
		tVInfo := session.DB("meijuAPI").C("TVInfo")
		iter := tVInfo.Find(bson.M{"name": t_name}).Iter()

		nickNameExit := false
		for iter.Next(&result) {
			nickNameExit = true
		}

		if (nickNameExit) {
			context.JSON(http.StatusOK, gin.H{
				"status": "TV name is exited",
			})
			return 
		} 

		form, err := context.MultipartForm()
		images := form.File["image"]
		//错误处理
		if err != nil {
		    context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		    })
		    return
		}
		for _, f := range images{
			filePath :=path.Join("./image/", t_name + ".jpg")
		    context.SaveUploadedFile( f, filePath)
		}

		videos := form.File["video"]
		//错误处理
		if err != nil {
		    context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		    })
		    return
		}
		for _, f := range videos{
				filePath :=path.Join("./video/", t_name + ".mp4")
		    context.SaveUploadedFile( f, filePath)
		}

		
		playInfo := session.DB("meijuAPI").C("PlayInfo")
		lms := session.DB("meijuAPI").C("LeavingMsgs")
		tvInfo_err := tVInfo.Insert( &TV{Name: t_name, ImageUrl: "http://172.26.65.118:8080/api/getImage/" + t_name + ".jpg"})
		if tvInfo_err != nil {
			log.Fatal(tvInfo_err)
		}

		t_tvInfo := TV{}
		tvInfo_err = tVInfo.Find(bson.M{"name": t_name}).One(&t_tvInfo)
		if tvInfo_err != nil {
			log.Fatal(tvInfo_err)
		}
		playInfo_err := playInfo.Insert( &PlayInfo{TVID: t_tvInfo.TVID, UpName: "woodx",  AllRateNum: 0, RatePeopleNum: 0, VideoUrl: "http://172.26.65.118:8080/api/getVideo/" + t_name + ".mp4", Brief: t_brief, ViewTime: 0, SetNum: 1})
		if playInfo_err != nil {
			log.Fatal( playInfo_err)
		}
		var messages []Message
		lms_err := lms.Insert( &LeavingMessage{TVID: t_tvInfo.TVID, Messages: messages})
		if lms_err != nil {
			log.Fatal(lms_err)
		}

		context.JSON(http.StatusOK, gin.H{
		    "status": "success",
		})
    })

	// 使用 gin.BasicAuth 中间件，设置授权用户
    // authorized := router.Group("/api", gin.BasicAuth(gin.Accounts{
    //     "woodx":    "test",
	// }))
	authorized := router.Group("/api")

	//post message
	authorized.GET("/postMessage", func(context *gin.Context){
		t_tvid, _ := context.GetQuery("tvid")
		t_sender, _ := context.GetQuery("sender")
		t_receiver, _ := context.GetQuery("receiver")
		t_content, _ := context.GetQuery("content")
		fmt.Println("postMessage")
		fmt.Println(t_tvid)
		fmt.Println(t_sender)
		fmt.Println(t_receiver)
		fmt.Println(t_content)
		fmt.Println("postMessage")

		lmsgs_collection := session.DB("meijuAPI").C("LeavingMsgs")
		//查询文档
		result := LeavingMessage{}
		//注意mongodb存储后的字段大小写问题
		err := lmsgs_collection.FindId(bson.ObjectIdHex(t_tvid)).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		result.Messages = append(result.Messages, Message{Sender: t_sender, Receiver: t_receiver, Content: t_content, Time: time.Now()})
		
		lmsgs_collection.UpsertId(bson.ObjectIdHex(t_tvid), bson.M{"$set": bson.M{"messages": result.Messages} })
		
		context.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	//sign up
	authorized.POST("/signUp", func(context *gin.Context){
		buf := make([]byte, 1024)
		n, _ := context.Request.Body.Read(buf)
		
		for i := 0; i < n; i +=1 {
			if(buf[i] ==' ' || buf[i] == '\x00') {
				buf = append(buf[:i], buf[(i + 1):]...)
				n -= 1
			}
		}

		map2 := make(map[string]string)

		err = json.Unmarshal(buf[:n], &map2)
		fmt.Println(map2)

		t_nickName := map2["nickname"]
		t_password := map2["password"]
		t_email := map2["email"]

		//出错判断
		if err != nil {
			log.Fatal(err)
		}

		//查询文档
		result := UserInfo{}
		//注意mongodb存储后的字段大小写问题
		iter := c.Find(bson.M{"nickname": t_nickName}).Iter()

		nickNameExit := false
		for iter.Next(&result) {
			nickNameExit = true
		}

		if (nickNameExit) {
			context.JSON(http.StatusOK, gin.H{"status": "nickName is existed"})
		} else {
			err = c.Insert(&UserInfo{Nickname: t_nickName, Password: t_password, Email: t_email})
			
			context.JSON(http.StatusOK, gin.H{"status": "sign up successfully"})
		}
	})

	//sign in
	authorized.POST("/signIn", func(context *gin.Context){
		buf := make([]byte, 1024)
		n, _ := context.Request.Body.Read(buf)
		
		for i := 0; i < n; i +=1 {
			if(buf[i] ==' ' || buf[i] == '\x00') {
				buf = append(buf[:i], buf[(i + 1):]...)
				n -= 1
			}
		}

		map2 := make(map[string]string)

		err = json.Unmarshal(buf[:n], &map2)
		fmt.Println(map2)

		t_nickName := map2["nickname"]
		t_password := map2["password"]
		//出错判断
		if err != nil {
			log.Fatal(err)
		}

		//查询文档
		result := UserInfo{}
		//注意mongodb存储后的字段大小写问题
		iter := c.Find(bson.M{"nickname": t_nickName}).Iter()

		nickNameExit := false
		for iter.Next(&result) {
			nickNameExit = true
		}

		//sigin success! 
		if (nickNameExit) {
			//
			_, ok := signInUserMap[t_nickName]
			if (ok) {
				context.JSON(http.StatusOK, gin.H{"status": "user has signed in"})
			} else {
				
				if (result.Password == t_password) {
					signInUserMap[t_nickName] = result
					context.JSON(http.StatusOK, gin.H{"status": "sign in success", "data": result.ID})
				} else {
					context.JSON(http.StatusOK, gin.H{"status": "password is wrong"})
				}
			}
		} else {
			context.JSON(http.StatusOK, gin.H{"status": "user didn't exit"})
		}
	})

	//log out
	authorized.POST("/logOut", func(context *gin.Context){
		buf := make([]byte, 1024)
		n, _ := context.Request.Body.Read(buf)
		
		for i := 0; i < n; i +=1 {
			if(buf[i] ==' ' || buf[i] == '\x00') {
				buf = append(buf[:i], buf[(i + 1):]...)
				n -= 1
			}
		}

		map2 := make(map[string]string)

		err = json.Unmarshal(buf[:n], &map2)
		fmt.Println(map2)

		t_nickName := map2["nickname"]

		fmt.Println(signInUserMap[t_nickName])
		_, ok := signInUserMap[t_nickName]
		if (ok) {
			delete(signInUserMap, t_nickName)
			context.JSON(http.StatusOK, gin.H{"status": "log out successfully"})
		} else {
			context.JSON(http.StatusOK, gin.H{"status": "didn't sign in"})
		}
	})

	//dianzan 
	authorized.POST("/thumbUp", func(context *gin.Context){
		buf := make([]byte, 4096)
		n, _ := context.Request.Body.Read(buf)
		fmt.Println(string(buf))
		for i := 0; i < n; i +=1 {
			if(buf[i] ==' ' || buf[i] == '\x00') {
				buf = append(buf[:i], buf[(i + 1):]...)
				n -= 1
			}
		}

		map2 := make(map[string]string)

		err = json.Unmarshal(buf[:n], &map2)
		fmt.Println(map2)

		t_rate := map2["rate"]
		t_tvid := map2["tvid"]

		tvInfo_collection := session.DB("meijuAPI").C("PlayInfo")
		//查询文档
		result := PlayInfo{}
		//注意mongodb存储后的字段大小写问题
		err := tvInfo_collection.FindId(bson.ObjectIdHex(t_tvid)).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		
		rate, _ := strconv.Atoi(t_rate)
		_, err = tvInfo_collection.UpsertId(bson.ObjectIdHex(t_tvid), bson.M{"$set": bson.M{"allratenum": result.AllRateNum + rate} })
		if err != nil {
			log.Fatal(err)
		}
		
		_, err = tvInfo_collection.UpsertId(bson.ObjectIdHex(t_tvid), bson.M{"$set": bson.M{"ratepeoplenum": (result.RatePeopleNum + 1)} })
		if err != nil {
			log.Fatal(err)
		}
		

		context.JSON(http.StatusOK, gin.H{"status": "rate successfully"})
	})

	// get self video
	authorized.POST("/getSelfVideo", func(context *gin.Context){
		buf := make([]byte, 4096)
		n, _ := context.Request.Body.Read(buf)
		fmt.Println(string(buf))
		for i := 0; i < n; i +=1 {
			if(buf[i] ==' ' || buf[i] == '\x00') {
				buf = append(buf[:i], buf[(i + 1):]...)
				n -= 1
			}
		}

		map2 := make(map[string]string)

		err = json.Unmarshal(buf[:n], &map2)
		fmt.Println(map2)

		t_upName := map2["upname"]

		playInfo_collection := session.DB("meijuAPI").C("PlayInfo")
		//查询文档
		result := PlayInfo{}
		//注意mongodb存储后的字段大小写问题
		iter := playInfo_collection.Find(bson.M{"upname": t_upName}).Iter()
		fmt.Println( t_upName )
		tVInfo_collection := session.DB("meijuAPI").C("TVInfo")
		var data []TV 
		for iter.Next(&result) {
			//查询文档
			fmt.Println(result.TVID)
			result_t := TV{}
			//注意mongodb存储后的字段大小写问题
			tVInfo_collection.FindId(result.TVID).One(&result_t)
			data = append(data, result_t)
			
		}

		context.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
	})


	authorized.POST("/deleteSelfVideo", func(context *gin.Context){
		buf := make([]byte, 4096)
		n, _ := context.Request.Body.Read(buf)
		fmt.Println(string(buf))
		for i := 0; i < n; i +=1 {
			if(buf[i] ==' ' || buf[i] == '\x00') {
				buf = append(buf[:i], buf[(i + 1):]...)
				n -= 1
			}
		}

		map2 := make(map[string]string)

		err = json.Unmarshal(buf[:n], &map2)
		fmt.Println(map2)

		t_tvid := map2["tvid"]

		playInfo_collection := session.DB("meijuAPI").C("PlayInfo")
		tVInfo_collection := session.DB("meijuAPI").C("TVInfo")
		lmsgs_collection := session.DB("meijuAPI").C("LeavingMsgs")
		tVInfo_collection.RemoveId(bson.ObjectIdHex(t_tvid))
		playInfo_collection.RemoveId(bson.ObjectIdHex(t_tvid))
		lmsgs_collection.RemoveId(bson.ObjectIdHex(t_tvid))

		context.JSON(http.StatusOK, gin.H{"status": "delete successfully"})
	})

	authorized.GET("/getVideoInfo", func(context *gin.Context)  {
		tVInfo_collection := session.DB("meijuAPI").C("TVInfo")
		//查询文档
		result := TV{}
		//注意mongodb存储后的字段大小写问题
		iter := tVInfo_collection.Find(bson.M{}).Iter()

		var data []TV 
		for iter.Next(&result) {
			data = append(data, result)
		}
		context.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
	})

	authorized.GET("/getPlayInfo", func(context *gin.Context)  {
		tvID, ok:= context.GetQuery("tvid")
		if !ok {
			fmt.Println("参数不存在")
			return
		}

		playInfo_collection := session.DB("meijuAPI").C("PlayInfo")
		//查询文档
		result := PlayInfo{}
		//注意mongodb存储后的字段大小写问题
		
		err := playInfo_collection.FindId(bson.ObjectIdHex(tvID)).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		playInfo_collection.UpsertId(bson.ObjectIdHex(tvID), bson.M{"$set": bson.M{"viewtime": result.ViewTime + 1} })

		result.ViewTime += 1
		context.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
	})

	authorized.GET("/getMessage", func(context *gin.Context)  {
		tvID, ok:= context.GetQuery("tvid")
		if !ok {
			fmt.Println("参数不存在")
			return
		}

		lmsgs_collection := session.DB("meijuAPI").C("LeavingMsgs")
		//查询文档
		result := LeavingMessage{}
		//注意mongodb存储后的字段大小写问题
		err := lmsgs_collection.FindId(bson.ObjectIdHex(tvID)).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		context.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
	})

	authorized.GET("/getSingleTVInfo", func(context *gin.Context)  {
		tvID, ok:= context.GetQuery("tvid")
		if !ok {
			fmt.Println("参数不存在")
			return
		}

		tvInfo_collection := session.DB("meijuAPI").C("TVInfo")
		//查询文档
		result := TV{}
		//注意mongodb存储后的字段大小写问题
		err := tvInfo_collection.FindId(bson.ObjectIdHex(tvID)).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		context.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
	})

	authorized.GET("/getSingleUserInfo", func(context *gin.Context)  {
		tvID, ok:= context.GetQuery("tvid")
		if !ok {
			fmt.Println("参数不存在")
			return
		}

		userInfo_collection := session.DB("meijuAPI").C("userInformation")
		//查询文档
		result := UserInfo{}
		//注意mongodb存储后的字段大小写问题
		err := userInfo_collection.FindId(bson.ObjectIdHex(tvID)).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		context.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
	})

	authorized.GET("/getVideo/:name", DownVideo)
	authorized.GET("/getImage/:name", DownImage)

	router.Run(":8080");
}


//文件下载功能实现
func DownVideo(c *gin.Context)  {
	//通过动态路由方式获取文件名，以实现下载不同文件的功能
	name:=c.Param("name")
	//拼接路径,如果没有这一步，则默认在当前路径下寻找
	filename:=path.Join("./video/",name)
	// fmt.Println(filename)
	//响应一个文件
	c.File(filename)
	
	return
}

//文件下载功能实现
func DownImage(c *gin.Context)  {
	//通过动态路由方式获取文件名，以实现下载不同文件的功能
	name:=c.Param("name")
	//拼接路径,如果没有这一步，则默认在当前路径下寻找
	filename:=path.Join("./image/",name)
	fmt.Println(filename)
	//响应一个文件
	c.File(filename)
	return
}
