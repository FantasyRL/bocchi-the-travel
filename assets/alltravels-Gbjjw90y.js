import{y as J,z as A,A as I,B as T,C as j,D as G,n as E,E as L,u as $,G as D,o as f,c as v,H as g,J as C,x as K,a as l,K as H,h as O,t as b,_ as k,r as M,b as q,F as V,q as z,w as N,p as Q,i as R}from"./index-Ph-8DMZH.js";const F="__epPropKey",U=s=>s,W=s=>A(s)&&!!s[F],X=(s,a)=>{if(!A(s)||W(s))return s;const{values:e,required:n,default:c,type:d,validator:_}=s,m={type:d,required:!!n,validator:e||_?u=>{let i=!1,y=[];if(e&&(y=Array.from(e),I(s,"default")&&y.push(c),i||(i=y.includes(u))),_&&(i||(i=_(u))),!i&&y.length>0){const P=[...new Set(y)].map(B=>JSON.stringify(B)).join(", ");T(`Invalid prop: validation failed${a?` for prop "${a}"`:""}. Expected one of [${P}], got value ${JSON.stringify(u)}.`)}return i}:void 0,[F]:!0};return I(s,"default")&&(m.default=c),m},Y=s=>J(Object.entries(s).map(([a,e])=>[a,X(e,a)])),Z=(s,a)=>(s.install=e=>{for(const n of[s,...Object.values({})])e.component(n.name,n)},s),S="el",x="is-",p=(s,a,e,n,c)=>{let d=`${s}-${a}`;return e&&(d+=`-${e}`),n&&(d+=`__${n}`),c&&(d+=`--${c}`),d},ee=Symbol("namespaceContextKey"),se=s=>{const a=j()?G(ee,E(S)):E(S);return L(()=>$(a)||S)},te=(s,a)=>{const e=se();return{namespace:e,b:(t="")=>p(e.value,s,t,"",""),e:t=>t?p(e.value,s,"",t,""):"",m:t=>t?p(e.value,s,"","",t):"",be:(t,r)=>t&&r?p(e.value,s,t,r,""):"",em:(t,r)=>t&&r?p(e.value,s,"",t,r):"",bm:(t,r)=>t&&r?p(e.value,s,t,"",r):"",bem:(t,r,o)=>t&&r&&o?p(e.value,s,t,r,o):"",is:(t,...r)=>{const o=r.length>=1?r[0]:!0;return t&&o?`${x}${t}`:""},cssVar:t=>{const r={};for(const o in t)t[o]&&(r[`--${e.value}-${o}`]=t[o]);return r},cssVarName:t=>`--${e.value}-${t}`,cssVarBlock:t=>{const r={};for(const o in t)t[o]&&(r[`--${e.value}-${s}-${o}`]=t[o]);return r},cssVarBlockName:t=>`--${e.value}-${s}-${t}`}};var ae=(s,a)=>{const e=s.__vccOpts||s;for(const[n,c]of a)e[n]=c;return e};const re=Y({header:{type:String,default:""},footer:{type:String,default:""},bodyStyle:{type:U([String,Object,Array]),default:""},bodyClass:String,shadow:{type:String,values:["always","hover","never"],default:"always"}}),ne=D({name:"ElCard"}),oe=D({...ne,props:re,setup(s){const a=te("card");return(e,n)=>(f(),v("div",{class:g([$(a).b(),$(a).is(`${e.shadow}-shadow`)])},[e.$slots.header||e.header?(f(),v("div",{key:0,class:g($(a).e("header"))},[C(e.$slots,"header",{},()=>[O(b(e.header),1)])],2)):K("v-if",!0),l("div",{class:g([$(a).e("body"),e.bodyClass]),style:H(e.bodyStyle)},[C(e.$slots,"default")],6),e.$slots.footer||e.footer?(f(),v("div",{key:1,class:g($(a).e("footer"))},[C(e.$slots,"footer",{},()=>[O(b(e.footer),1)])],2)):K("v-if",!0)],2))}});var le=ae(oe,[["__file","card.vue"]]);const ce=Z(le),de={data(){return{trnumber:10}}},h=s=>(Q("data-v-9bf6467d"),s=s(),R(),s),ue={class:"travels"},ie={class:"center"},pe=h(()=>l("br",null,null,-1)),fe={class:"card-header"},ve=h(()=>l("br",null,null,-1)),_e=h(()=>l("br",null,null,-1)),me=h(()=>l("div",{class:"null"},null,-1)),ye=h(()=>l("br",null,null,-1));function $e(s,a,e,n,c,d){const _=M("a-page-header"),w=ce;return f(),v(V,null,[l("div",ue,[q(_,{style:{border:"1px solid rgb(235, 237, 240)"},title:"所有行程",onBack:a[0]||(a[0]=()=>s.$router.go(-1))})]),l("div",ie,[(f(!0),v(V,null,z(c.trnumber,m=>(f(),v("div",{key:m},[pe,q(w,{style:{"min-width":"90vw"}},{header:N(()=>[l("div",fe,[l("span",null,"Card name "+b(m),1)])]),footer:N(()=>[O("Footer content")]),default:N(()=>[(f(),v(V,null,z(4,u=>l("p",{key:u,class:"text item"},b("List item "+u),1)),64))]),_:2},1024),ve]))),128)),_e,me,ye])],64)}const be=k(de,[["render",$e],["__scopeId","data-v-9bf6467d"]]);export{be as default};
