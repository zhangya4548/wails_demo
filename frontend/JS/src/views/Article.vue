<template>
  <el-form ref="formRef" :model="form" label-width="120px">
    <el-form-item label="声音选择">

      <el-select v-model="form.resource" class="m-2" placeholder="Select">
        <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        >
        </el-option>
      </el-select>
    </el-form-item>

    <el-form-item label="图文内容" prop="goodsDetail">
      <editor v-model="content" :init="init"/>

<!--      <div style="border: 1px solid #ccc">-->
<!--        <Toolbar-->
<!--            :editorId="editorId"-->
<!--            :defaultConfig="toolbarConfig"-->
<!--            :mode="mode"-->
<!--            style="border-bottom: 1px solid #ccc"-->
<!--        />-->
<!--        <Editor-->
<!--            :editorId="editorId"-->
<!--            :defaultConfig="editorConfig"-->
<!--            :defaultContent="getDefaultContent"-->
<!--            :defaultHtml="defaultHtml"-->
<!--            :mode="mode"-->
<!--            style="height: 500px; overflow-y: hidden;"-->
<!--        />-->
<!--        &lt;!&ndash; 注意: defaultContent (JSON 格式) 和 defaultHtml (HTML 格式) ，二选一 &ndash;&gt;-->
<!--      </div>-->
    </el-form-item>



    <el-form-item>
      <el-button  @click="onSubmit">确认生成</el-button>
    </el-form-item>
  </el-form>
</template>
<!--<style src="@wangeditor/editor/dist/css/style.css"></style>-->
<script>

import {onMounted, computed, onBeforeUnmount, reactive, ref,watch} from 'vue'
// import { Editor, Toolbar, getEditor, removeEditor } from '@wangeditor/editor-for-vue'
// import cloneDeep from 'lodash.clonedeep'


import tinymce from 'tinymce/tinymce'
import Editor from '@tinymce/tinymce-vue'
import 'tinymce/themes/silver/theme'; // 引用主题文件
import 'tinymce/icons/default'; // 引用图标文件
import "tinymce/plugins/image";
import "tinymce/plugins/link";
import "tinymce/plugins/code";
import "tinymce/plugins/table";
import "tinymce/plugins/lists";

// import zh_CN from "@/assets/langs/zh_CN";

export default {
  components: {
    'editor': Editor,
    // Editor, Toolbar
  },

  setup() {

    // const editorId = `w-e-${Math.random().toString().slice(-5)}` //【注意】编辑器 id ，要全局唯一
    //
    // // defaultContent (JSON 格式) 和 defaultHtml (HTML 格式) ，二选一
    // const defaultHtml = ''
    // const defaultContent = [
    //   { type: 'paragraph', children: [{ text: '' }] }
    // ]
    // const getDefaultContent = computed(() => cloneDeep(defaultContent)) // 注意，要深拷贝 defaultContent ，否则报错
    //
    // const toolbarConfig = {}
    // const editorConfig = { placeholder: '请输入格式:(一张图片配上一段内容往下排)' }
    //
    // // 组件销毁时，也及时销毁编辑器
    // onBeforeUnmount(() => {
    //   const editor = getEditor(editorId)
    //   if (editor == null) return
    //
    //   editor.destroy()
    //   removeEditor(editorId)
    // })


    const content = ref('');
    const init = {
      language_url: "../assets/langs/zh_CN.js", // 中文语言包路径
      // language_url: zh_CN, // 中文语言包路径
      language: "zh_CN",
      skin_url: "../assets/skins/ui/oxide/skin.css", // 编辑器皮肤样式
      menubar: false, // 隐藏菜单栏
      height: 320,
      toolbar_mode: "none",
      plugins: 'image link code table lists', // 插件需要import进来
      toolbar: 'fontselect fontsizeselect | link lineheight forecolor backcolor bold italic underline strikethrough | alignleft aligncenter alignright alignjustify | image quicklink h2 h3 blockquote table numlist bullist preview fullscreen',
      content_style: "p {margin: 5px 0; font-size: 14px}",
      fontsize_formats: "12px 14px 16px 18px 24px 36px 48px 56px 72px",
      font_formats: "微软雅黑=Microsoft YaHei,Helvetica Neue,PingFang SC,sans-serif;苹果苹方=PingFang SC,Microsoft YaHei,sans-serif;宋体=simsun,serif;仿宋体=FangSong,serif;黑体=SimHei,sans-serif;Arial=arial,helvetica,sans-serif;Arial Black=arial black,avant garde;Book Antiqua=book antiqua,palatino;",
      branding: false,
      elementpath: false,
      resize: false, // 禁止改变大小
      statusbar: false, // 隐藏底部状态栏
    };

    // 监听编辑器内容变化，自定义处理内容
    watch(() => content.value, (value) => {
      // TODO
    })




    const onSubmit = () => {
      console.log('---------')
      console.log(form)
    };

    const form = reactive({
      wx_url: '',
      resource:'Option2',
    })

    const options = [
      {
        value: 'Option1',
        label: 'Option1',
      },
      {
        value: 'Option2',
        label: 'Option2',
      },
    ]

    onMounted(async()=>{
      //dom 挂载后
      console.log("*******onMounted******")
      tinymce.init(); // 初始化
    })

    return {
      content,
      init,
      options,
      onSubmit,
      mode: 'default',
      // editorId,
      // defaultHtml,
      // getDefaultContent,
      // toolbarConfig,
      // editorConfig,
      form,
    };
  },
};
</script>
