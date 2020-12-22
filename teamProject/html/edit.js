var self = this;

var Main = {
    data() {
      return {
        drawer: false,
        innerDrawer: false,
        TVinfos: [],
        show_TVinfos: [],
        id: '',
        input: '',
        username: '',
        currentDate: new Date(),
        TVinfos: [],
        url: "172.26.29.223",
        fileList: [{name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}, {name: 'food2.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}]
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
        this.id = newsid

        // /getSingleUserInfo
        this.$http.get('http://172.26.65.118:8080/api/getSingleUserInfo',{
            params: {
              tvid: newsid
            },
            },{
              // emulateJSON: true
            }).then(function(res){
              // this.userInfo = res.body.data;
              // console.log(this.userInfo);
              console.log(res)
              this.username = res.body.data.Nickname
              this.getVideo(res.body.data.Nickname)
          });
      },
      getVideo: function (name) {
        this.$http.post('http://172.26.65.118:8080/api/getSelfVideo',{
            upname: name
          },{
            // emulateJSON: true
          }).then(function(res){
            this.TVinfos = res.body.data;
            this.show_TVinfos = this.TVinfos;
            console.log(this.TVinfos);
        });
      }
      ,
      changetoedit: function(){
        window.location.href = url;
      },
      search_click(event) {
        this.$http.get('http://172.26.29.223:8080/api/getVideoInfo',{
          params: {
            
          },
          },{
            // emulateJSON: true
          }).then(function(res){
            this.TVinfos = res.body.data;
            console.log(this.TVinfos);
          });
  
      },

      myDetail_click(i) {
        var url = "http://172.26.29.223:8080/html/play.html?tvid=" + this.TVinfos[i].TVID + "=" + this.id;
        console.log(url);
        window.location.href = url;
      },

      k_delete(i){
        var t_tvid = this.TVinfos[i].TVID;
        console.log(t_tvid);
        this.$http.post('http://172.26.65.118:8080/api/deleteSelfVideo',{
            tvid: t_tvid 
          },{
            // emulateJSON: true
          }).then(function(res){
            console.log(res)
        });
        this.getVideo(this.username)
      }
    }
  }
var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')