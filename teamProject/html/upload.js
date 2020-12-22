// var Main = {
//     data() {
//         return {
//     　　　　form: {
//       　　　　　name: '',
//               videoUrl1: ''
//     　　　　　},
//             videoForm: {
//                 storageurl: ''// 视频地址
//             },
//             actionUrl: '',
//             load:false,
//            fileList:[],
//     　　　　param:new FormData(),//表单要提交的参数
//     　　　　src:"https://afp.alicdn.com/afp-creative/creative/u124884735/14945f2171400c10764ab8f3468470e4.jpg" //展示图片的地址
//         };
//     },
//     methods: {
//         //阻止upload的自己上传，进行再操作
//         onchange(file,filesList) {
//             console.log(file);
//             // //创建临时的路径来展示图片
//             // var windowURL = window.URL || window.webkitURL;
//             // this.src=windowURL.createObjectURL(file.raw);
//             //重新写一个表单上传的方法
//             this.param.append('file', file.raw, file.name);
//             console.log(this.param);
//         },      
//         beforeUploadVideo (file) {
//             console.log(file);
//             // //创建临时的路径来展示图片
//             // var windowURL = window.URL || window.webkitURL;
//             // this.src=windowURL.createObjectURL(file.raw);
//             //重新写一个表单上传的方法
//             this.param.append('file', file.raw, file.name);
//             console.log(this.param);
//             // const isLt20M = file.size / 1024 / 1024 < 30;
//             // if (['video/mp4'].indexOf(file.type) == -1) { //'video/ogg', 'video/flv', 'video/avi', 'video/wmv', 'video/rmvb'
//             //     this.$message.error('请上传正确的视频格式');
//             //     return false;
//             // }
//             // if (!isLt20M) {
//             //     this.$message.error('上传视频大小不能超过20MB哦!');
//             //     return false;
//             // }
//             // this.isShowUploadVideo = false;
//             // this.load=true;
//         },

	
// 	//上传成功回调
// 	handleVideoSuccess (res, file) {
//             this.form.videoUrl1 = URL.createObjectURL(file.raw);
//             this.form.videoUrl = res.data;
             
            
//             if (res.code == 1) {
//               this.load = false;

                
//             } else {
//                 this.$message.error('视频上传失败，请重新上传！');
//             }
              
//   },
//         handleRemove(file,filesList){
//             this.param.delete('file')
//         },
//         onSubmit(){//表单提交的事件
//             var names = this.form.name;
//             //下面append的东西就会到form表单数据的fields中；
//             this.param.append('message', names);
//             let config = {
//                     headers: {
//                         'Content-Type': 'multipart/form-data'
//                     }
//                 };
//             //然后通过下面的方式把内容通过axios来传到后台
//             //下面的this.$reqs 是在主js中通过Vue.prototype.$reqs = axios 来把axios赋给它;
//             this.$reqs.post("/upload", this.param, config).then(function(result) {
//                     console.log(result);
//             })
//         }
//     }
//   }
// var Ctor = Vue.extend(Main)
// new Ctor().$mount('#app')
// // new Ctor().$mount('#app')

// <div id = "app">
//     <el-row>
//             <el-col :span="8" :offset="8">
//                 <div id="upload">
//             <!--elementui的form组件-->
//                     <el-form ref="form" :model="form" label-width="80px">
//                         <el-form-item label="主题">
//                         <el-input v-model="form.name" name="names" style="width:360px;"></el-input>
//                     </el-form-item>
    
//               <label class="el-form-item__label" style="width: 80px;">上传图片</label>
//               <!--elementui的上传图片的upload组件-->
//               <!--
//                 :auto-upload=false  // 取消自动上传
//                 :on-remove="handleRemove" // 处理删除图片的操作
//                 :on-change="onchange" // 通过onchange这个属性来获取现在的图片和所有准备上传的图片
//                 :limit=1 // 限制只能上传一张，这里暂时只考虑一张图片的情况
//                 drag // 设置这个让可以把图片拖进来上传
//                 action="" // 这里暂时不设置上传地址，因为我们是要拦截在form中上传
//                 -->
    
//               <el-upload
//                 class="upload-demo"
//                 :auto-upload=false
//                 :on-change="onchange"
//                 :on-remove="handleRemove"
//                 :limit=1
//                 :file-list="fileList"
//                 list-type="picture"
//                 drag
//                 action=""
//                 style="margin-left:80px;">
//                 <i class="el-icon-upload"></i>
//                 <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
//                 <div class="el-upload__tip" slot="tip">这里只能上传一张,如需更换请先手动删除列表中的！</div>
//               </el-upload>
//               <el-form-item style="padding-top:20px;">
//                 <el-button type="primary" @click="onSubmit">立即创建</el-button>
//               </el-form-item>
//               <el-form-item label="上传视频" prop="videoUrl">   

//                 <el-upload
//                 class="avatar-uploader"
//                 :action="actionUrl"
//                 :show-file-list="false"
//                 :on-success="handleVideoSuccess"
//                 :before-upload="beforeUploadVideo"
//                 >
//                 <video v-if="form.videoUrl1" :src="form.videoUrl1" class="avatar"></video>
//                 <i v-else class="el-icon-plus avatar-uploader-icon"></i>
//                   <p v-if='load'> 上传中</p>
//                    <P style="color:#000000">请保证视频格式正确，且不超过20M</P>
    
//                 </el-upload>
                
//               </el-form-item>
//                     </el-form>
//                 </div>
//             </el-col>
            
//         </el-row>
// </div>


new Vue({
    el: "#uploadBox",
    data(){
        return {
            loading : false,
            loading_message : "上传",
            dialogVisible: false
        }
    },
    methods:{
        handleClose(done) {
            this.$confirm('确认关闭？')
                .then(_ => {
                done();
                })
                .catch(_ => {});
        },
        change_to_index: function(){
            this.dialogVisible = false
            const loading = this.$loading({
                lock: true,
                text: 'Loading',
                spinner: 'el-icon-loading',
                background: 'rgba(0, 0, 0, 0.7)'
                });
                setTimeout(() => {
                loading.close();
            }, 2000);

            var form = document.forms["upload"];
            console.log(form);
            form.action = "http://172.26.65.118:8080/uploadfiles";
            form.submit();
            // .then(function(res){
            //     console.log(res);
            //     window.location.href = "http://172.26.65.118:8080/html/index.html";
            // });
        
        },
    }
})


$('#hidden_frame').load(function(){
    window.location.href = "http://172.26.65.118:8080/html/index.html";
});

