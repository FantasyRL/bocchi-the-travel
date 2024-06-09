import{D as _,Q as w,Z as B,B as E,U as y,N,G as S,o as r,c as m,b as c,V as u,u as n,a3 as $,A as g,w as l,W as O,h as p,t as i,Y as P,a1 as V,a2 as z,l as v,a,F as j,f as b,g as D,a4 as F,r as k,j as d}from"./index-t4Ap-Vw9.js";/* empty css                                                  */import{C as H,H as Q,P as q,E as A}from"./PlayCircleOutlined-BhkrH4Zb.js";import{i as G,E as K}from"./index-C4NRWXjE.js";const U=_({name:"ElTimeline",setup(o,{slots:e}){const t=w("timeline");return B("timeline",e),()=>E("ul",{class:[t.b()]},[y(e,"default")])}}),W=N({timestamp:{type:String,default:""},hideTimestamp:{type:Boolean,default:!1},center:{type:Boolean,default:!1},placement:{type:String,values:["top","bottom"],default:"bottom"},type:{type:String,values:["primary","success","warning","danger","info"],default:""},color:{type:String,default:""},size:{type:String,values:["normal","large"],default:"normal"},icon:{type:G},hollow:{type:Boolean,default:!1}}),Y=_({name:"ElTimelineItem"}),Z=_({...Y,props:W,setup(o){const e=o,t=w("timeline-item"),f=S(()=>[t.e("node"),t.em("node",e.size||""),t.em("node",e.type||""),t.is("hollow",e.hollow)]);return(s,h)=>(r(),m("li",{class:u([n(t).b(),{[n(t).e("center")]:s.center}])},[c("div",{class:u(n(t).e("tail"))},null,2),s.$slots.dot?p("v-if",!0):(r(),m("div",{key:0,class:u(n(f)),style:$({backgroundColor:s.color})},[s.icon?(r(),g(n(K),{key:0,class:u(n(t).e("icon"))},{default:l(()=>[(r(),g(O(s.icon)))]),_:1},8,["class"])):p("v-if",!0)],6)),s.$slots.dot?(r(),m("div",{key:1,class:u(n(t).e("dot"))},[y(s.$slots,"dot")],2)):p("v-if",!0),c("div",{class:u(n(t).e("wrapper"))},[!s.hideTimestamp&&s.placement==="top"?(r(),m("div",{key:0,class:u([n(t).e("timestamp"),n(t).is("top")])},i(s.timestamp),3)):p("v-if",!0),c("div",{class:u(n(t).e("content"))},[y(s.$slots,"default")],2),!s.hideTimestamp&&s.placement==="bottom"?(r(),m("div",{key:1,class:u([n(t).e("timestamp"),n(t).is("bottom")])},i(s.timestamp),3)):p("v-if",!0)],2)],2))}});var T=P(Z,[["__file","timeline-item.vue"]]);const J=V(U,{TimelineItem:T}),L=z(T),M={class:"travels"},R=c("br",null,null,-1),X={key:0,class:"itinerary"},x=c("span",{class:"shadow"},null,-1),ee=c("span",{class:"edge"},null,-1),te=c("span",{class:"front text"}," 删除计划 ",-1),se=[x,ee,te],ne={key:1},ae={props:{},data(){return{id:1,party:{start_time:"2006-01-02",end_time:"2006-01-03"},info:{},partynull:!1}},methods:{deleteItinerary(){const o="/bocchi/party/itinerary/delete?itinerary_id="+this.id;b.get(o,{headers:{"access-token":this.access_token}}).then(e=>{console.log(e)}).catch(e=>{console.error(e)})},init(){const o="/bocchi/party/itinerary/info?itinerary_id="+this.id,e={};b.get(o,e).then(t=>{console.log(t),this.partynull=!0,this.info=t.data.itinerary,t.data.base.code==10007&&(this.partynull=0,console.log(this.partynull))}).catch(t=>{console.error(t)})}},mounted(){this.id=Number(this.$route.params.id),this.init(),this.access_token=D.get("access_token")},computed:{getIcon(){return o=>{switch(o){case 1:return A;case 2:return q;case 3:return Q;case 4:return H;case 5:return F}}},getType(){return o=>{switch(o){case 1:return"路线";case 2:return"活动";case 3:return"住宿";case 4:return"餐饮";case 5:return"其他";default:return"未知类型"}}}}},ue=Object.assign(ae,{__name:"itinerary",setup(o){return v(!1),v(!1),(e,t)=>{const f=k("a-page-header"),s=L,h=J,C=k("a-empty");return r(),m(j,null,[c("div",M,[a(f,{style:{border:"1px solid rgb(235, 237, 240)"},title:"计划详情",onBack:t[0]||(t[0]=()=>e.$router.go(-1))})]),R,e.partynull?(r(),m("div",X,[a(h,{style:{"max-width":"600px","margin-left":"10%"}},{default:l(()=>[a(s,null,{default:l(()=>[d("计划名:"+i(e.info.title),1)]),_:1}),a(s,null,{default:l(()=>[d("序列:"+i(e.info.sequence),1)]),_:1}),a(s,null,{default:l(()=>[d(" 创建者:"+i(e.info.founder_id),1)]),_:1}),a(s,null,{default:l(()=>[d(" 类型："+i(e.getType(e.info.action_type)),1)]),_:1}),a(s,null,{default:l(()=>[d(" 备注："+i(e.info.remark),1)]),_:1}),a(s,null,{default:l(()=>[d(" 地点："+i(e.info.rectangle),1)]),_:1}),a(s,null,{default:l(()=>[d(" 路线："+i(e.info.route_json),1)]),_:1}),a(s,null,{default:l(()=>[d(" 开始时间："+i(e.info.schedule_start_time),1)]),_:1}),a(s,null,{default:l(()=>[d(" 结束时间: "+i(e.info.schedule_end_time),1)]),_:1})]),_:1}),c("button",{onClick:t[1]||(t[1]=(...I)=>e.deleteItinerary&&e.deleteItinerary(...I))},se)])):p("",!0),e.partynull?p("",!0):(r(),m("div",ne,[a(C)]))],64)}}});export{ue as default};