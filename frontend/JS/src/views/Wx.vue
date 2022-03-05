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

    <el-form-item label="微信文章url">
      <el-input v-model="form.wx_url"></el-input>
    </el-form-item>

    <el-form-item>
      <el-button  @click="onSubmit">确认生成</el-button>
    </el-form-item>
  </el-form>
</template>

<script>

import { reactive } from 'vue'
import {ElMessage} from "element-plus";

export default {
  components: {
  },
  setup() {

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

    const form = reactive({
      wx_url: '',
      resource:'Option2',
    })


    const onSubmit = () => {
      console.log('---------')
      console.log(form)

      // 成功提示
      ElMessage({
        message: form,
        type: 'success',
      })
    };

    return {
      options,
      onSubmit,
      form,
    };
  },
};
</script>
