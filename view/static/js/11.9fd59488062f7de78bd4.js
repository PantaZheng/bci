webpackJsonp([11],{J8xq:function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var i={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"header"},[t("div",{staticClass:"header-left back-icon"},[t("router-link",{attrs:{to:"/"}},[t("svg",{staticClass:"icon arrow",attrs:{"aria-hidden":"true"}},[t("use",{attrs:{"xlink:href":"#iconleft"}})])])],1),this._v(" "),t("div",{staticClass:"header-title v-center"},[this._v("\n    添加用户\n  ")]),this._v(" "),t("div",{staticClass:"header-right"})])},staticRenderFns:[]};var s=a("VU/8")({name:"addUserHeader"},i,!1,function(e){a("WudP")},"data-v-7bb5092c",null).exports,l=a("mvHQ"),n=a.n(l),r={name:"addUserForm",data:function(){return{item:{name:"",telephone:"",level:Number},level:[{value:6,label:"最高负责人"},{value:4,label:"老师"},{value:5,label:"助理"},{value:3,label:"学生"},{value:2,label:"专家教授"}]}},methods:{deleteMe:function(){this.$axios.delete(this.url+"/user/id/7")},submit:function(){var e=this;""!==this.item.name&&""!==this.item.telephone?1!==this.item.level?this.$axios({method:"post",url:this.url+"/user/",data:n()(this.item)}).then(function(t){200===t.status?(alert("添加成功"),e.item=[]):alert(t)}).catch(function(e){alert(e)}):this.$notify.warning({title:"请选择用户身份!"}):this.$notify.warning({title:"请输入姓名和电话!"})}},created:function(){this.item.level=1}},u={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("div",{staticClass:"weui-cells form weui-cells_form"},[a("div",{staticClass:"weui-cell"},[e._m(0),e._v(" "),a("div",{staticClass:"weui-cell__bd "},[a("input",{directives:[{name:"model",rawName:"v-model",value:e.item.name,expression:"item.name"}],staticClass:"weui-input",attrs:{pattern:"[0-9]*",required:"",placeholder:"请输入姓名",type:"name"},domProps:{value:e.item.name},on:{input:function(t){t.target.composing||e.$set(e.item,"name",t.target.value)}}})])]),e._v(" "),a("div",{staticClass:"weui-cell"},[e._m(1),e._v(" "),a("div",{staticClass:"weui-cell__bd"},[a("input",{directives:[{name:"model",rawName:"v-model",value:e.item.telephone,expression:"item.telephone"}],staticClass:"weui-input",attrs:{placeholder:"请输入手机号",type:"tel"},domProps:{value:e.item.telephone},on:{input:function(t){t.target.composing||e.$set(e.item,"telephone",t.target.value)}}})])]),e._v(" "),a("div",{staticClass:"weui_cell weui_cell_select"},[a("div",{staticClass:"weui_cell_bd weui_cell_primary"},[a("select",{directives:[{name:"model",rawName:"v-model",value:e.item.level,expression:"item.level"}],staticClass:"weui_select",on:{change:function(t){var a=Array.prototype.filter.call(t.target.options,function(e){return e.selected}).map(function(e){return"_value"in e?e._value:e.value});e.$set(e.item,"level",t.target.multiple?a:a[0])}}},[a("option",{attrs:{value:"1"}},[e._v("请选择用户角色")]),e._v(" "),e._l(e.level,function(e){return a("option",{key:e.value,attrs:{label:e.label},domProps:{value:e.value}})})],2)])])]),e._v(" "),a("div",{staticClass:"weui-btn-area",on:{click:function(t){return e.submit()}}},[a("a",{staticClass:"weui-btn weui-btn_primary",attrs:{href:"javascript:;"}},[e._v("提交")])]),e._v(" "),a("div",{staticClass:"weui-btn-area",on:{click:function(t){return e.deleteMe()}}},[a("a",{staticClass:"weui-btn weui-btn_primary",attrs:{href:"javascript:;"}},[e._v("删除我")])])])},staticRenderFns:[function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"weui-cell__hd"},[t("label",{staticClass:"weui-label"},[this._v("姓名")])])},function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"weui-cell__hd"},[t("label",{staticClass:"weui-label"},[this._v("手机号")])])}]};var c=a("VU/8")(r,u,!1,function(e){a("eanI")},"data-v-4e156cae",null).exports,o={name:"addUser",components:{AddUserHeader:s,AddUserForm:c}},v={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",[t("add-user-header"),this._v(" "),t("add-user-form")],1)},staticRenderFns:[]};var d=a("VU/8")(o,v,!1,function(e){a("VDyT")},"data-v-cdf741ec",null);t.default=d.exports},VDyT:function(e,t){},WudP:function(e,t){},eanI:function(e,t){}});
//# sourceMappingURL=11.9fd59488062f7de78bd4.js.map