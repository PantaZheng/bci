webpackJsonp([3],{"09xf":function(e,t){},"3g8P":function(e,t){},BBNV:function(e,t){},BcqE:function(e,t){},CJrA:function(e,t){},GkLy:function(e,t){},HhP9:function(e,t){},g0cZ:function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0});String;var s={name:"projectHeader",props:{Name:String},data:function(){return{project_id:sessionStorage.getItem("projectId")}}},a={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"header"},[t("div",{staticClass:"header-left back-icon"},[t("router-link",{attrs:{to:"/project/"+this.project_id}},[t("svg",{staticClass:"icon arrow",attrs:{"aria-hidden":"true"}},[t("use",{attrs:{"xlink:href":"#iconleft"}})])])],1),this._v(" "),t("div",{staticClass:"header-title v-center"},[this._v("\n    "+this._s(this.Name)+"\n  ")]),this._v(" "),t("div",{staticClass:"header-right"},[t("svg",{staticClass:"icon ",attrs:{"aria-hidden":"true"}},[t("use",{attrs:{"xlink:href":"#icondian2"}})])])])},staticRenderFns:[]};var n=i("VU/8")(s,a,!1,function(e){i("09xf")},"data-v-8ea6f558",null).exports,l=(Object,{name:"detail",props:{module:Object}}),c={render:function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",[i("el-divider"),e._v(" "),i("div",[i("div",{staticClass:"weui-cell"},[e._v(e._s(e.module.content))]),e._v(" "),i("div",{staticClass:"weui-cells"},[i("div",{staticClass:"weui-cell"},[e._m(0),e._v(" "),i("div",{staticClass:"weui-cell__bd cellMargin"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[e._v("\n            "+e._s(e.module.startTime)+"\n          ")])])]),e._v(" "),i("div",{staticClass:"weui-cell"},[e._m(1),e._v(" "),i("div",{staticClass:"weui-cell__bd cellMargin"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[e._v("\n            "+e._s(e.module.endTime)+"\n          ")])])]),e._v(" "),i("div",{staticClass:"weui-cell"},[e._m(2),e._v(" "),i("div",{staticClass:"weui-cell__bd cellMargin"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[e._v("\n            "+e._s(e.module.leader.name)+"\n          ")])])]),e._v(" "),i("div",{staticClass:"weui-cell"},[e._m(3),e._v(" "),i("div",{staticClass:"weui-cell__bd cellMargin"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},e._l(e.module.participants,function(t){return i("span",{key:t.id,staticClass:"nameMargin"},[e._v(e._s(t.name))])}),0)])])])])],1)},staticRenderFns:[function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"weui-cell__hd"},[t("label",{staticClass:"weui-label"},[this._v("开始时间")])])},function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"weui-cell__hd"},[t("label",{staticClass:"weui-label"},[this._v("结束时间")])])},function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"weui-cell__hd"},[t("label",{staticClass:"weui-label"},[this._v("模块负责人")])])},function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"weui-cell__hd"},[t("label",{staticClass:"weui-label"},[this._v("模块参与学生")])])}]};var o=i("VU/8")(l,c,!1,function(e){i("HhP9")},"data-v-34cec0c5",null).exports,r=(Object,{name:"detailHeader",components:{Detail:o},props:{module:Object},methods:{handleChange:function(e){console.log(e)}}}),u={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",[t("el-collapse",{on:{change:this.handleChange}},[t("el-collapse-item",{attrs:{title:"点击查看模块详情"}},[t("detail",{attrs:{module:this.module}})],1)],1)],1)},staticRenderFns:[]};var d=i("VU/8")(r,u,!1,function(e){i("CJrA")},"data-v-55e9a743",null).exports,m=i("mtWM"),_=i.n(m),v={name:"runMission",data:function(){return{missionList:[]}},methods:{getMissionList:function(){_.a.get(this.url+"/mission/module/"+sessionStorage.getItem("moduleId")).then(this.getMissionListSucc)},getMissionListSucc:function(e){(e=e.data)&&(this.missionList=e,console.log(this.missionList))}},mounted:function(){this.getMissionList()}},h={render:function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",[i("div",{staticClass:"weui_tab_bd_item "},[i("div",{staticClass:"weui-news"},[i("ul",{staticClass:"weui-news-list"},e._l(e.missionList,function(t){return i("li",{key:t.id},[i("router-link",{staticClass:"weui-news-item",attrs:{to:"/project/module/mission/"+t.id}},[i("div",{staticClass:"weui-news-inners"},[i("div",{staticClass:"weui-news-title"},[e._v(e._s(t.name)),i("span",{staticClass:"timeStyle"},[e._v(e._s(t.startTime))])]),e._v(" "),i("div",{staticClass:"weui-news-text"},[i("p",{staticClass:"weui-news-p"},[e._v(e._s(t.content))])])])])],1)}),0)])])])},staticRenderFns:[]};var w=i("VU/8")(v,h,!1,function(e){i("GkLy")},"data-v-95bc9148",null).exports,f=(Array,{name:"finishedMission",props:{mission:Array}}),p={render:function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",[i("div",{staticClass:"weui_tab_bd_item "},[i("div",{staticClass:"weui-news"},[i("ul",{staticClass:"weui-news-list"},e._l(e.mission,function(t){return 1==t.tag?i("li",{key:t.id},[i("router-link",{staticClass:"weui-news-item",attrs:{to:"/project/module/mission/"+t.id}},[i("div",{staticClass:"weui-news-inners"},[i("div",{staticClass:"weui-news-title"},[e._v(e._s(t.name)),i("span",{staticClass:"timeStyle"},[e._v(e._s(t.create_time))])]),e._v(" "),i("div",{staticClass:"weui-news-text"},[i("p",{staticClass:"weui-news-p"},[e._v(e._s(t.content))])])])])],1):e._e()}),0)])])])},staticRenderFns:[]};var C=i("VU/8")(f,p,!1,function(e){i("yNT2")},"data-v-050cc107",null).exports,g=i("mvHQ"),b=i.n(g),y={name:"newMission",data:function(){return{fileList:[],mission:{creatorID:localStorage.getItem("id"),name:"",startTime:"",endTime:"",content:"",moduleID:parseInt(sessionStorage.getItem("moduleId")),participants:[]}}},methods:{handleRemove:function(e,t){console.log(e,t)},handlePreview:function(e){console.log(e)},beforeRemove:function(e,t){return this.$confirm("确定移除"+e.name+"？")},getStudentList:function(){_.a.get(this.url+"/user/level/3").then(this.getStudentListSucc)},getStudentListSucc:function(e){(e=e.data)&&(this.StudentList=e)},submit:function(){this.$axios({method:"post",url:this.url+"/mission",data:b()(this.mission)}).then(function(e){200===e.status?alert("成功创建任务！"):alert(e)}).catch(function(){alert("创建失败！")})},mounted:function(){this.getStudentList()}}},x={render:function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",{staticClass:"weui_tab_bd_item formMargin"},[i("form",[i("div",{staticClass:"weui_cells"},[i("div",{staticClass:"weui_cell weui"},[e._m(0),e._v(" "),i("div",{staticClass:"weui_cell_bd weui_cell_primary "},[i("input",{directives:[{name:"model",rawName:"v-model",value:e.mission.name,expression:"mission.name"}],staticClass:"weui_input",attrs:{type:"",value:"",placeholder:"请输入任务名称"},domProps:{value:e.mission.name},on:{input:function(t){t.target.composing||e.$set(e.mission,"name",t.target.value)}}})])]),e._v(" "),i("div",{staticClass:"weui_cell"},[e._m(1),e._v(" "),i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[i("input",{directives:[{name:"model",rawName:"v-model",value:e.mission.startTime,expression:"mission.startTime"}],staticClass:"weui_input",attrs:{type:"date",value:""},domProps:{value:e.mission.startTime},on:{input:function(t){t.target.composing||e.$set(e.mission,"startTime",t.target.value)}}})])]),e._v(" "),i("div",{staticClass:"weui_cell"},[e._m(2),e._v(" "),i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[i("input",{directives:[{name:"model",rawName:"v-model",value:e.mission.endTime,expression:"mission.endTime"}],staticClass:"weui_input",attrs:{type:"date",value:""},domProps:{value:e.mission.endTime},on:{input:function(t){t.target.composing||e.$set(e.mission,"endTime",t.target.value)}}})])])]),e._v(" "),i("el-select",{staticClass:"selectStyle",attrs:{multiple:"",placeholder:"请选择参与模块学生","value-key":"id"},model:{value:e.mission.participants,callback:function(t){e.$set(e.mission,"participants",t)},expression:"mission.participants"}},e._l(e.StudentList,function(e){return i("el-option",{key:e.id,attrs:{label:e.name,value:e}})}),1),e._v(" "),i("div",{staticClass:"weui_cells"},[i("div",{staticClass:"weui_cell"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[i("textarea",{directives:[{name:"model",rawName:"v-model",value:e.mission.content,expression:"mission.content"}],staticClass:"weui_textarea",attrs:{placeholder:"请输入主要内容",rows:"3"},domProps:{value:e.mission.content},on:{input:function(t){t.target.composing||e.$set(e.mission,"content",t.target.value)}}}),e._v(" "),i("div",{staticClass:"weui_textarea_counter iconSize"},[i("el-upload",{ref:"upload",staticClass:"upload-demo",attrs:{action:"https://jsonplaceholder.typicode.com/posts/","on-preview":e.handlePreview,"on-remove":e.handleRemove,"file-list":e.fileList,"auto-upload":!1}},[i("el-button",{attrs:{slot:"trigger",type:"primary",size:"small"},slot:"trigger"},[e._v("\n                  点击上传附件\n                ")])],1)],1)])])]),e._v(" "),i("div",{staticClass:"weui-btn-area",on:{click:function(t){return e.submit()}}},[i("a",{staticClass:"weui-btn weui-btn_primary",attrs:{href:"javascript:;"}},[e._v("提交")])])],1)])},staticRenderFns:[function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"weui_cell_hd "},[t("label",{staticClass:"weui_label weui_cell_primary"},[this._v("名称")])])},function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"weui_cell_hd"},[t("label",{staticClass:"weui_label"},[this._v("开始")])])},function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"weui_cell_hd"},[t("label",{staticClass:"weui_label"},[this._v("结束")])])}]};var M=i("VU/8")(y,x,!1,function(e){i("BBNV")},"data-v-14fd2aa5",null).exports,S=(Array,{name:"moduleHeader",components:{RunMission:w,FinishedMission:C,NewMission:M},props:{mission:Array},data:function(){return{activeIndex:"1",Run:"",Finished:"",New:""}},mounted:function(){this.Run=!0},methods:{handleSelect:function(e,t){console.log(e,t)},showRun:function(){this.Run=!0,this.New=!1,this.Finished=!1},showFinished:function(){this.Finished=!1,this.Finished=!0,this.Run=!1,this.New=!1},showNew:function(){this.Run=!1,this.New=!1,this.New=!0,this.Finished=!1}}}),N={render:function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",[i("div",{staticClass:"weui_tab"},[i("el-menu",{staticClass:"el-menu-demo",attrs:{"default-active":e.activeIndex,mode:"horizontal","background-color":"#fff","text-color":"#303133","active-text-color":"#16b0ff"},on:{select:e.handleSelect}},[i("div",{staticClass:"weui_navbar"},[i("el-menu-item",{staticClass:"weui_navbar_item ",attrs:{index:"1"},on:{click:e.showRun}},[e._v("正在进行")]),e._v(" "),i("el-menu-item",{staticClass:"weui_navbar_item ",attrs:{index:"2"},on:{click:e.showFinished}},[e._v("已完成")]),e._v(" "),i("el-menu-item",{staticClass:"weui_navbar_item ",attrs:{index:"3"},on:{click:e.showNew}},[e._v("新建任务")])],1)])],1),e._v(" "),e.New?i("new-mission",{staticClass:"componentsMargin"}):e._e(),e._v(" "),e.Run?i("run-mission",{staticClass:"componentsMargin",attrs:{mission:e.mission}}):e._e(),e._v(" "),e.Finished?i("finished-mission",{staticClass:"componentsMargin",attrs:{mission:e.mission}}):e._e()],1)},staticRenderFns:[]};var $=i("VU/8")(S,N,!1,function(e){i("3g8P")},"data-v-144bcf3d",null).exports,k={name:"project",components:{ModuleHeader:n,ModuleDetail:d,ModuleMissionHeader:$},data:function(){return{module:Object,mission:[],Name:this.Name}},methods:{getModule:function(){sessionStorage.setItem("moduleId",this.$route.params.moId),_.a.get(this.url+"/module/id/"+this.$route.params.moId).then(this.getModuleSucc)},getModuleSucc:function(e){if(e){var t=e.data;this.module=t,this.mission=t.missions,this.Name=t.name}}},mounted:function(){this.getModule()}},R={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",[t("module-header",{attrs:{Name:this.Name}}),this._v(" "),t("module-detail",{attrs:{module:this.module}}),this._v(" "),t("module-mission-header",{attrs:{mission:this.mission}})],1)},staticRenderFns:[]};var L=i("VU/8")(k,R,!1,function(e){i("BcqE")},"data-v-7ff74a68",null);t.default=L.exports},yNT2:function(e,t){}});
//# sourceMappingURL=3.df29373ecdb01f324f7b.js.map