<template>
  <el-form ref="formRef" :model="form" label-width="120px">
    <el-form-item label="账号">
      <el-input v-model="form.uri"></el-input>
    </el-form-item>
    <el-form-item label="保存目录">
      <el-input v-model="form.dir"></el-input>
    </el-form-item>

    <el-form-item>
      <el-button  @click="onSubmit">保存</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { reactive } from 'vue'
import { ElMessage,ElLoading } from 'element-plus'
import  axios  from 'axios'

export default {
  components: {
  },
  setup() {
    const onSubmit = () => {
      console.log('---------')
      console.log(form)


      axios.post('http://localhost:9999/home', form)
          .then(function (response) {
            if (response.data.msg === "成功") {
              // 成功提示
              ElMessage({
                message: response.data.data,
                type: 'success',
              })
            } else {
              // alert(response.data.msg)
            }
          })
          .catch(function (error) {
            // alert("异常,请联系反馈");
          });
    };

    const form = reactive({
      uri: '',
      dir: '',
    })

    return {
      onSubmit,
      form,
    };
  },
};
</script>
