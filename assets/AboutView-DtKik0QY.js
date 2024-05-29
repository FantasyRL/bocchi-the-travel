import{_ as p,r as g,o as v,c as m,a as t,b as u,t as n,w as f,T as k,d as l,v as d,e as _,F as w,f as h,g as i,p as b,h as C}from"./index-DrSOrVwU.js";const r=s=>(b("data-v-2745259f"),s=s(),C(),s),F={class:"about"},y={class:"settings"},S=r(()=>t("svg",{xmlns:"http://www.w3.org/2000/svg",width:"20",viewBox:"0 0 20 20",height:"20",fill:"none",class:"svg-icon"},[t("g",{"stroke-width":"1.5","stroke-linecap":"round",stroke:"#5d41de"},[t("circle",{r:"2.5",cy:"10",cx:"10"}),t("path",{"fill-rule":"evenodd",d:"m8.39079 2.80235c.53842-1.51424 2.67991-1.51424 3.21831-.00001.3392.95358 1.4284 1.40477 2.3425.97027 1.4514-.68995 2.9657.82427 2.2758 2.27575-.4345.91407.0166 2.00334.9702 2.34248 1.5143.53842 1.5143 2.67996 0 3.21836-.9536.3391-1.4047 1.4284-.9702 2.3425.6899 1.4514-.8244 2.9656-2.2758 2.2757-.9141-.4345-2.0033.0167-2.3425.9703-.5384 1.5142-2.67989 1.5142-3.21831 0-.33914-.9536-1.4284-1.4048-2.34247-.9703-1.45148.6899-2.96571-.8243-2.27575-2.2757.43449-.9141-.01669-2.0034-.97028-2.3425-1.51422-.5384-1.51422-2.67994.00001-3.21836.95358-.33914 1.40476-1.42841.97027-2.34248-.68996-1.45148.82427-2.9657 2.27575-2.27575.91407.4345 2.00333-.01669 2.34247-.97026z","clip-rule":"evenodd"})])],-1)),A={class:"lable"},B={class:"info"},I=["src"],D={class:"name"},V={class:"shejiao"},j=r(()=>t("div",null,null,-1)),T={class:"setpage"},N={blur:!0},z=r(()=>t("text",null,"设置界面",-1)),E={class:"setpage-flur"},M={class:"avatarpage"},O=r(()=>t("div",{class:"bg"},null,-1)),$={data(){return{setname:"设置",setcount:3,setshow:0,id:1,name:"未登录",avatar:"./src/assets/ava.png",mail:"",signature:"这个人很懒，什么都没写。",access_token:"",refresh_token:"",avatarFile:null,url:"",showavatarpage:0}},methods:{modavatarpage(){this.showavatarpage=!this.showavatarpage,console.log(this.showavatarpage)},onFileChange(s){this.avatarFile=s.target.files[0]},putavatar(){if(!this.avatarFile){alert("请选择一个文件");return}let s=new FormData;s.append("avatar",this.avatarFile),h.put(this.url+"/bocchi/user/avatar/upload",s,{headers:{"Content-Type":"multipart/form-data","access-token":this.access_token,Accept:"*/*"}}).then(e=>{console.log(e)}).catch(e=>{console.error(e)})},settings(){this.setshow=!this.setshow,this.setcount++},refreshtoken(){h.get(this.url+"/bocchi/access_token/get",{headers:{"refresh-token":this.refresh_token,Accept:"*/*"}}).then(s=>{console.log(s),console.log(s.data.access_token),this.access_token=s.data.access_token,i.remove("access_token"),this.cookiesSet=i.set("access_token",s.data.access_token,{expires:1})}).catch(s=>{console.error(s)})},logout(){i.remove("id"),i.remove("access_token"),i.remove("refresh_token"),window.location.href="/"},init(){h.get(this.url+"/bocchi/user/info?user_id="+this.id).then(s=>{console.log(s),this.name=s.data.user.name,this.avatar=s.data.user.avatar,this.mail=s.data.user.email,this.signature=s.data.user.signature,console.log(this.avatar),this.avatar===""&&(this.avatar="//xiey.work/640.jpg"),this.signature===""&&(this.signature="这个人很懒，什么都没写。")}).catch(s=>{console.error(s)})}},mounted(){this.id=i.get("id"),this.access_token=i.get("access_token"),this.refresh_token=i.get("refresh_token"),this.id&&this.init()},components:{},computed:{},watch:{setcount(s,e){s%2==1&&(this.setname="设置"),s%2!=1&&(this.setname="关闭")}},props:{},emits:{},setup(){}},q=Object.assign($,{__name:"AboutView",setup(s){return(e,a)=>{const c=g("a-page-header");return v(),m(w,null,[t("div",F,[u(c,{style:{border:"1px solid rgb(235, 237, 240)"},title:"个人中心","sub-title":"副标题",onBack:a[0]||(a[0]=()=>e.$router.go(-1))}),t("div",y,[t("button",{class:"button",onClick:a[1]||(a[1]=(...o)=>e.settings&&e.settings(...o))},[S,t("span",A,n(e.setname),1)])]),t("div",B,[t("div",{class:"touxiang",onClick:a[2]||(a[2]=(...o)=>e.modavatarpage&&e.modavatarpage(...o))},[t("div",null,[t("img",{src:e.avatar,class:"touxiangimg rounded-image"},null,8,I),t("div",null,[t("text",D,n(e.name),1)])])]),t("div",V,[t("text",null,"uid:"+n(e.id),1),t("text",null,"邮箱:"+n(e.mail),1),t("text",null,"签名:"+n(e.signature),1)])]),j,u(k,{name:"fade"},{default:f(()=>[l(t("div",T,[t("mdui-list",N,[z,t("mdui-list-item",{onClick:a[3]||(a[3]=(...o)=>e.refreshtoken&&e.refreshtoken(...o))},"刷新令牌"),t("mdui-list-item",{onClick:a[4]||(a[4]=(...o)=>e.logout&&e.logout(...o))},"退出登陆")])],512),[[d,e.setshow]])]),_:1}),l(t("div",E,null,512),[[d,e.setshow]]),l(t("div",M,[t("form",{onSubmit:a[7]||(a[7]=_((...o)=>e.uploadAvatar&&e.uploadAvatar(...o),["prevent"]))},[t("input",{type:"file",ref:"avatarInput",onChange:a[5]||(a[5]=(...o)=>e.onFileChange&&e.onFileChange(...o))},null,544),t("button",{type:"submit",onClick:a[6]||(a[6]=(...o)=>e.putavatar&&e.putavatar(...o))},"上传头像")],32)],512),[[d,1]])]),O],64)}}}),H=p(q,[["__scopeId","data-v-2745259f"]]);export{H as default};
