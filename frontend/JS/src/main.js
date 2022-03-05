import { createApp } from "vue";
import App from "./App.vue";
import router from "@/router";
import i18n from "@/i18n";

import ElementPlus from 'element-plus'
// import 'element-plus/dist/index.css'
// import 'element-plus/lib/theme-chalk/index.css'
import 'element-plus/theme-chalk/index.css'

// Register global common components
// 注册全局通用组件
import publicComponents from "@/components/public";

createApp(App).use(router).use(i18n).use(publicComponents).use(ElementPlus).mount("#app");
