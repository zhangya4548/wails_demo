import{_ as l}from"./index.js";import{a as _,o as d,b as p,d as t,t as e,e as a,f as i,g as n}from"./vendor.js";var r="/assets/comeon.gif";const h={setup(){return{comeonGif:r}}},m={class:"about"},u={class:"title"},f={class:"content"},b={class:"comeon"},g=["src"],v={class:"info"},k={class:"info-item"},$={class:"name"},w=n("https://github.com/misitebao/wails-template-vue"),y={class:"info-item"},B={class:"name"},N=n("https://github.com/wailsapp/wails"),V={class:"info-item"},A={class:"name"},C={class:"thank"};function G(s,L,O,c,j,D){const o=_("OpenLink");return d(),p("div",m,[t("div",u,e(s.$t("aboutpage.title")),1),t("div",f,[t("div",b,[t("img",{src:c.comeonGif,alt:""},null,8,g)]),t("ul",v,[t("li",k,[t("div",$,e(s.$t("aboutpage.project-repository")),1),a(o,{class:"link",href:"https://github.com/misitebao/wails-template-vue"},{default:i(()=>[w]),_:1})]),t("li",y,[t("div",B,e(s.$t("aboutpage.wails-repository")),1),a(o,{class:"link",href:"https://github.com/wailsapp/wails"},{default:i(()=>[N]),_:1})]),t("li",V,[t("div",A,e(s.$t("aboutpage.author")),1),a(o,{class:"link",href:"https://github.com/misitebao"},{default:i(()=>[n(e(s.$t("aboutpage.misitebao")),1)]),_:1})])])]),t("div",C,e(s.$t("aboutpage.thanks")),1)])}var T=l(h,[["render",G]]);export{T as default};