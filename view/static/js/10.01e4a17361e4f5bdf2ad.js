webpackJsonp([10],{QEAu:function(t,e){},iIZT:function(t,e){},jdtw:function(t,e){},zHGF:function(t,e,i){"use strict";Object.defineProperty(e,"__esModule",{value:!0});String;var a={name:"examineHeader",props:{name:String}},s={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"header"},[e("div",{staticClass:"header-left back-icon"},[e("router-link",{attrs:{to:"/newProject"}},[e("svg",{staticClass:"icon arrow",attrs:{"aria-hidden":"true"}},[e("use",{attrs:{"xlink:href":"#iconleft"}})])])],1),this._v(" "),e("div",{staticClass:"header-title v-center"},[this._v("\n    "+this._s(this.name)+"\n  ")]),this._v(" "),e("div",{staticClass:"header-right"})])},staticRenderFns:[]};var l=i("VU/8")(a,s,!1,function(t){i("jdtw")},"data-v-18e6a81a",null).exports,n=i("mvHQ"),r=i.n(n),c=(Object,{name:"assistantForm",props:{project:Object},data:function(){return{item:{name:"",creatorID:this.project.creatorID,start_time:"",end_time:"",content:"",leaderID:this.project.leaderID,targets:this.tags},list:[],target:""}},methods:{getTarget:function(){this.item.targets=this.target.split("，")},submit:function(){console.log(this.project),this.tags=this.project.tagSet,console.log(this.tags),this.tags.forEach(function(t){t.id===sessionStorage.getItem("id")&&(t.tag=!0)}),this.$axios({method:"post",url:this.url+"/project/",data:{creatorID:this.project.creatorID,leaderID:this.project.leaderID,startTime:this.item.startTime,tagSet:this.tags}}).then(function(t){console.log(t),alert("成功创建项目！")}).catch(function(t){alert(t)}),this.$axios({method:"post",url:this.url+"/project/",data:r()(this.item)}).then(function(t){console.log(t),alert("成功创建项目！")}).catch(function(t){alert(t)})}}}),o={render:function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",{staticClass:"page-bg"},[i("div",{staticClass:"weui-cells"},[i("div",{staticClass:"weui-cell"},[t._m(0),t._v(" "),i("div",{staticClass:"weui-cell__bd "},[i("input",{directives:[{name:"model",rawName:"v-model",value:t.item.name,expression:"item.name"}],staticClass:"weui-input",attrs:{pattern:"[0-9]*",placeholder:"请输入项目名称",type:"name",value:"project.name"},domProps:{value:t.item.name},on:{input:function(e){e.target.composing||t.$set(t.item,"name",e.target.value)}}})])]),t._v(" "),i("div",{staticClass:"weui-cell"},[t._m(1),t._v(" "),i("div",{staticClass:"weui-cell__bd"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[i("input",{directives:[{name:"model",rawName:"v-model",value:t.item.startTime,expression:"item.startTime"}],staticClass:"weui_input",attrs:{type:"date",value:"project.startTime"},domProps:{value:t.item.startTime},on:{input:function(e){e.target.composing||t.$set(t.item,"startTime",e.target.value)}}})])])]),t._v(" "),i("div",{staticClass:"weui-cell"},[t._m(2),t._v(" "),i("div",{staticClass:"weui-cell__bd"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[i("input",{directives:[{name:"model",rawName:"v-model",value:t.item.endTime,expression:"item.endTime"}],staticClass:"weui_input",attrs:{type:"date",value:"project.endTime"},domProps:{value:t.item.endTime},on:{input:function(e){e.target.composing||t.$set(t.item,"endTime",e.target.value)}}})])])])]),t._v(" "),i("div",{staticClass:"weui-cell"},[t._m(3),t._v(" "),i("div",{staticClass:"weui-cell__bd "},[i("input",{directives:[{name:"model",rawName:"v-model",value:t.item.leader,expression:"item.leader"}],staticClass:"weui-input",attrs:{pattern:"[0-9]*",placeholder:"请输入项目名称",type:"name",value:"project.leader"},domProps:{value:t.item.leader},on:{input:function(e){e.target.composing||t.$set(t.item,"leader",e.target.value)}}})])]),t._v(" "),i("div",{staticClass:"weui-cell"},[t._m(4),t._v(" "),i("div",{staticClass:"weui-cell__bd cellMargin"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},t._l(t.project.participants,function(e){return e.level>4?i("span",{key:e.id,staticClass:"nameMargin"},[t._v("\n                "+t._s(e.name)+"\n            ")]):t._e()}),0)])]),t._v(" "),i("div",{staticClass:"weui-cells weui-cells_form"},[i("el-input",{attrs:{type:"textarea",rows:4,placeholder:"请输入项目验收目标，多个目标请用'，'分隔",clearable:""},on:{input:function(e){return t.getTarget()}},model:{value:t.target,callback:function(e){t.target=e},expression:"target"}})],1),t._v(" "),i("div",{staticClass:"weui-cells weui-cells_form"},[i("el-input",{attrs:{type:"textarea",rows:4,placeholder:"请输入项目主要内容",value:"project.content",clearable:""},model:{value:t.item.content,callback:function(e){t.$set(t.item,"content",e)},expression:"item.content"}})],1),t._v(" "),i("div",{staticClass:"weui-btn-area"},[i("a",{staticClass:"weui-btn weui-btn_primary",attrs:{href:"javascript:;"},on:{click:function(e){return t.submit()}}},[t._v("通过审核")])])])},staticRenderFns:[function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("项目名称")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("开始时间")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("结束时间")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("项目负责人")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("项目参与老师")])])}]};var u=i("VU/8")(c,o,!1,function(t){i("iIZT")},"data-v-4fc6f492",null).exports,m=i("mtWM"),d=i.n(m),v={name:"examine",components:{ExamineHeader:l,ExamineDetail:u},data:function(){return{project:Object,name:String}},methods:{getProject:function(){d.a.get(this.url+"/project/id/"+this.$route.params.pId).then(this.getProjectSucc)},getProjectSucc:function(t){console.log(t.status),(t=t.data)&&(this.project=t,console.log(this.project),this.name=t.name)}},mounted:function(){this.getProject()}},_={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("examine-header",{attrs:{name:this.name}}),this._v(" "),e("examine-detail",{attrs:{project:this.project}})],1)},staticRenderFns:[]};var p=i("VU/8")(v,_,!1,function(t){i("QEAu")},"data-v-5c933889",null);e.default=p.exports}});
//# sourceMappingURL=10.01e4a17361e4f5bdf2ad.js.map