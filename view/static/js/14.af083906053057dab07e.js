webpackJsonp([14],{"75Op":function(t,e,i){"use strict";Object.defineProperty(e,"__esModule",{value:!0});String;var s={name:"header",props:{name:String}},a={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"header"},[e("div",{staticClass:"header-left back-icon"},[e("router-link",{attrs:{to:"/mission"}},[e("svg",{staticClass:"icon arrow",attrs:{"aria-hidden":"true"}},[e("use",{attrs:{"xlink:href":"#iconleft"}})])])],1),this._v(" "),e("div",{staticClass:"header-title v-center"},[this._v("\n    "+this._s(this.name)+"\n  ")]),this._v(" "),e("div",{staticClass:"header-right"})])},staticRenderFns:[]};var l=i("VU/8")(s,a,!1,function(t){i("u5nj")},"data-v-6285d02a",null).exports,n=(Object,{name:"detail",props:{mission:Object}}),c={render:function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",t._l(t.mission,function(e){return 1==e.id?i("div",{key:e.id},[i("div",{staticClass:"weui-cell textColor"},[t._v(t._s(e.content))]),t._v(" "),i("div",{staticClass:"weui-cells"},[i("div",{staticClass:"weui-cell"},[t._m(0,!0),t._v(" "),i("div",{staticClass:"weui-cell__bd cellMargin"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[t._v("\n            "+t._s(e.start_time)+"\n          ")])])]),t._v(" "),i("div",{staticClass:"weui-cell"},[t._m(1,!0),t._v(" "),i("div",{staticClass:"weui-cell__bd cellMargin"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[t._v("\n            "+t._s(e.end_time)+"\n          ")])])]),t._v(" "),i("div",{staticClass:"weui-cell"},[t._m(2,!0),t._v(" "),i("div",{staticClass:"weui-cell__bd cellMargin"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},t._l(e.participants,function(e){return i("span",{key:e.id,staticClass:"nameMargin"},[t._v(t._s(e.name))])}),0)])]),t._v(" "),e.file?i("div",{staticClass:"weui-cell"},[t._m(3,!0),t._v(" "),t._m(4,!0)]):t._e(),t._v(" "),t._m(5,!0),t._v(" "),i("div",{staticClass:"weui-cells"},t._l(e.gains,function(e){return i("div",{key:e.id,staticClass:"weui-cell"},[i("div",{staticClass:"weui-cell__bd"},[i("p",[t._v(t._s(e.name)+"发表"+t._s(e.type))])]),t._v(" "),i("div",{staticClass:"weui-cell__ft"},[t._v(t._s(e.uptime))])])}),0)])]):t._e()}),0)},staticRenderFns:[function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("开始时间")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("结束时间")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("参与任务学生")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("任务附件")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__bd cellMargin"},[e("div",{staticClass:"weui_cell_bd weui_cell_primary"},[e("button",{staticClass:"buttonShow"},[this._v("点击下载")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell"},[e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label v-center"},[this._v("成果")])])])}]};var r=i("VU/8")(n,c,!1,function(t){i("xbNc")},"data-v-3d9d1976",null).exports,_=i("mtWM"),u=i.n(_),d={name:"detail",components:{Header:l,Detail:r},data:function(){return{mission:[],name:""}},methods:{getMission:function(){sessionStorage.setItem("mission_id",this.$route.params.miId),u.a.get(this.url+"/Mission/id/"+this.$route.params.miId).then(this.getMissionSucc)},getMissionSucc:function(t){202===t.status?this.$router.push("/noProject"):t&&(t=t.data,this.mission=t,this.name=t.name)}},mounted:function(){this.getMission()}},v={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("header",{attrs:{name:this.name}}),this._v(" "),e("detail",{attrs:{mission:this.mission}})],1)},staticRenderFns:[]};var o=i("VU/8")(d,v,!1,function(t){i("YU7B")},"data-v-03dfe498",null);e.default=o.exports},YU7B:function(t,e){},u5nj:function(t,e){},xbNc:function(t,e){}});
//# sourceMappingURL=14.af083906053057dab07e.js.map