var self = this;

var Main = {
    data() {
      return {
        id: '',
        input: '',
        currentDate: new Date(),
        TVinfos: [],
        show_TVinfos: [],
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
        this.id = newsid;
        this.$http.get('http://172.26.65.118:8080/api/getVideoInfo',{
          params: {
            
          },
          },{
            // emulateJSON: true
          }).then(function(res){
            this.TVinfos = res.body.data;
            this.show_TVinfos = this.TVinfos;
            console.log(this.TVinfos);
        });
      },
      search_click(event) {
        var index = 0;
        this.show_TVinfos = [];
        for (var i = 0; i < this.TVinfos.length; ++i) {
          if (this.TVinfos[i].Name.indexOf(this.input) != -1) {
            console.log(this.TVinfos[i].Name);
            console.log(this.input);
            this.show_TVinfos.push(this.TVinfos[i]);
          }
        }
        console.log(this.show_TVinfos);
      },

      edit_click(event) {
        var url = "http://172.26.65.118:8080/html/upload.html?id=" + this.id;
        window.location.href = url;
      },

      changetoedit:function(){
        var url = "http://172.26.65.118:8080/html/edit.html?id=" + this.id;
        window.location.href = url;
      },
      
      logout: function(){
        this.$http.get('http://172.26.65.118:8080/api/getSingleUserInfo',{
            params: {
              tvid: this.id
            },
            },{
              // emulateJSON: true
            }).then(function(res){
              // this.userInfo = res.body.data;
              // console.log(this.userInfo);
              console.log(res)
              this.username = res.body.data.Nickname
              this.t_logout(res.body.data.Nickname)
          });
      },

      t_logout: function(name) {
        this.$http.post('http://172.26.65.118:8080/api/logOut',{
          nickname: name   
        },{
          // emulateJSON: true
        }).then(function(res){
          console.log(res)
        });
      },

      myDetail_click(i) {
        var url = "http://172.26.65.118:8080/html/play.html?tvid=" + this.TVinfos[i].TVID + "=" + this.id;
        console.log(url);
        window.location.href = url;
      }
    }
  }
var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')