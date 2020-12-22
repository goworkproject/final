var Main = {
  data() {
    return {
      PlayInfo: {},
      //播放页面信息
      // type PlayInfo struct {
      // 	TVID	bson.ObjectId `bson:"_id,omitempty"` //类型是bson.ObjectId
      // 	UpID	bson.ObjectId 
      // 	VideoUrl string //shi pin dizhi
      // 	Brief string 	//jian jie
      // 	ViewTime int	//guankan cishu
      // 	SetNum int		//duoshao ji
      // }
      TVInfo: {},
      //TV信息，生成的ID主键与PlayInfo、LeavingMessage的ID主键是一样的
      // type TV struct {
      //   TVID	bson.ObjectId `bson:"_id,omitempty"` //类型是bson.ObjectId
      //   Name string
      //   ImageUrl string
      // }


      Messages: {},
      //留言的单元，储存留言的信息
      // type Message struct {
      //   Sender string
      //   Receiver string
      //   Content string
      //   Time time.Time
      // }

      // //TV的留言
      // type LeavingMessage struct {
      //   TVID	bson.ObjectId `bson:"_id,omitempty"` //类型是bson.ObjectId
      //   Messages []Message
      // }
      userInfo: {},
      //用户个人信息
      // type UserInfo struct {
      //   ID	bson.ObjectId `bson:"_id,omitempty"` //类型是bson.ObjectId
      //   Nickname string
      //   Password string
      // }

      value1: null,
      value2: null,
      value: 3.7,
      show3: true,
      colors: ['#99A9BF', '#F7BA2A', '#FF9900'],  // 等同于 { 2: '#99A9BF', 4: { value: '#F7BA2A', excluded: true }, 5: '#FF9900' }
      show_time: 0,
      description:"hello,world",
      author_name: "Σ(☉▽☉翔",
      saying_message: "五得闲天哈第一",
      chat_message:'',
      detail:''
    };
  },
  // invoke when start 
  mounted:function(){
      this.getTVID();
  },
  methods: {
      getTVID() {
          urlinfo=window.location.href;  //获取当前页面的url
          len=urlinfo.length;//获取url的长度
          offset=urlinfo.indexOf("?");//设置参数字符串开始的位置
          newsidinfo=urlinfo.substr(offset,len)//取出参数字符串 这里会获得类似“id=1”这样的字符串
          newsids=newsidinfo.split("=");//对获得的参数字符串按照“=”进行分割
          newsid=newsids[1];//得到参数值
          newUserId = newsids[2];
          //
          console.log(newUserId)
          //
          this.$http.get('http://172.26.65.118:8080/api/getSingleTVInfo',{
            params: {
              tvid: newsid
            },
            },{
              // emulateJSON: true
            }).then(function(res){
              this.TVInfo = res.body.data;
              this.description = this.TVInfo.Name;
              // console.log(this.TVInfo);
          });

          this.$http.get('http://172.26.65.118:8080/api/getSingleUserInfo',{
            params: {
              tvid: newUserId
            },
            },{
              // emulateJSON: true
            }).then(function(res){
              this.userInfo = res.body.data;
              console.log(this.userInfo);
          });

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
                this.value2  = 5;
              } else {
                this.value2 = (this.PlayInfo.AllRateNum / this.PlayInfo.RatePeopleNum);
              }
          });

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
      },

      post_click(event) {
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
      },
      
      rate_click (event) {

        this.$http.post('http://172.26.65.118:8080/api/thumbUp',{
            tvid: this.PlayInfo.TVID,
            rate: ""+this.value2,
            },{
              // emulateJSON: true
            }).then(function(res){
              console.log(res);
          });

      },

      errorHandler() {
        return true;
      }
}
}

var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')