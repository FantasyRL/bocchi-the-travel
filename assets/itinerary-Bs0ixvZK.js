import{o as n,c as r,j as l,a as e,t as a,F as s,f as o}from"./index-B9MZhPqo.js";const u=e("br",null,null,-1),_=e("br",null,null,-1),c=e("br",null,null,-1),m={props:{},data(){return{id:1,data:{id:2,title:"1",founder_id:6,action_type:2,rectangle:"1",route_json:"1",remark:"1",sequence:0,schedule_start_time:"2006-01-02 15:40",schedule_end_time:"2006-01-02 15:40",party_id:5,is_merged:1}}},methods:{init(){const d="/bocchi/party/itinerary/info?itinerary_id="+this.id,t={};o.get(d,t).then(i=>{console.log(i),this.data=i.data.itinerary}).catch(i=>{console.error(i)})}},mounted(){this.id=Number(this.$route.params.id),this.init()}},p=Object.assign(m,{__name:"itinerary",setup(d){return(t,i)=>(n(),r(s,null,[l(" itinerary debug part: "),e("div",null,"itinerary id:"+a(t.id),1),e("div",null,"原始数据:"+a(t.data),1),u,_,c,e("div",null,"tittle:"+a(t.data.title),1),e("div",null,"founder_id:"+a(t.data.founder_id),1),e("div",null,"action_type:"+a(t.data.action_type),1),e("div",null,"rectangle:"+a(t.data.rectangle),1),e("div",null,"route_json:"+a(t.data.route_json),1),e("div",null,"remark:"+a(t.data.remark),1),e("div",null,"sequence:"+a(t.data.sequence),1),e("div",null,"schedule_start_time:"+a(t.data.schedule_start_time),1),e("div",null,"schedule_end_time:"+a(t.data.schedule_end_time),1),e("div",null,"party_id:"+a(t.data.party_id),1),e("div",null,"is_merged:"+a(t.data.is_merged),1)],64))}});export{p as default};
