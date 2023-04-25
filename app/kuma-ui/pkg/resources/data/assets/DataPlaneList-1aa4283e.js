var Me=Object.defineProperty;var qe=(t,s,a)=>s in t?Me(t,s,{enumerable:!0,configurable:!0,writable:!0,value:a}):t[s]=a;var ee=(t,s,a)=>(qe(t,typeof s!="symbol"?s+"":s,a),a);import{d as de,r as O,c as S,v as Q,o as c,j as _,i as y,h as l,g as T,b as p,K as Re,t as P,a0 as Ee,a1 as je,Z as De,F as N,q as K,f as V,k as ze,a2 as Be,p as ce,m as pe,M as xe,Y as Ce,L as Oe,N as He,x as Qe,e as L,w as D,u as Ge,C as Pe,H as Se,Q as Ye,R as Ze,W as Je,X as We,U as Xe,V as et,a3 as tt,a4 as at}from"./index-834ac640.js";import{d as te,x as nt,k as st,D as lt,M as Ae}from"./kongponents.es-131ffe48.js";import{C as ot}from"./ContentWrapper-be0c9d71.js";import{D as it}from"./DataOverview-9d4155bf.js";import{D as F,a as ue,_ as rt}from"./DefinitionListItem-e9b95b5e.js";import{_ as fe}from"./_plugin-vue_export-helper-c27b6911.js";import{S as ut}from"./StatusBadge-1d0340ff.js";import{T as dt}from"./TagList-7389a145.js";import{_ as ct}from"./YamlView.vue_vue_type_script_setup_true_lang-05c499f0.js";import{u as pt}from"./store-bb95959d.js";import{d as ft}from"./datadogLogEvents-302eea7b.js";import{Q as z}from"./QueryParameter-70743f73.js";const Ie=[{key:"status",label:"Status"},{key:"name",label:"DPP"},{key:"type",label:"Type"},{key:"service",label:"Service"},{key:"protocol",label:"Protocol"},{key:"zone",label:"Zone"},{key:"lastConnected",label:"Last Connected"},{key:"lastUpdated",label:"Last Updated"},{key:"totalUpdates",label:"Total Updates"},{key:"dpVersion",label:"Kuma DP version"},{key:"envoyVersion",label:"Envoy version"},{key:"details",label:"Details",hideLabel:!0}],mt=["name","details"],vt=Ie.filter(t=>!mt.includes(t.key)).map(t=>({tableHeaderKey:t.key,label:t.label,isChecked:!1})),Ue=["status","name","type","service","protocol","zone","lastUpdated","dpVersion","details"];function gt(t,s=Ue){return Ie.filter(a=>s.includes(a.key)?t?!0:a.key!=="zone":!1)}function yt(t,s,a){return Math.max(s,Math.min(t,a))}const ht=["ControlLeft","ControlRight","ShiftLeft","ShiftRight","AltLeft"];class bt{constructor(s,a){ee(this,"commands");ee(this,"keyMap");ee(this,"boundTriggerShortcuts");this.commands=a,this.keyMap=Object.fromEntries(Object.entries(s).map(([w,u])=>[w.toLowerCase(),u])),this.boundTriggerShortcuts=this.triggerShortcuts.bind(this)}registerListener(){document.addEventListener("keydown",this.boundTriggerShortcuts)}unRegisterListener(){document.removeEventListener("keydown",this.boundTriggerShortcuts)}triggerShortcuts(s){_t(s,this.keyMap,this.commands)}}function _t(t,s,a){const w=kt(t.code),u=[t.ctrlKey?"ctrl":"",t.shiftKey?"shift":"",t.altKey?"alt":"",w].filter(k=>k!=="").join("+"),f=s[u];if(!f)return;const b=a[f];b.isAllowedContext&&!b.isAllowedContext(t)||(b.shouldPreventDefaultAction&&t.preventDefault(),!(b.isDisabled&&b.isDisabled())&&b.trigger(t))}function kt(t){return ht.includes(t)?"":t.replace(/^Key/,"").toLowerCase()}function wt(t,s){const a=" "+t,w=a.matchAll(/ ([-\s\w]+):\s*/g),u=[];for(const f of Array.from(w)){if(f.index===void 0)continue;const b=Tt(f[1]);if(s.length>0&&!s.includes(b))throw new Error(`Unknown field “${b}”. Known fields: ${s.join(", ")}`);const k=f.index+f[0].length,o=a.substring(k);let m;if(/^\s*["']/.test(o)){const d=o.match(/['"](.*?)['"]/);if(d!==null)m=d[1];else throw new Error(`Quote mismatch for field “${b}”.`)}else{const d=o.indexOf(" "),g=d===-1?o.length:d;m=o.substring(0,g)}m!==""&&u.push([b,m])}return u}function Tt(t){return t.trim().replace(/\s+/g,"-").replace(/-[a-z]/g,(s,a)=>a===0?s:s.substring(1).toUpperCase())}const Ve=t=>(ce("data-v-715248af"),t=t(),pe(),t),Dt=Ve(()=>y("span",{class:"visually-hidden"},"Focus filter",-1)),Ct=["for"],Pt=["id","placeholder"],St={key:0,class:"k-suggestion-box","data-testid":"k-filter-bar-suggestion-box"},At={class:"k-suggestion-list"},Et={key:0,class:"k-filter-bar-error"},xt={key:0},Ot=["title","data-filter-field"],It={class:"visually-hidden"},Ut=Ve(()=>y("span",{class:"visually-hidden"},"Clear query",-1)),Vt=de({__name:"KFilterBar",props:{id:{type:String,required:!0},fields:{type:Object,required:!0},placeholder:{type:String,required:!1,default:null},query:{type:String,required:!1,default:""}},emits:["fields-change"],setup(t,{emit:s}){const a=t,w=O(null),u=O(null),f=O(a.query),b=O([]),k=O(null),o=O(!1),m=O(-1),h=S(()=>Object.keys(a.fields)),d=S(()=>Object.entries(a.fields).slice(0,5).map(([e,r])=>({fieldName:e,...r}))),g=S(()=>h.value.length>0?`Filter by ${h.value.join(", ")}`:"Filter"),A=S(()=>a.placeholder??g.value);Q(()=>b.value,function(e,r){n(e,r)||(k.value=null,s("fields-change",{fields:e,query:f.value}))}),Q(()=>f.value,function(){f.value===""&&(k.value=null),o.value=!0});const E={Enter:"submitQuery",Escape:"closeSuggestionBox",ArrowDown:"jumpToNextSuggestion",ArrowUp:"jumpToPreviousSuggestion"},I={submitQuery:{trigger:Y,isAllowedContext(e){return u.value!==null&&e.composedPath().includes(u.value)},shouldPreventDefaultAction:!0},jumpToNextSuggestion:{trigger:ae,isAllowedContext(e){return u.value!==null&&e.composedPath().includes(u.value)},shouldPreventDefaultAction:!0},jumpToPreviousSuggestion:{trigger:ne,isAllowedContext(e){return u.value!==null&&e.composedPath().includes(u.value)},shouldPreventDefaultAction:!0},closeSuggestionBox:{trigger:H,isAllowedContext(e){return w.value!==null&&e.composedPath().includes(w.value)}}};function $(){const e=new bt(E,I);ze(function(){e.registerListener()}),Be(function(){e.unRegisterListener()}),q(f.value)}$();function G(e){const r=e.target;q(r.value)}function Y(){if(u.value instanceof HTMLInputElement)if(m.value===-1)q(u.value.value),o.value=!1;else{const e=d.value[m.value].fieldName;e&&Z(u.value,e)}}function ae(){M(1)}function ne(){M(-1)}function M(e){m.value=yt(m.value+e,-1,d.value.length-1)}function se(){u.value instanceof HTMLInputElement&&u.value.focus()}function le(e){const i=e.currentTarget.getAttribute("data-filter-field");i&&u.value instanceof HTMLInputElement&&Z(u.value,i)}function Z(e,r){const i=f.value===""||f.value.endsWith(" ")?"":" ";f.value+=i+r+":",e.focus(),m.value=-1}function J(){f.value="",u.value instanceof HTMLInputElement&&(u.value.value="",u.value.focus(),q(""))}function B(e){e.relatedTarget===null&&H(),w.value instanceof HTMLElement&&e.relatedTarget instanceof Node&&!w.value.contains(e.relatedTarget)&&H()}function H(){o.value=!1}function q(e){k.value=null;try{const r=wt(e,h.value);r.sort((i,C)=>i[0].localeCompare(C[0])),b.value=r}catch(r){if(r instanceof Error)k.value=r,o.value=!0;else throw r}}function n(e,r){return JSON.stringify(e)===JSON.stringify(r)}return(e,r)=>(c(),_("div",{ref_key:"filterBar",ref:w,class:"k-filter-bar","data-testid":"k-filter-bar"},[y("button",{class:"k-focus-filter-input-button",title:"Focus filter",type:"button","data-testid":"k-filter-bar-focus-filter-input-button",onClick:se},[Dt,l(),T(p(te),{"aria-hidden":"true",class:"k-filter-icon",color:"var(--grey-400)","data-testid":"k-filter-bar-filter-icon","hide-title":"",icon:"filter",size:"20"})]),l(),y("label",{for:`${a.id}-filter-bar-input`,class:"visually-hidden"},[Re(e.$slots,"default",{},()=>[l(P(p(g)),1)],!0)],8,Ct),l(),Ee(y("input",{id:`${a.id}-filter-bar-input`,ref_key:"filterInput",ref:u,"onUpdate:modelValue":r[0]||(r[0]=i=>f.value=i),class:"k-filter-bar-input",type:"text",placeholder:p(A),"data-testid":"k-filter-bar-filter-input",onFocus:r[1]||(r[1]=i=>o.value=!0),onBlur:B,onChange:G},null,40,Pt),[[je,f.value]]),l(),o.value?(c(),_("div",St,[y("div",At,[k.value!==null?(c(),_("p",Et,P(k.value.message),1)):(c(),_("button",{key:1,class:De(["k-submit-query-button",{"k-submit-query-button-is-selected":m.value===-1}]),title:"Submit query",type:"button","data-testid":"k-filter-bar-submit-query-button",onClick:Y},`
          Submit `+P(f.value),3)),l(),(c(!0),_(N,null,K(p(d),(i,C)=>(c(),_("div",{key:`${a.id}-${C}`,class:De(["k-suggestion-list-item",{"k-suggestion-list-item-is-selected":m.value===C}])},[y("b",null,P(i.fieldName),1),i.description!==""?(c(),_("span",xt,": "+P(i.description),1)):V("",!0),l(),y("button",{class:"k-apply-suggestion-button",title:`Add ${i.fieldName}:`,type:"button","data-filter-field":i.fieldName,"data-testid":"k-filter-bar-apply-suggestion-button",onClick:le},[y("span",It,"Add "+P(i.fieldName)+":",1),l(),T(p(te),{"aria-hidden":"true",color:"currentColor","hide-title":"",icon:"chevronRight",size:"16"})],8,Ot)],2))),128))])])):V("",!0),l(),f.value!==""?(c(),_("button",{key:1,class:"k-clear-query-button",title:"Clear query",type:"button","data-testid":"k-filter-bar-clear-query-button",onClick:J},[Ut,l(),T(p(te),{"aria-hidden":"true",color:"currentColor",icon:"clear","hide-title":"",size:"20"})])):V("",!0)],512))}});const $t=fe(Vt,[["__scopeId","data-v-715248af"]]),$e=t=>(ce("data-v-3fa16dcc"),t=t(),pe(),t),Lt={class:"entity-section-list"},Nt={class:"entity-title","data-testid":"data-plane-proxy-title"},Ft={class:"mt-2 heading-with-icon"},Kt={key:0},Mt=$e(()=>y("h4",null,"Insights",-1)),qt={class:"block-list"},Rt={key:0,class:"mt-2"},jt=$e(()=>y("summary",null,`
                  Responses (acknowledged / sent)
                `,-1)),zt={class:"config-section"},Bt=de({__name:"DataPlaneEntitySummary",props:{dataPlaneOverview:{type:Object,required:!0}},setup(t){const s=t,a=S(()=>{const{name:o,mesh:m,dataplane:h}=s.dataPlaneOverview;return{type:"Dataplane",name:o,mesh:m,networking:h.networking}}),w=S(()=>xe(s.dataPlaneOverview.dataplane)),u=S(()=>{var m;const o=Array.from(((m=s.dataPlaneOverview.dataplaneInsight)==null?void 0:m.subscriptions)??[]);return o.reverse(),o.map(h=>{const d=h.connectTime!==void 0?Ce(h.connectTime):"—",g=h.disconnectTime!==void 0?Ce(h.disconnectTime):"—",A=Object.entries(h.status).filter(([E])=>!["total","lastUpdateTime"].includes(E)).map(([E,I])=>{const $=`${I.responsesAcknowledged??0} / ${I.responsesSent??0}`;return{type:E.toUpperCase(),ratio:$,responsesSent:I.responsesSent??0,responsesAcknowledged:I.responsesAcknowledged??0,responsesRejected:I.responsesRejected??0}});return{subscription:h,formattedConnectDate:d,formattedDisconnectDate:g,statuses:A}})}),f=S(()=>{const{status:o}=Oe(s.dataPlaneOverview.dataplane,s.dataPlaneOverview.dataplaneInsight);return o}),b=S(()=>{const o=He(s.dataPlaneOverview.dataplaneInsight);return o!==null?Object.entries(o).map(([m,h])=>({name:m,version:h})):[]}),k=S(()=>{var I;const o=((I=s.dataPlaneOverview.dataplaneInsight)==null?void 0:I.subscriptions)??[];if(o.length===0)return[];const m=o[o.length-1];if(!m.version)return[];const h=[],d=m.version.envoy,g=m.version.kumaDp;if(!(d.kumaDpCompatible!==void 0?d.kumaDpCompatible:!0)){const $=`Envoy ${d.version} is not supported by Kuma DP ${g.version}.`;h.push($)}if(!(g.kumaCpCompatible!==void 0?g.kumaCpCompatible:!0)){const $=`Kuma DP ${g.version} is not supported by this Kuma control plane.`;h.push($)}return h});return(o,m)=>{const h=Qe("router-link");return c(),L(p(nt),null,{body:D(()=>[y("div",Lt,[y("section",null,[y("h3",Nt,[y("span",null,[l(`
              DPP:

              `),T(h,{to:{name:"data-plane-detail-view",params:{mesh:t.dataPlaneOverview.mesh,dataPlane:t.dataPlaneOverview.name}}},{default:D(()=>[l(P(t.dataPlaneOverview.name),1)]),_:1},8,["to"])]),l(),T(ut,{status:p(f)},null,8,["status"])]),l(),T(ue,{class:"mt-4"},{default:D(()=>[T(F,{term:"Mesh"},{default:D(()=>[l(P(t.dataPlaneOverview.mesh),1)]),_:1}),l(),p(w)!==null?(c(),L(F,{key:0,term:"Tags"},{default:D(()=>[T(dt,{tags:p(w)},null,8,["tags"])]),_:1})):V("",!0),l(),p(b).length>0?(c(),L(F,{key:1,term:"Dependencies"},{default:D(()=>[y("ul",null,[(c(!0),_(N,null,K(p(b),(d,g)=>(c(),_("li",{key:g},P(d.name)+": "+P(d.version),1))),128))]),l(),p(k).length>0?(c(),_(N,{key:0},[y("h5",Ft,[l(`
                  Warnings

                  `),T(p(te),{class:"ml-1",icon:"warning",color:"var(--black-500)","secondary-color":"var(--yellow-300)",size:"20"})]),l(),(c(!0),_(N,null,K(p(k),(d,g)=>(c(),_("p",{key:g},P(d),1))),128))],64)):V("",!0)]),_:1})):V("",!0)]),_:1})]),l(),p(u).length>0?(c(),_("section",Kt,[Mt,l(),y("div",qt,[(c(!0),_(N,null,K(p(u),(d,g)=>(c(),_("div",{key:g},[T(ue,null,{default:D(()=>[T(F,{term:"Connect time","data-testid":`data-plane-connect-time-${g}`},{default:D(()=>[l(P(d.formattedConnectDate),1)]),_:2},1032,["data-testid"]),l(),T(F,{term:"Disconnect time","data-testid":`data-plane-disconnect-time-${g}`},{default:D(()=>[l(P(d.formattedDisconnectDate),1)]),_:2},1032,["data-testid"]),l(),T(F,{term:"CP instance ID"},{default:D(()=>[l(P(d.subscription.controlPlaneInstanceId),1)]),_:2},1024)]),_:2},1024),l(),d.statuses.length>0?(c(),_("details",Rt,[jt,l(),T(ue,null,{default:D(()=>[(c(!0),_(N,null,K(d.statuses,(A,E)=>(c(),L(F,{key:`${g}-${E}`,term:A.type,"data-testid":`data-plane-subscription-status-${g}-${E}`},{default:D(()=>[l(P(A.ratio),1)]),_:2},1032,["term","data-testid"]))),128))]),_:2},1024)])):V("",!0)]))),128))])])):V("",!0),l(),y("section",zt,[T(ct,{id:"code-block-data-plane-summary",content:p(a),"code-max-height":"250px"},null,8,["content"])])])]),_:1})}}});const Ht=fe(Bt,[["__scopeId","data-v-3fa16dcc"]]),Qt=["protocol","service","zone"];function Gt(t){const s=new Map;for(const[a,w]of t){const u=Qt.includes(a),f=u?"tag":a;s.has(f)||s.set(f,[]);const b=s.get(f);let k;f==="tag"?k=(u?`kuma.io/${a}:${w}`:w).replace(/\s+/g,""):k=w,b.push(k.trim())}return s}const Yt=t=>(ce("data-v-dcc59ae1"),t=t(),pe(),t),Zt={key:0},Jt=Yt(()=>y("label",{for:"data-planes-type-filter",class:"mr-2"},`
              Type:
            `,-1)),Wt=["value"],Xt=["for"],ea=["id","checked","onChange"],ta=de({__name:"DataPlaneList",props:{dataPlaneOverviews:{type:Array,required:!0},isLoading:{type:Boolean,required:!1,default:!1},error:{type:[Error,null],required:!1,default:null},nextUrl:{type:String,required:!1,default:null},pageOffset:{type:Number,required:!1,default:0},selectedDppName:{type:String,required:!1,default:null},isGatewayView:{type:Boolean,required:!1,default:!1},gatewayType:{type:String,required:!1,default:"true"},dppFilterFields:{type:Object,required:!0}},emits:["load-data"],setup(t,{emit:s}){const a=t,w=50,u={true:"All",builtin:"Builtin",delegated:"Delegated"},f={title:"No Data",message:"There are no data plane proxies present."},b=Ge(),k=pt(),o=O(Ue),m=O({headers:[],data:[]}),h=O(z.get("filterQuery")??""),d=O(a.gatewayType),g=O({}),A=O(null),E=S(()=>k.getters["config/getMulticlusterStatus"]),I=S(()=>({name:k.getters["config/getEnvironment"]==="universal"?"universal-dataplane":"kubernetes-dataplane"})),$=S(()=>"tag"in a.dppFilterFields?'tag: "kuma.io/protocol: http"':"name"in a.dppFilterFields?"name: cluster":"field: value"),G=S(()=>{let n=gt(E.value,o.value);return a.isGatewayView?n=n.filter(e=>e.key!=="protocol"):n=n.filter(e=>e.key!=="type"),{data:m.value.data,headers:n}}),Y=S(()=>vt.filter(n=>a.isGatewayView?n.tableHeaderKey!=="protocol":n.tableHeaderKey!=="type").filter(n=>E.value?!0:n.tableHeaderKey!=="zone").map(n=>{const e=o.value.includes(n.tableHeaderKey);return{...n,isChecked:e}}));Q(d,function(){M(0)}),Q(g,function(){M(0)}),Q(()=>a.dataPlaneOverviews,function(){J()});function ae(){const n=Pe.get("dpVisibleTableHeaderKeys");Array.isArray(n)&&(o.value=n),J()}ae();function ne(n){M(n)}function M(n){const e={...g.value};"gateway"in e||(e.gateway=d.value),s("load-data",n,e)}function se(n){n.stopPropagation()}function le(n,e){const r=n.target,i=o.value.findIndex(C=>C===e);r.checked&&i===-1?o.value.push(e):!r.checked&&i>-1&&o.value.splice(i,1),Pe.set("dpVisibleTableHeaderKeys",Array.from(new Set(o.value)))}function Z(){at.logger.info(ft.CREATE_DATA_PLANE_PROXY_CLICKED)}async function J(){try{Array.isArray(a.dataPlaneOverviews)&&a.dataPlaneOverviews.length>0?(B(a.selectedDppName??a.dataPlaneOverviews[0].name),m.value.data=await Promise.all(a.dataPlaneOverviews.map(n=>H(n)))):(B(null),m.value.data=[])}catch(n){console.error(n)}}function B(n){n&&a.dataPlaneOverviews.length>0?(A.value=a.dataPlaneOverviews.find(e=>e.name===n)??a.dataPlaneOverviews[0],z.set(a.isGatewayView?"gateway":"dpp",A.value.name)):(A.value=null,z.set(a.isGatewayView?"gateway":"dpp",null))}async function H(n){var ge,ye,he,be,_e;const e=n.mesh,r=n.name,i=((ge=n.dataplane.networking.gateway)==null?void 0:ge.type)||"STANDARD",C={name:i==="STANDARD"?"data-plane-detail-view":"gateway-detail-view",params:{mesh:e,dataPlane:r}},R={name:"mesh-detail-view",params:{mesh:e}},W=["kuma.io/protocol","kuma.io/service","kuma.io/zone"],X=xe(n.dataplane).filter(v=>W.includes(v.label)),oe=(ye=X.find(v=>v.label==="kuma.io/service"))==null?void 0:ye.value,Le=(he=X.find(v=>v.label==="kuma.io/protocol"))==null?void 0:he.value,ie=(be=X.find(v=>v.label==="kuma.io/zone"))==null?void 0:be.value;let me;oe!==void 0&&(me={name:"service-detail-view",params:{mesh:e,service:oe}});let ve;ie!==void 0&&(ve={name:"zones",query:{ns:ie}});const{status:Ne}=Oe(n.dataplane,n.dataplaneInsight),Fe=((_e=n.dataplaneInsight)==null?void 0:_e.subscriptions)??[],Ke={totalUpdates:0,totalRejectedUpdates:0,dpVersion:null,envoyVersion:null,selectedTime:NaN,selectedUpdateTime:NaN,version:null},x=Fe.reduce((v,U)=>{var ke,we;if(U.connectTime){const Te=Date.parse(U.connectTime);(!v.selectedTime||Te>v.selectedTime)&&(v.selectedTime=Te)}const re=Date.parse(U.status.lastUpdateTime);return re&&(!v.selectedUpdateTime||re>v.selectedUpdateTime)&&(v.selectedUpdateTime=re),{totalUpdates:v.totalUpdates+parseInt(U.status.total.responsesSent??"0",10),totalRejectedUpdates:v.totalRejectedUpdates+parseInt(U.status.total.responsesRejected??"0",10),dpVersion:((ke=U.version)==null?void 0:ke.kumaDp.version)||v.dpVersion,envoyVersion:((we=U.version)==null?void 0:we.envoy.version)||v.envoyVersion,selectedTime:v.selectedTime,selectedUpdateTime:v.selectedUpdateTime,version:U.version||v.version}},Ke),j={name:r,nameRoute:C,mesh:e,meshRoute:R,type:i,zone:ie??"—",zoneRoute:ve,service:oe??"—",serviceInsightRoute:me,protocol:Le??"—",status:Ne,totalUpdates:x.totalUpdates,totalRejectedUpdates:x.totalRejectedUpdates,dpVersion:x.dpVersion??"—",envoyVersion:x.envoyVersion??"—",warnings:[],unsupportedEnvoyVersion:!1,unsupportedKumaDPVersion:!1,kumaDpAndKumaCpMismatch:!1,lastUpdated:x.selectedUpdateTime?Se(new Date(x.selectedUpdateTime).toUTCString()):"—",lastConnected:x.selectedTime?Se(new Date(x.selectedTime).toUTCString()):"—",overview:n};if(x.version){const{kind:v}=Ye(x.version);switch(v!==Ze&&j.warnings.push(v),v){case We:j.unsupportedEnvoyVersion=!0;break;case Je:j.unsupportedKumaDPVersion=!0;break}}return E.value&&x.dpVersion&&X.find(U=>U.label===Xe)&&typeof x.version.kumaDp.kumaCpCompatible=="boolean"&&!x.version.kumaDp.kumaCpCompatible&&(j.warnings.push(et),j.kumaDpAndKumaCpMismatch=!0),j}function q({fields:n,query:e}){const r=z.get("filterFields"),i=r!==null?JSON.parse(r):{},C=JSON.stringify(i),R=Object.fromEntries(Gt(n)),W=JSON.stringify(R);z.set("filterQuery",e||null),z.set("filterFields",W),C!==W&&(g.value=R)}return(n,e)=>(c(),L(ot,null,{content:D(()=>{var r;return[T(it,{"selected-entity-name":(r=A.value)==null?void 0:r.name,"page-size":w,"is-loading":a.isLoading,error:t.error,"empty-state":f,"table-data":p(G),"table-data-is-empty":p(G).data.length===0,"show-details":"",next:a.nextUrl!==null,"page-offset":a.pageOffset,onTableAction:e[1]||(e[1]=i=>B(i.name)),onLoadData:ne},{additionalControls:D(()=>[T($t,{id:"data-plane-proxy-filter",class:"data-plane-proxy-filter",placeholder:p($),query:h.value,fields:a.dppFilterFields,onFieldsChange:q},null,8,["placeholder","query","fields"]),l(),a.isGatewayView?(c(),_("div",Zt,[Jt,l(),Ee(y("select",{id:"data-planes-type-filter","onUpdate:modelValue":e[0]||(e[0]=i=>d.value=i),"data-testid":"data-planes-type-filter"},[(c(!0),_(N,null,K(p(u),(i,C)=>(c(),_("option",{key:C,value:C},P(i),9,Wt))),128))],512),[[tt,d.value]])])):V("",!0),l(),T(p(st),{label:"Columns",icon:"cogwheel","button-appearance":"outline"},{items:D(()=>[y("div",{onClick:se},[(c(!0),_(N,null,K(p(Y),(i,C)=>(c(),L(p(lt),{key:C,class:"table-header-selector-item",item:i},{default:D(()=>[y("label",{for:`data-plane-table-header-checkbox-${C}`,class:"k-checkbox table-header-selector-item-checkbox"},[y("input",{id:`data-plane-table-header-checkbox-${C}`,checked:i.isChecked,type:"checkbox",class:"k-input",onChange:R=>le(R,i.tableHeaderKey)},null,40,ea),l(" "+P(i.label),1)],8,Xt)]),_:2},1032,["item"]))),128))])]),_:1}),l(),T(p(Ae),{appearance:"creation",to:p(I),icon:"plus","data-testid":"data-plane-create-data-plane-button",onClick:Z},{default:D(()=>[l(`
            Create data plane proxy
          `)]),_:1},8,["to"]),l(),p(b).query.ns?(c(),L(p(Ae),{key:1,appearance:"primary",icon:"arrowLeft",to:{name:p(b).name},"data-testid":"data-plane-ns-back-button"},{default:D(()=>[l(`
            View All
          `)]),_:1},8,["to"])):V("",!0)]),_:1},8,["selected-entity-name","is-loading","error","table-data","table-data-is-empty","next","page-offset"])]}),sidebar:D(()=>[A.value!==null?(c(),L(Ht,{key:0,"data-plane-overview":A.value},null,8,["data-plane-overview"])):(c(),L(rt,{key:1}))]),_:1}))}});const va=fe(ta,[["__scopeId","data-v-dcc59ae1"]]);export{va as D};
