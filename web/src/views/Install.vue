<template>

<div class="setting">
        <div class="setting-container">
            <h2 class="setting-title">系统初始化</h2>
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
                    <el-button type="primary" class="setting-button" @click="submitForm">初始化</el-button>
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
            const res = await createAliyunSecret(formData.value)
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


</script>

<style scoped>
.setting {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 40px 20px;
    background: linear-gradient(135deg, #f5f9ff 0%, #f0f5ff 50%, #fdf6ff 100%);
}

.setting-container {
    width: 100%;
    max-width: 960px;
    background: #ffffffcc;
    backdrop-filter: blur(12px);
    border-radius: 20px;
    box-shadow: 0 25px 60px rgba(64, 99, 255, 0.15);
    padding: 48px 56px;
    border: 1px solid rgba(99, 102, 241, 0.12);
}

.setting-title {
    font-size: 34px;
    font-weight: 700;
    color: #334155;
    text-align: center;
    margin-bottom: 32px;
    letter-spacing: 2px;
}

.setting-form {
    display: flex;
    flex-direction: column;
    gap: 32px;
}

:deep(.el-space) {
    width: 100%;
    background: rgba(255, 255, 255, 0.75);
    border-radius: 16px;
    padding: 24px;
    box-shadow: inset 0 0 0 1px rgba(99, 102, 241, 0.08);
    gap: 18px !important;
}

:deep(.el-alert) {
    border-radius: 12px;
    padding: 12px 16px;
    color: #3b82f6;
    background: rgba(59, 130, 246, 0.08);
}

:deep(.el-alert__title) {
    font-weight: 600;
}

:deep(.el-form-item__label) {
    font-size: 14px;
    color: #475569;
    font-weight: 500;
}

:deep(.el-input__wrapper),
:deep(.el-select .el-input__wrapper) {
    
    background: rgba(255, 255, 255, 0.9);
    border: 1px solid rgba(148, 163, 184, 0.3);
    box-shadow: 0 4px 14px rgba(15, 23, 42, 0.06);
    transition: border 0.2s ease, box-shadow 0.2s ease;
}

:deep(.el-input__wrapper:hover),
:deep(.el-input__wrapper.is-focus) {
    border-color: #6366f1;
    box-shadow: 0 8px 20px rgba(99, 102, 241, 0.18);
}

:deep(.el-select .el-select__caret) {
    color: #6366f1;
}

.setting-button {
    width: 100%;
    height: 52px;
    border-radius: 14px;
    font-size: 18px;
    font-weight: 600;
    letter-spacing: 1px;
    background: linear-gradient(135deg, #6366f1, #8b5cf6);
    border: none;
    box-shadow: 0 18px 30px rgba(99, 102, 241, 0.25);
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.setting-button:hover {
    transform: translateY(-2px);
    box-shadow: 0 20px 35px rgba(99, 102, 241, 0.35);
}

.setting-button:active {
    transform: translateY(0);
}

@media screen and (max-width: 768px) {
    .setting-container {
        padding: 32px 24px;
    }

    :deep(.el-space) {
        padding: 20px;
    }

    .setting-title {
        font-size: 28px;
        margin-bottom: 24px;
    }

    :deep(.el-form-item__label) {
        width: 160px !important;
    }
}
</style>