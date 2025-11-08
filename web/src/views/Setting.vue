<template>
    <div class="setting">
        <div class="setting-container">
            <h2 class="setting-title">系统设置</h2>
            <el-form 
                class="setting-form" 
                label-width="120px"
                :model="formData"
                :rules="rules"
                ref="formRef"
            >
            <el-space fill>
                <el-alert type="info" show-icon :closable="false">
                    <p>基础信息配置:</p>
                </el-alert>
                <el-form-item label="登录密码" label-width="250px" prop="sys_password">
                    <el-input class="setting-input" type="text" v-model="formData.sys_password" />
                </el-form-item>
                
                <el-form-item label="请填写host地址" label-width="250px" prop="host">
                    <el-input class="setting-input" v-model="formData.host" />
                </el-form-item>
          
                <el-form-item label="请填写默认路径" label-width="250px" prop="default_path">
                    <el-input class="setting-input" v-model="formData.default_path" />
                </el-form-item>

                <el-form-item label="保存图片url模式" label-width="250px" prop="url_mode">
                    <el-select class="setting-input" v-model="formData.url_mode" placeholder="请选择url模式">
                        <el-option label="保持上传时图片名，不做任何修改" value="keep" />
                        <el-option label="年月/时间戳（例子：202511/1731031200.png）" value="ym_timestamp" />
                    </el-select>
                </el-form-item>
                

                <el-form-item label="复制图片时格式" label-width="250px" prop="copy_img_format">
                    <el-select class="setting-input" v-model="formData.copy_img_format" placeholder="请选择复制图片格式">
                        <el-option label="默认url" value="default" />
                        <el-option label="markdown格式" value="markdown" />
                    </el-select>
                </el-form-item>
            </el-space>

            <el-space fill>
                <el-alert type="info" show-icon :closable="false">
                    <p>系统信息配置:</p>
                </el-alert>

                <el-form-item label="请填写region" label-width="250px" prop="region">
                    <el-input class="setting-input" v-model="formData.region" />
                </el-form-item>

                <el-form-item label="请填写bucket" label-width="250px" prop="bucket">
                    <el-input class="setting-input" v-model="formData.bucket" />
                </el-form-item>

                <el-form-item label="请填写access_key_id" label-width="250px" prop="access_key_id">
                    <el-input class="setting-input" v-model="formData.access_key_id" />
                </el-form-item>

                <el-form-item label="请填写access_key_secret" label-width="250px" prop="access_key_secret">
                    <el-input class="setting-input" v-model="formData.access_key_secret" />
                </el-form-item>
               
            </el-space>

                <el-form-item>
                    <el-button type="primary" class="setting-button" @click="submitForm">提交</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script setup>
import { getAliyunSecret, saveAliyunSecret } from '@/api/system'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const formData = ref({
    region: '',
    bucket: '',
    host: '',
    access_key_id: '',
    access_key_secret: '',
    default_path: '',
    sys_password: '',
    url_mode: 'keep',
    copy_img_format: 'default',
})

const rules = ref({
    region: [
        { required: true, message: '请填写region', trigger: 'blur' }
    ],
    host: [
        { required: true, message: '请填写host地址', trigger: 'blur' }
    ],
    url_mode: [
        { required: true, message: '请选择url模式', trigger: 'blur' }
    ],
    copy_img_format: [
        { required: true, message: '请选择复制图片格式', trigger: 'blur' }
    ],
    bucket: [
        { required: true, message: '请填写bucket', trigger: 'blur' }
    ],
    access_key_id: [
        { required: true, message: '请填写access_key_id', trigger: 'blur' }
    ],
    access_key_secret: [
        { required: true, message: '请填写access_key_secret', trigger: 'blur' }
    ],
   
    sys_password: [
        { required: true, message: '请填写登录密码', trigger: 'blur' }
    ]
})

const formRef = ref(null)

//提交表单
const submitForm = async () => {
    formRef.value.validate( async (valid) => {
        if (valid) {
            const res = await saveAliyunSecret(formData.value)
            if (res.code === 0) {
                ElMessage({
                    type: 'success',
                    message: '修改成功！',
                    showClose: true,
                    onClose: () => {
                       //刷新页面
                       window.location.reload()
                    }
                })
               
            } else {
                ElMessage({
                    type: 'error',
                    message: res.message,
                    showClose: true,
                   
                })
            }
        }
    })
}

//初始化表单
const initFormData = async () => {
    const  res = await getAliyunSecret()
    if (res.code  == 0) {
        if (res.data.url_mode == '') {
            res.data.url_mode = 'keep'
        }
        if (res.data.copy_img_format == '') {
            res.data.copy_img_format = 'default'
        }
        formData.value = res.data
    }
}
initFormData()


</script>

<style scoped>
.setting {
    background-color: #fff;
    min-height: calc(100vh - 60px);
}

.setting-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 40px 20px;
}

.setting-title {
    font-size: 28px;
    color: #2d6cdf;
    margin-bottom: 40px;
    font-weight: 500;
    position: relative;
    padding-left: 15px;
}

.setting-title::before {
    content: '';
    position: absolute;
    left: 0;
    top: 50%;
    transform: translateY(-50%);
    width: 4px;
    height: 20px;
    background: linear-gradient(120deg, #2d6cdf 0%, #4c89f0 100%);
    border-radius: 2px;
}

.setting-form {
    max-width: 800px;
}

.setting-input {
    width: 400px;  /* 统一输入框宽度 */
}

.setting-input :deep(.el-input__wrapper) {
    box-shadow: 0 0 0 1px #dcdfe6 inset;
}

.setting-input :deep(.el-input__wrapper:hover) {
    box-shadow: 0 0 0 1px #2d6cdf inset;
}

.setting-button {
    width: 200px;
    background: linear-gradient(120deg, #2d6cdf 0%, #4c89f0 50%, #6ea5ff 100%);
    border: none;
    height: 40px;
    margin-top: 20px;
}

.setting-button:hover {
    background: linear-gradient(120deg, #245bb9 0%, #4075d1 50%, #5c8ee6 100%);
}

:deep(.el-form-item__label) {
    color: #606266;
    font-weight: 500;
    font-size: 16px;
}
</style>