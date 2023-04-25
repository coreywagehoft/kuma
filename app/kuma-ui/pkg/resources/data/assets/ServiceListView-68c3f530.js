import{d as E,u as C,r,v as L,e as w,w as T,o as k,g as R,f as V}from"./index-834ac640.js";import{S as q}from"./ServiceSummary-79cad6f3.js";import{C as z}from"./ContentWrapper-be0c9d71.js";import{D as B}from"./DataOverview-9d4155bf.js";import{u as M}from"./index-0e952743.js";import{Q as P}from"./QueryParameter-70743f73.js";import"./kongponents.es-131ffe48.js";import"./DefinitionListItem-e9b95b5e.js";import"./ErrorBlock-42ddf946.js";import"./_plugin-vue_export-helper-c27b6911.js";import"./LoadingBlock.vue_vue_type_script_setup_true_lang-4325e9aa.js";import"./StatusBadge-1d0340ff.js";import"./TagList-7389a145.js";import"./store-bb95959d.js";import"./YamlView.vue_vue_type_script_setup_true_lang-05c499f0.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-82e53469.js";import"./toYaml-4e00099e.js";import"./datadogLogEvents-302eea7b.js";const re=E({__name:"ServiceListView",props:{selectedServiceName:{type:String,required:!1,default:null},offset:{type:Number,required:!1,default:0}},setup(_){const c=_,m=M(),A=[{label:"Service",key:"name"},{label:"Type",key:"serviceType"},{label:"Address",key:"addressPort"},{label:"Status",key:"status"},{label:"DP proxies (online / total)",key:"dpProxiesStatus"}],S=50,N={title:"No Data",message:"There are no service insights present."},u=C(),p=r(!0),v=r(null),x=r(null),h=r(c.offset),o=r(null),b=r(null),l=r({headers:A,data:[]});L(()=>u.params.mesh,function(){u.name==="service-list-view"&&d(0)}),d(c.offset);async function d(e){h.value=e,P.set("offset",e>0?e:null),p.value=!0,v.value=null;const t=u.params.mesh,s=S;try{const{items:a,next:f}=await m.getAllServiceInsightsFromMesh({mesh:t},{size:s,offset:e});if(x.value=f,Array.isArray(a)&&a.length>0){a.sort((n,i)=>n.name>i.name?1:n.name<i.name?-1:0),l.value.data=a.map(n=>D(n));const y=c.selectedServiceName??a[0].name;await g({name:y,mesh:t})}else l.value.data=[]}catch(a){a instanceof Error?v.value=a:console.error(a)}finally{p.value=!1}}function D(e){const t={name:"service-detail-view",params:{mesh:e.mesh,service:e.name}},s={name:"mesh-detail-view",params:{mesh:e.mesh}};let a="—";if(e.dataplanes){const{online:n=0,total:i=0}=e.dataplanes;a=`${n} / ${i}`}const f=e.addressPort,y=e.serviceType??"internal";return{...e,serviceType:y,nameRoute:t,meshRoute:s,dpProxiesStatus:a,addressPort:f}}async function g({mesh:e,name:t}){o.value=await m.getServiceInsight({mesh:e,name:t}),o.value.serviceType==="external"&&(b.value=await m.getExternalServiceByServiceInsightName(e,t)),P.set("service",t)}return(e,t)=>(k(),w(z,null,{content:T(()=>{var s;return[R(B,{"selected-entity-name":(s=o.value)==null?void 0:s.name,"page-size":S,error:v.value,"is-loading":p.value,"empty-state":N,"table-data":l.value,"table-data-is-empty":l.value.data.length===0,next:x.value,"page-offset":h.value,onTableAction:g,onLoadData:d},null,8,["selected-entity-name","error","is-loading","table-data","table-data-is-empty","next","page-offset"])]}),sidebar:T(()=>[o.value!==null?(k(),w(q,{key:0,service:o.value,"external-service":b.value},null,8,["service","external-service"])):V("",!0)]),_:1}))}});export{re as default};
