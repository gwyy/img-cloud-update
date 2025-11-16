<template>
    <div class="login">
        <el-form :model="form" :rules="rules" ref="formRef">
        <div class="login-container">
            <div class="login-title">请输入密码</div>
            <div class="login-form">
                <el-form-item prop="password">
                    <el-input 
                        v-model="form.password" 
                        type="password" 
                        placeholder="请输入密码"
                        size="large"
                        @keyup.enter="login"
                    />
                </el-form-item>
                <el-button type="primary" size="large" @click="login">登录</el-button>
            </div>
        </div>
    </el-form>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { loginApi } from '@/api/system'
import { ElMessage } from 'element-plus'
import { isInstall } from '@/api/system'
import { useTokenStore } from '@/store'
const router = useRouter()

const tokenStore = useTokenStore()


//检查是否安装
const checkInstall = async () => {
    const res = await isInstall()
    if (res.code === 0) {
        if (res.data.isInstall === "false") {
            router.push('/install')
        }
    } else {
        ElMessage.error(res.message)
    }
}
 checkInstall()



const form = ref({
    password: ''
})
const rules = ref({
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' }
    ]
})
const formRef = ref(null)


const login = () => {
    formRef.value.validate( async (valid) => {
        if (valid) {
            const res = await loginApi(form.value)
            if (res.code === 0) {
                // 保存token
                tokenStore.login(form.value.password)
                router.push('/index')
            } else {
                ElMessage.error(res.message)
            }
        }
    })
}
</script>

<style scoped>
.login {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    background: #0a0e27;
    position: relative;
    overflow: hidden;
}

/* 动态网格背景 */
.login::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-image: 
        linear-gradient(rgba(0, 255, 255, 0.1) 1px, transparent 1px),
        linear-gradient(90deg, rgba(0, 255, 255, 0.1) 1px, transparent 1px);
    background-size: 50px 50px;
    animation: gridMove 20s linear infinite;
    transform: perspective(500px) rotateX(60deg);
}

/* 发光粒子效果 */
.login::after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: radial-gradient(circle at 20% 50%, rgba(120, 0, 255, 0.1) 0%, transparent 50%),
                radial-gradient(circle at 80% 80%, rgba(0, 255, 255, 0.1) 0%, transparent 50%);
    animation: pulse 4s ease-in-out infinite;
}

@keyframes gridMove {
    0% {
        background-position: 0 0;
    }
    100% {
        background-position: 50px 50px;
    }
}

@keyframes pulse {
    0%, 100% {
        opacity: 0.5;
    }
    50% {
        opacity: 1;
    }
}

.login-container {
    background: rgba(15, 20, 40, 0.8);
    backdrop-filter: blur(10px);
    padding: 60px 80px;
    border-radius: 20px;
    border: 1px solid rgba(0, 255, 255, 0.3);
    box-shadow: 
        0 0 60px rgba(0, 255, 255, 0.3),
        inset 0 0 60px rgba(0, 255, 255, 0.05);
    min-width: 500px;
    position: relative;
    z-index: 1;
}

.login-title {
    font-size: 32px;
    font-weight: 700;
    background: linear-gradient(90deg, #00ffff, #7000ff);
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
    text-align: center;
    margin-bottom: 40px;
    text-transform: uppercase;
    letter-spacing: 2px;
    text-shadow: 0 0 30px rgba(0, 255, 255, 0.5);
}

.login-form {
    display: flex;
    gap: 16px;
    align-items: center;
}

.login-form :deep(.el-form-item) {
    flex: 1;
    margin-bottom: 0;
}

.login-form :deep(.el-form-item__error) {
    color: #ff0080;
    text-shadow: 0 0 10px rgba(255, 0, 128, 0.8);
}

.login-form :deep(.el-input) {
    width: 100%;
}

.login-form :deep(.el-input__wrapper) {
    background: rgba(0, 0, 0, 0.3);
    border: 1px solid rgba(0, 255, 255, 0.3);
    box-shadow: 0 0 20px rgba(0, 255, 255, 0.2);
    transition: all 0.3s;
}

.login-form :deep(.el-input__wrapper:hover),
.login-form :deep(.el-input__wrapper.is-focus) {
    border-color: #00ffff;
    box-shadow: 0 0 30px rgba(0, 255, 255, 0.5);
}

.login-form :deep(.el-input__inner) {
    color: #00ffff;
    font-size: 16px;
}

.login-form :deep(.el-input__inner::placeholder) {
    color: rgba(0, 255, 255, 0.4);
}

.login-form :deep(.el-button) {
    padding: 12px 40px;
    background: linear-gradient(135deg, #7000ff, #00ffff);
    border: none;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 1px;
    box-shadow: 0 0 30px rgba(0, 255, 255, 0.4);
    transition: all 0.3s;
}

.login-form :deep(.el-button:hover) {
    box-shadow: 0 0 50px rgba(0, 255, 255, 0.8);
    transform: translateY(-2px);
}
</style>