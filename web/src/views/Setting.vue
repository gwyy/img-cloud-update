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
                <el-form-item label="请填写appid" prop="appid">
                    <el-input class="setting-input" v-model="formData.appid" />
                </el-form-item>

                <el-form-item label="请填写ac" prop="ac">
                    <el-input class="setting-input" v-model="formData.ac" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" class="setting-button" @click="submitForm">提交</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive } from 'vue'


const formRef = ref(null)
const formData = reactive({
    appid: '',
    ac: ''
})

const rules = {
    appid: [
        { required: true, message: '请输入appid', trigger: 'blur' }
    ],
    ac: [
        { required: true, message: '请输入ac', trigger: 'blur' }
    ]
}

const submitForm = async () => {
    if (!formRef.value) return
    
    try {
        await formRef.value.validate()
        // TODO: 这里添加提交逻辑
        console.log('表单提交成功', formData)
    } catch (error) {
        console.error('表单验证失败')
    }
}
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