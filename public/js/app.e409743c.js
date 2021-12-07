(function(t){function e(e){for(var a,r,l=e[0],s=e[1],u=e[2],d=0,b=[];d<l.length;d++)r=l[d],Object.prototype.hasOwnProperty.call(c,r)&&c[r]&&b.push(c[r][0]),c[r]=0;for(a in s)Object.prototype.hasOwnProperty.call(s,a)&&(t[a]=s[a]);i&&i(e);while(b.length)b.shift()();return o.push.apply(o,u||[]),n()}function n(){for(var t,e=0;e<o.length;e++){for(var n=o[e],a=!0,l=1;l<n.length;l++){var s=n[l];0!==c[s]&&(a=!1)}a&&(o.splice(e--,1),t=r(r.s=n[0]))}return t}var a={},c={app:0},o=[];function r(e){if(a[e])return a[e].exports;var n=a[e]={i:e,l:!1,exports:{}};return t[e].call(n.exports,n,n.exports,r),n.l=!0,n.exports}r.m=t,r.c=a,r.d=function(t,e,n){r.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:n})},r.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},r.t=function(t,e){if(1&e&&(t=r(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var a in t)r.d(n,a,function(e){return t[e]}.bind(null,a));return n},r.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return r.d(e,"a",e),e},r.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},r.p="/";var l=window["webpackJsonp"]=window["webpackJsonp"]||[],s=l.push.bind(l);l.push=e,l=l.slice();for(var u=0;u<l.length;u++)e(l[u]);var i=s;o.push([0,"chunk-vendors"]),n()})({0:function(t,e,n){t.exports=n("56d7")},"387a":function(t,e,n){},"56d7":function(t,e,n){"use strict";n.r(e);n("e260"),n("e6cf"),n("cca6"),n("a79d");var a=n("7a23"),c={class:"shadow-2xl p-8 m-8"};function o(t,e,n,o,r,l){var s=Object(a["j"])("Tasker");return Object(a["h"])(),Object(a["c"])("div",c,[Object(a["f"])(s,{msg:"Welcome to Go Tasker!"})])}var r=n("aea3"),l=n.n(r),s={class:"text-6xl my-12"},u=Object(a["d"])("img",{src:l.a,style:{display:"inline-flex"},width:"50",height:"50"},null,-1),i=Object(a["d"])("label",{class:"uppercase italic",for:"task"},"New Task: ",-1),d={class:"w-full h-full"},b={class:"my-0 mx-auto"},f=Object(a["d"])("tr",null,[Object(a["d"])("th",{class:"table-header"},"Task"),Object(a["d"])("th",{class:"table-header"},"Status"),Object(a["d"])("th",{class:"table-header text-left"},"Mark Complete"),Object(a["d"])("th",{class:"table-header text-left"},"Remove")],-1),p=["task"],h={class:"table-data"},j={class:"table-data"},O=["onClick","disabled"],k=Object(a["d"])("svg",{class:"w-6 h-6",fill:"none",stroke:"currentColor",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},[Object(a["d"])("path",{"stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"2",d:"M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"})],-1),m=[k],g={class:"table-data"},v=["onClick"],w=Object(a["d"])("svg",{class:"w-6 h-6",fill:"none",stroke:"currentColor",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},[Object(a["d"])("path",{"stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"2",d:"M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"})],-1),y=[w];function x(t,e,n,c,o,r){return Object(a["h"])(),Object(a["c"])(a["a"],null,[Object(a["d"])("h1",s,[u,Object(a["e"])(Object(a["k"])(n.msg),1)]),Object(a["d"])("div",null,[i,Object(a["m"])(Object(a["d"])("input",{class:"bg-gray-200","onUpdate:modelValue":e[0]||(e[0]=function(t){return o.task=t})},null,512),[[a["l"],o.task]]),Object(a["d"])("button",{onClick:e[1]||(e[1]=function(){return r.handleSubmit&&r.handleSubmit.apply(r,arguments)}),class:"bg-blue-400 w-24 my-12"},"Add")]),Object(a["d"])("div",d,[Object(a["d"])("table",b,[f,(Object(a["h"])(!0),Object(a["c"])(a["a"],null,Object(a["i"])(o.tasks,(function(t){return Object(a["h"])(),Object(a["c"])("tr",{task:t,key:t.ID},[Object(a["d"])("td",h,Object(a["k"])(t.Text),1),Object(a["d"])("td",{class:Object(a["g"])(["table-data",{"bg-green-300":"Complete"===t.Status,"bg-yellow-300":"Pending"===t.Status}])},Object(a["k"])(t.Status),3),Object(a["d"])("td",j,[Object(a["d"])("button",{onClick:function(e){return r.markComplete(t.Text)},disabled:"Complete"===t.Completed},m,8,O)]),Object(a["d"])("td",g,[Object(a["d"])("button",{onClick:function(e){return r.deleteTask(t.Text)}},y,8,v)])],8,p)})),128))])])],64)}n("d3b7"),n("159b");var T=n("bc3a"),C=n.n(T),S={name:"Tasker",props:{msg:String},data:function(){return{task:"",tasks:[]}},created:function(){this.fetchData()},methods:{fetchData:function(){var t=this;C.a.get("/api/index").then((function(e){e.data.forEach((function(t){!0===t.Completed?t.Status="Complete":t.Status="Pending"})),t.tasks=e.data})).catch((function(t){return console.err(t)}))},handleSubmit:function(){var t=this,e=this.fetchData;C.a.post("/api/add",{task:this.task}).then((function(){e(),t.task=""})).then((function(){return console.log("Task successfully added.")})).catch((function(t){return console.log(t)}))},markComplete:function(t){var e=this.fetchData;C.a.post("/api/done",{task:t}).then((function(){return e()})).then((function(){return console.log("Task successfully marked complete.")})).catch((function(t){return console.log(t)}))},deleteTask:function(t){var e=this.fetchData;C.a.post("/api/rm",{task:t}).then((function(){return e()})).then((function(){return console.log("Task successfully deleted.")})).catch((function(t){return console.log(t)}))}}},M=(n("9899"),n("6b0d")),P=n.n(M);const D=P()(S,[["render",x]]);var _=D,A={name:"App",components:{Tasker:_}};n("7b85");const V=P()(A,[["render",o]]);var B=V;n("a766");Object(a["b"])(B).mount("#app")},"7b85":function(t,e,n){"use strict";n("387a")},9899:function(t,e,n){"use strict";n("98a0")},"98a0":function(t,e,n){},a766:function(t,e,n){},aea3:function(t,e,n){t.exports=n.p+"img/gopher-logo.f287df43.png"}});
//# sourceMappingURL=app.e409743c.js.map