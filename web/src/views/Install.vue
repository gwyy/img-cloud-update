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
                <el-form-item label="请设置登录密码" label-width="250px" prop="password">
                    <el-input class="setting-input" type="text" v-model="formData.password" />
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
import { install,isInstall } from '@/api/system'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useTokenStore } from '@/store'
import { useRouter } from 'vue-router'
const router = useRouter()


//判断是否安装
const checkInstall = async () => {
    const res = await isInstall()
    if (res.code === 0) {
        if (res.data.isInstall === "true") {
            router.push('/login')
        } else {
            return
        }
    } else {
        ElMessage.error(res.message)
    }
}
checkInstall()




const formData = ref({
    password: '',
})

const rules = ref({
   
    password: [
        { required: true, message: '请填写登录密码', trigger: 'blur' }
    ]
})

const formRef = ref(null)

//提交表单
const submitForm = async () => {
    formRef.value.validate( async (valid) => {
        if (valid) {
            const res = await install(formData.value)
            if (res.code === 0) {

                //保存密码到本地
                useTokenStore().login(formData.value.password)
                //刷新页面
                ElMessage({
                    type: 'success',
                    message: '初始化成功！',
                    showClose: true,
                    onClose: () => {
                       //跳转到首页
                       router.push('/index')
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