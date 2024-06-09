import{D as _,Q as C,Z as E,B as N,U as h,N as $,G as S,o as u,c as p,b as d,V as m,u as l,a3 as z,A as k,w as s,W as O,h as f,t as r,Y as P,a1 as V,a2 as j,l as w,a as n,F as D,f as g,g as F,a4 as H,r as y,j as i}from"./index-BwwmACn5.js";/* empty css                                                  */import{C as Q,H as q,P as A,E as G}from"./PlayCircleOutlined-DJpXC4jI.js";import{i as K,E as U}from"./index-CT41R4hy.js";const W=_({name:"ElTimeline",setup(o,{slots:e}){const t=C("timeline");return E("timeline",e),()=>N("ul",{class:[t.b()]},[h(e,"default")])}}),Y=$({timestamp:{type:String,default:""},hideTimestamp:{type:Boolean,default:!1},center:{type:Boolean,default:!1},placement:{type:String,values:["top","bottom"],default:"bottom"},type:{type:String,values:["primary","success","warning","danger","info"],default:""},color:{type:String,default:""},size:{type:String,values:["normal","large"],default:"normal"},icon:{type:K},hollow:{type:Boolean,default:!1}}),Z=_({name:"ElTimelineItem"}),J=_({...Z,props:Y,setup(o){const e=o,t=C("timeline-item"),c=S(()=>[t.e("node"),t.em("node",e.size||""),t.em("node",e.type||""),t.is("hollow",e.hollow)]);return(a,v)=>(u(),p("li",{class:m([l(t).b(),{[l(t).e("center")]:a.center}])},[d("div",{class:m(l(t).e("tail"))},null,2),a.$slots.dot?f("v-if",!0):(u(),p("div",{key:0,class:m(l(c)),style:z({backgroundColor:a.color})},[a.icon?(u(),k(l(U),{key:0,class:m(l(t).e("icon"))},{default:s(()=>[(u(),k(O(a.icon)))]),_:1},8,["class"])):f("v-if",!0)],6)),a.$slots.dot?(u(),p("div",{key:1,class:m(l(t).e("dot"))},[h(a.$slots,"dot")],2)):f("v-if",!0),d("div",{class:m(l(t).e("wrapper"))},[!a.hideTimestamp&&a.placement==="top"?(u(),p("div",{key:0,class:m([l(t).e("timestamp"),l(t).is("top")])},r(a.timestamp),3)):f("v-if",!0),d("div",{class:m(l(t).e("content"))},[h(a.$slots,"default")],2),!a.hideTimestamp&&a.placement==="bottom"?(u(),p("div",{key:1,class:m([l(t).e("timestamp"),l(t).is("bottom")])},r(a.timestamp),3)):f("v-if",!0)],2)],2))}});var T=P(J,[["__file","timeline-item.vue"]]);const L=V(W,{TimelineItem:T}),M=j(T),R={class:"travels"},X=d("br",null,null,-1),x={key:0,class:"itinerary"},ee=["src"],te=d("span",{class:"shadow"},null,-1),ae=d("span",{class:"edge"},null,-1),oe=d("span",{class:"front text"}," 删除计划 ",-1),ne=[te,ae,oe],se={key:1},le={props:{},data(){return{datago:null,datadata:null,id:1,party:{start_time:"2006-01-02",end_time:"2006-01-03"},info:{},partynull:!1}},methods:{gogogo(){window.location.replace("//uri.amap.com/navigation?from="+this.datadata+",startpoint&to="+this.datago+",endpoint&via=&mode=car&policy=1&src=mypage&coordinate=gaode&callnative=0")},deleteItinerary(){const o="/bocchi/party/itinerary/delete?itinerary_id="+this.id;g.get(o,{headers:{"access-token":this.access_token}}).then(e=>{console.log(e)}).catch(e=>{console.error(e)})},init(){const o="/bocchi/party/itinerary/info?itinerary_id="+this.id,e={};g.get(o,e).then(t=>{console.log(t),this.partynull=!0,this.info=t.data.itinerary,t.data.base.code==10007&&(this.partynull=0,console.log(this.partynull))}).catch(t=>{console.error(t)})}},mounted(){this.id=Number(this.$route.params.id),this.init(),this.access_token=F.get("access_token")},computed:{getlocal(){return o=>{if(o!=="")return"https://restapi.amap.com/v3/staticmap?zoom=17&size=250*250&key=eae4d0491385d75b43d247afaef4247d&location="+o}},getroad(){return o=>{if(o!==void 0){const e=o.split(",")[0],t=o.split(",")[1];return g.get("https://restapi.amap.com/v3/geocode/geo?key=eae4d0491385d75b43d247afaef4247d&address="+e).then(c=>{console.log(c.data.geocodes[0].location),this.datadata=c.data.geocodes[0].location}),g.get("https://restapi.amap.com/v3/geocode/geo?key=eae4d0491385d75b43d247afaef4247d&address="+t).then(c=>{console.log(c.data.geocodes[0].location),this.datago=c.data.geocodes[0].location}),"从 "+e+" 到 "+t}}},getIcon(){return o=>{switch(o){case 1:return G;case 2:return A;case 3:return q;case 4:return Q;case 5:return H}}},getType(){return o=>{switch(o){case 1:return"路线";case 2:return"活动";case 3:return"住宿";case 4:return"餐饮";case 5:return"其他";default:return"未知类型"}}}}},ce=Object.assign(le,{__name:"itinerary",setup(o){return w(!1),w(!1),(e,t)=>{const c=y("a-page-header"),a=M,v=y("a-button"),I=L,B=y("a-empty");return u(),p(D,null,[d("div",R,[n(c,{style:{border:"1px solid rgb(235, 237, 240)"},title:"计划详情",onBack:t[0]||(t[0]=()=>e.$router.go(-1))})]),X,e.partynull?(u(),p("div",x,[n(I,{style:{"max-width":"600px","margin-left":"10%"}},{default:s(()=>[n(a,null,{default:s(()=>[i("计划名:"+r(e.info.title),1)]),_:1}),n(a,null,{default:s(()=>[i("序列:"+r(e.info.sequence),1)]),_:1}),n(a,null,{default:s(()=>[i(" 创建者:"+r(e.info.founder_id),1)]),_:1}),n(a,null,{default:s(()=>[i(" 类型："+r(e.getType(e.info.action_type)),1)]),_:1}),n(a,null,{default:s(()=>[i(" 备注："+r(e.info.remark),1)]),_:1}),n(a,null,{default:s(()=>[i(" 地点："+r(e.info.rectangle),1)]),_:1}),n(a,null,{default:s(()=>[d("div",null,[d("img",{src:e.getlocal(e.info.rectangle)},null,8,ee)])]),_:1}),n(a,null,{default:s(()=>[i(r(e.getroad(e.info.route_json)),1)]),_:1}),n(a,null,{default:s(()=>[i(" 路线： "),n(v,{type:"primary",onClick:t[1]||(t[1]=b=>e.gogogo())},{default:s(()=>[i("点击跳转到大地图")]),_:1})]),_:1}),n(a,null,{default:s(()=>[i(" 开始时间："+r(e.info.schedule_start_time),1)]),_:1}),n(a,null,{default:s(()=>[i(" 结束时间: "+r(e.info.schedule_end_time),1)]),_:1})]),_:1}),d("button",{onClick:t[2]||(t[2]=(...b)=>e.deleteItinerary&&e.deleteItinerary(...b))},ne)])):f("",!0),e.partynull?f("",!0):(u(),p("div",se,[n(B)]))],64)}}});export{ce as default};
