webpackJsonp([11],{AqbR:function(t,e){},IJqf:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var i=n("lbHh"),a=n.n(i),r={render:function(){var t=this.$createElement;return(this._self._c||t)("div",{staticClass:"v-center header page-bg page-hd-title"},[this._v("\n  项目列表\n")])},staticRenderFns:[]};var s=n("VU/8")({name:"frameHeader"},r,!1,function(t){n("tKMr")},"data-v-4d3776de",null).exports,c=n("ifoU"),o=n.n(c),l=n("OMN4"),u=n.n(l),d={name:"frameList",data:function(){return{tag:this.Run,list:[]}},mounted:function(){var t=parseInt(a.a.get("level"));t===this.PI?this.getProjectManager():t===this.assistant?this.getProjectList():(this.getProjectMember(),this.getProjectManager())},methods:{goFrameDetail:function(t,e){sessionStorage.setItem("pid",t),sessionStorage.setItem("pname",e),this.$router.push("/projectFrame")},getProjectList:function(){u.a.get(this.url+"/project/all").then(this.getProjectListSucc)},getProjectManager:function(){u.a.get(this.url+"/project/manager/"+a.a.get("id")).then(this.getProjectListSucc)},getProjectMember:function(){u.a.get(this.url+"/project/member/"+a.a.get("id")).then(this.getProjectListSucc)},getProjectListSucc:function(t){if(200===t.status){t=t.data;var e=this.list.concat(t),n=new o.a;this.list=e.filter(function(t){return!n.has(t.id)&&n.set(t.id,1)})}else alert(t.data)}}},h={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",t._l(t.list,function(e){return e.state===t.tag?n("el-container",{key:e.id,staticStyle:{height:"auto",width:"100%",border:"1px solid #eee",display:"flex"},nativeOn:{click:function(n){return t.goFrameDetail(e.id,e.name)}}},[n("el-main",{staticStyle:{width:"50px",display:"flex",height:"100%"}},[n("span",[n("svg",{staticClass:"icon projectIcon",attrs:{"aria-hidden":"true"}},[n("use",{attrs:{"xlink:href":"#icon-"}})])])]),t._v(" "),n("el-main",{staticStyle:{width:"60%",display:"flex",height:"100%"}},[n("span",[t._v("\n          "+t._s(e.name)+"\n        ")])]),t._v(" "),n("el-main",{staticStyle:{width:"30%",display:"flex",height:"100%"}},[n("span",[n("svg",{staticClass:"icon fuzerenIcon",attrs:{"aria-hidden":"true"}},[n("use",{attrs:{"xlink:href":"#iconren"}})]),t._v("\n          "+t._s(e.managerName)+"\n        ")])])],1):t._e()}),1)},staticRenderFns:[]};var f=n("VU/8")(d,h,!1,function(t){n("g61W")},"data-v-7a059a02",null).exports,g={name:"frame",components:{FrameHeader:s,FrameList:f},methods:{judge:function(){a.a.get("id")||this.$router.push("/loginError")}},created:function(){this.judge()},mounted:function(){window.scrollTo(0,0)}},m={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("frame-header"),this._v(" "),e("frame-list")],1)},staticRenderFns:[]};var p=n("VU/8")(g,m,!1,function(t){n("AqbR")},"data-v-4a0288cf",null);e.default=p.exports},g61W:function(t,e){},tKMr:function(t,e){}});
//# sourceMappingURL=11.8ee2bee3ea5b81bc6cce.js.map