import{_ as x,d as D,u as T,e as V,f as y,r,o as l,c as u,a as c,g as a,h as t,t as i,i as d,F as B,j as S,w as I}from"./index-a5NbRkrj.js";const L={class:"content"},j={class:"view-img"},C=["src"],F={class:"details"},U=["href"],_="https://severj.top/img/",E={__name:"DetailView",setup(M){let s=D(null);const m=T();V(async()=>{s.value=m.params.id,await new Promise(e=>setTimeout(e,2e3)),y().commit("setLoading",!1)});const g=e=>`${_}background${e}.webp`,p=e=>`name${e}`,b=e=>`description${e}`,f=e=>e,h=(e,n)=>e+n,v=e=>`${_}background${e}.webp`,w=e=>e;return(e,n)=>{const N=r("a-page-header"),$=r("a-rate"),k=r("a-tag");return l(),u("div",null,[c(N,{style:{border:"1px solid rgb(235, 237, 240)"},title:"景点详情","sub-title":`ID: ${a(s)}`,onBack:n[0]||(n[0]=()=>e.$router.go(-1))},null,8,["sub-title"]),t("div",L,[t("div",j,[t("img",{src:g(Number(a(s))),alt:""},null,8,C)]),t("div",F,[t("h2",null,i(p(Number(a(s)))),1),t("p",null,"简介: "+i(b(Number(a(s)))),1),t("p",null,[d("评分："),c($,{value:f(4.5),"allow-half":"",disabled:""},null,8,["value"])]),t("p",null,[d("标签： "),(l(!0),u(B,null,S(w(3),o=>(l(),u("a",{key:o},[c(k,null,{default:I(()=>[t("a",{href:v(Number(o))},i(h(Number(a(s)),o)),9,U)]),_:2},1024)]))),128))])])])])}}},R=x(E,[["__scopeId","data-v-fdfa66b5"]]);export{R as default};
