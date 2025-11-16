<template>
  <div class="index">
    <el-form :model="formData" ref="formRef" :rules="rules">

      <!-- 选择地址区域 -->
      <div class="select-container">
        <div class="title">请选择云厂商</div>
        <div class="yun-complate">阿里云</div>
      </div>
      
      <!-- 选择上传到哪个文件夹 -->
      <div class="select-container">
        <div class="title">选择上传路径</div>
        <div class="yun-path">
          <span>{{ host }}</span>
          <el-input class="update-path" v-model="formData.path"  placeholder="请输入路径" />
        </div>
      </div>


      <!-- 上传区域 -->
      <div class="upload-container">
        <el-upload
          class="upload-area"
          drag
          ref="uploadRef"
          action="#"
          accept="image/*"
          :auto-upload="false"
          :show-file-list="false"
          :http-request="handleUpload"
          :on-change="handleFileChange"
        >
          <div class="upload-content">
            <i class="el-icon-upload"></i>
            <div class="upload-text">请上传文件</div>
          </div>
        </el-upload>
      </div>

      <!-- 缩略图显示区域 -->
      <div class="thumbnail-container">
        <div class="thumbnail-list">
        
          <div class="thumbnail-item" v-for="(image, index) in uploadedImages" :key="index" >
            <div class="thumbnail-wrapper">
              <div class="thumbnail-visual">
                <el-image :src="image.url" 
                :preview-src-list="uploadedImages.map(item => item.url)"
                fit="cover"
                :initial-index="index"
                scale="1"
                :zoom-rate="1.2"
                :max-scale="7"
                :min-scale="0.2"
                hide-on-click-modal=true />
                <div class="thumbnail-actions">
                  <el-button type="danger" size="small" @click="onDeleteImage(image.name)">
                    <el-icon  class="el-icon-delete" >
                      <Delete />
                    </el-icon>
                  </el-button>
                </div>
              </div>
              <button class="thumbnail-copy" type="button" @click="copyFilePath(image.url)">复制路径</button>
            </div>
          </div>

        </div>
      </div>
      
      <!-- 上传完成后文件路径显示区域 -->
      <div class="file-path-container">
        <el-input
          v-model="serverFileName"
          class="path-input"
          :placeholder="'请上传文件，服务器路径将在此显示'"
        >
          <template #append>
            <el-button type="primary" @click="copyFilePath('')" :disabled="!serverFileName">
              <i class="el-icon-document-copy"></i>
              复制路径
            </el-button>
          </template>
        </el-input>
      </div>
    </el-form>
  </div>
</template>

<script setup>
// 从 Vue 中导入 ref 响应式引用
import { ref,reactive } from 'vue'
// 导入 Element Plus 的消息提示组件
import { ElMessage } from 'element-plus'
import { uploadImage,deleteImage,getAliyunImgData } from '@/api/system'

import { useRouter } from 'vue-router'
const router = useRouter()


// 表单引用
const formRef = ref(null)
const uploadRef = ref(null)
// 控制是否第一次上传
const isFirstUpload = ref(true)
//上传回调地址
const serverFileName = ref('')

// 二次构造的表单数据
const formData = reactive({
  path: '',
  file: null
})

//复制图片格式 default or  markdown
var copyImgFormat = ""
var host = ref('')
//初始化阿里云信息
const initAliyunImgData = async () => {
  const response = await getAliyunImgData()
  if (response.code === 0) {
    formData.path = response.data.path
    copyImgFormat = response.data.copyImgFormat
    host.value = response.data.host

    if (host.value == '') {
      host.value = '-- 未设置host地址 --'
    }
  }
}
initAliyunImgData()



//上传的图片列表
const uploadedImages = ref([]);
// const uploadedImages = ref([  {
//     url: 'https://img1.liangtian.me/auto/202511/1762610929.jpg',
//     name: 'auto/202511/1762610929.jpg'
//   },
// {
//   url: 'https://img1.liangtian.me/auto/202511/1762611086.jpg',
//   name: 'auto/202511/1762611086.jpg'
// },])







// 表单校验规则
const rules = reactive({
  path: [{ required: true, message: '请输入路径', trigger: 'blur' }],
  file: [{ required: true, message: '请上传文件', trigger: 'change' }]
})





// 文件选择后的处理
const handleFileChange = async (file) => {
  formData.file = file.raw
  if (isFirstUpload.value) {
    isFirstUpload.value = false
    try {
      // 触发表单验证
      await formRef.value.validate()
      // 手动触发上传
      await uploadRef.value.submit()
    } catch (error) {
      ElMessage.error('表单验证失败')
      isFirstUpload.value = true // 验证失败时允许下次重试
    }
  }
}

// 上传文件
const handleUpload = async () => {
  if(isFirstUpload.value) {
    return
  }

  // 创建 FormData 对象
  const form = new FormData()
  form.append('path', formData.path)
  form.append('file', formData.file)


  //发送请求 
  try {
    // 使用 axios 发送请求
    const response = await uploadImage(form)
    if (response.code === 0) {
      ElMessage.success('上传成功')
      //存储上传地址
      serverFileName.value = setFinalImageUrl(response.data.url)
      // 添加新上传的图片到缩略图列表
      addThumbnail(response.data.url,response.data.name)
    
      console.log('服务器响应:', response.data)
    } else {
      ElMessage.error(response.message)
    }
  } catch (error) {
     ElMessage.error('错误：'+ error.message)
  }

  // 清空文件和表单
  // 清空表单
  formRef.value.resetFields()
  uploadRef.value.clearFiles()
  formData.file = null
  isFirstUpload.value = true // 重置为允许下次上传
}



//设置最终图片地址
const setFinalImageUrl = (url) => {
  if(url == ''){
    return ''
  }
  if(copyImgFormat == 'default'){
    return url
  }else if(copyImgFormat == 'markdown'){
    return "![]("+url+")"
  }
}



//添加缩略图
const addThumbnail = (url,name) => {
  //判断长度 超过10个就去掉最后一个
  if(uploadedImages.value.length >= 10){
    uploadedImages.value.shift()
  }
  uploadedImages.value.push({
    url: url,
    name: name
  })
}


// 删除图片
const onDeleteImage = async (name) => {
  if (name == "") {
    return
  }
  //  //删除缩略图
  //  uploadedImages.value = uploadedImages.value.filter(item => item.name !== name)
  //   //清空file
  //   serverFileName.value = ''
  //   return 
  //发送请求 删除图片
  const response = await deleteImage(name)
  if (response.code === 0) {
    ElMessage.success('删除成功')
    //删除缩略图
    uploadedImages.value = uploadedImages.value.filter(item => item.name !== name)
    //清空file
    serverFileName.value = ''
  } else {
    ElMessage.error('删除失败')
  }
}



// 定义复制文件路径的异步函数
const copyFilePath = async (url) => {
  var finalValue = ""
  if (url != '') {
    finalValue = setFinalImageUrl(url)
  } else {
    finalValue = serverFileName.value
  }
  if (finalValue == '') {
    ElMessage.warning('暂无可复制的路径')
    return
  }
  try {
    await navigator.clipboard.writeText(finalValue)
    ElMessage.success('路径已复制到剪贴板')
  } catch (err) {
    ElMessage.error('复制失败，请手动复制')
  }
  serverFileName.value = finalValue
}
</script>

<style scoped>
.index {
  width: 100%;
  overflow: hidden;
  display: flex;
  flex-direction: column;  /* 改为纵向排列 */
  align-items: center;
  padding: 0;
  margin: 0;
}

.select-container {
  width: 80vw;           /* 与上传区域保持同宽 */
  max-width: 1200px;
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 10px;
  margin-top: 10px;
}

.upload-container {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0;
  margin: 20px;
}

/* 上传区域的外层容器，用于居中定位上传组件 */
.upload-area {
  width: 80vw;           
  min-width: 300px;     /* 修改最小宽度，适应移动端 */
  max-width: 1200px;    /* 新增最大宽度限制 */
  height: 300px;        
}

/* Element Plus上传组件的拖拽区域样式自定义 */
:deep(.el-upload-dragger) {
  width: 100%;              
  height: 300px !important; /* 减小高度从300px到200px */
  display: flex;            /* 使用弹性布局，便于内容居中 */
  align-items: center;      /* 内容垂直居中 */
  justify-content: center;  /* 内容水平居中 */
  border: 4px dashed #B3D8FF;  /* 虚线边框，提供视觉引导 */
  border-radius: 12px;      /* 圆角边框，提升视觉体验 */
  background: rgba(179, 216, 255, 0.03);  /* 轻微的背景色，区分上传区域 */
  transition: all 0.5s;     /* 添加过渡效果，使状态变化更平滑 */
  position: relative;       /* 为动画效果提供定位上下文 */
  overflow: hidden;         /* 隐藏超出部分，主要用于动画效果 */
}

/* 创建流动动画效果的光效 */
:deep(.el-upload-dragger)::before {
  content: '';             /* 创建伪元素 */
  position: absolute;      /* 绝对定位，相对于父元素 */
  top: 0;
  left: -100%;            /* 初始位置在可视区域左侧 */
  width: 100%;
  height: 100%;
  /* 创建渐变效果，形成流动的光带 */
  background: linear-gradient(
    90deg,
    transparent,
    rgba(179, 216, 255, 0.2),
    transparent
  );
  animation: flow 3s infinite;  /* 应用流动动画，无限循环 */
}

/* 定义流动动画关键帧 */
@keyframes flow {
  0% { left: -100%; }     /* 动画开始：在左侧不可见区域 */
  50% { left: 100%; }     /* 动画中点：移动到右侧 */
  100% { left: 100%; }    /* 动画结束：保持在右侧 */
}

/* 鼠标悬停时的交互效果 */
:deep(.el-upload-dragger:hover) {
  border-color: #B3D8FF;   /* 加深边框颜色 */
  background: rgba(179, 216, 255, 0.08);  /* 增加背景不透明度 */
  transform: scale(1.02);   /* 轻微放大效果 */
  box-shadow: 0 0 30px rgba(179, 216, 255, 0.3);  /* 添加发光效果 */
}

/* 上传内容区域（图标和文字的容器） */
.upload-content {
  text-align: center;      /* 文字居中对齐 */
}

/* 上传图标样式 */
.upload-content i {
  font-size: 80px;        /* 设置图标大小 */
  color: #409EFF;         /* 设置图标颜色为主题蓝 */
  margin-bottom: 20px;    /* 与下方文字保持间距 */
  transition: all 0.3s;    /* 添加过渡效果 */
}

/* 鼠标悬停时图标的动画效果 */
:deep(.el-upload-dragger:hover) .upload-content i {
  transform: scale(1.1) translateY(-5px);  /* 放大并向上浮动 */
  color: #79bbff;         /* 图标颜色变浅 */
}

/* 上传区域的提示文字样式 */
.upload-text {
  font-size: 28px;        /* 设置文字大小 */
  color: #409EFF;         /* 文字颜色与图标保持一致 */
  font-weight: 500;       /* 适中的文字粗细 */
  transition: all 0.3s;    /* 添加过渡效果 */
}

/* 鼠标悬停时文字的动画效果 */
:deep(.el-upload-dragger:hover) .upload-text {
  transform: scale(1.05);  /* 轻微放大效果 */
  color: #79bbff;         /* 文字颜色变浅 */
}

/* 文件路径容器样式 */
.file-path-container {
  min-width: 300px;       /* 修改最小宽度 */
  max-width: 1200px;      /* 新增最大宽度限制 */
  width:80vw;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 10px;
  padding: 20px 10px;
  margin-top: 20px;
  width: 90vw;
}

/* 路径输入框样式 */
.path-input {
  --el-input-height: 60px;
  font-size: 20px;
}

:deep(.path-input .el-input__wrapper) {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  padding: 0 20px;
}

/* 路径输入框的按钮组样式 */
:deep(.path-input .el-input-group__append) {
  padding: 0;
  border: none;
  background: transparent;
}

/* 复制按钮样式优化 */
:deep(.path-input .el-button) {
  height: 60px;
  padding: 0 30px;
  font-size: 18px;
  font-weight: 500;
  border: none;
  border-radius: 0 8px 8px 0;
  background: linear-gradient(135deg, #409EFF, #79bbff);
  color: white;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

:deep(.path-input .el-button:hover) {
  transform: translateY(-1px);
  background: linear-gradient(135deg, #66b1ff, #8cc5ff);
  box-shadow: 0 4px 12px rgba(64, 158, 255, 0.4);
  color:#fff;
}

:deep(.path-input .el-button:active) {
  transform: translateY(1px);
  box-shadow: 0 2px 6px rgba(64, 158, 255, 0.4);
}

:deep(.path-input .el-button:disabled) {
  background: #c0c4cc;
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
}

:deep(.path-input .el-button i) {
  margin-right: 4px;
  font-size: 20px;
}

/* 新增选择地址区域样式 */
.title {
  font-size: 28px;
    color: #2d6cdf;
    font-weight: 500;
    position: relative;
    padding-left: 15px;
}
.title::before {
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

.path-select {
  width: 200px;
}

:deep(.path-select .el-input__wrapper) {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
}

.custom-option {
  display: flex;
  align-items: center;
  gap: 8px;
}

.custom-option i {
  color: #409EFF;
  font-size: 18px;
}

/* 修改移动端适配 */
@media screen and (max-width: 768px) {
  .select-container {
    width: 90vw;
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
    padding: 10px;
  }

  .upload-area {
    height: 200px;      /* 移动端减小高度 */
  }

  :deep(.el-upload-dragger) {
    height: 200px !important;
  }

  .upload-content i {
    font-size: 50px;    /* 减小图标大小 */
    margin-bottom: 10px;
  }

  .upload-text {
    font-size: 20px;    /* 减小文字大小 */
  }

  .file-path-container {
    bottom: 100px;      /* 调整位置 */
  }

  .path-input {
    --el-input-height: 45px;  /* 减小输入框高度 */
    font-size: 16px;          /* 调整字体大小 */
  }

  :deep(.path-input .el-button) {
    height: 45px;             /* 匹配输入框高度 */
    padding: 0 15px;          /* 减小按钮内边距 */
    font-size: 14px;          /* 调整按钮字体大小 */
  }

  :deep(.path-input .el-button i) {
    font-size: 16px;
  }

  .title {
    font-size: 20px;
  }

  .path-select {
    width: 100%;
  }
}

.yun-path {
  display: flex;
  align-items: center;
  gap: 8px;
}

.update-path {
  width: 200px;
}

/* 缩略图容器样式 */
.thumbnail-container {
  width: 80vw;
  max-width: 1200px;
  margin: 10px auto 0 auto;
  padding: 0px;
}

.thumbnail-list {
  display: flex;
  flex-wrap: nowrap;
  justify-content: center; /* 水平居中 */
  gap: 10px;
  overflow-x: auto;
  padding: 10px 0 ;
}

.thumbnail-item {
  flex: 0 0 auto;
  width: 100px;
}

.thumbnail-wrapper {
  position: relative;
  width: 100%;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  background: #fff;
}

.thumbnail-visual {
  position: relative;
  width: 100%;
 
  overflow: hidden;
}

.thumbnail-visual :deep(.el-image) {
  width: 100%;
  height: auto;
}

.thumbnail-visual :deep(img) {
  object-fit: cover;
  width: 100%;
  height: auto;
  object-fit: contain; /* 或者 cover，看你想要的效果 */

}

.thumbnail-actions {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(180deg, transparent, rgba(0, 0, 0, 0.65));
  opacity: 0;
  pointer-events: none;
  transition: opacity 0.3s ease;
}

.thumbnail-visual:hover .thumbnail-actions {
  opacity: 1;
  pointer-events: auto;
}

.thumbnail-actions .el-button {
  padding: 2px 10px;
}

.thumbnail-copy {
  width: 100%;
  height: 30px;
  border: none;
  background: rgba(45, 108, 223, 0.85);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.thumbnail-copy:hover {
  background: rgba(45, 108, 223, 1);
}

/* 移动端适配 */
@media screen and (max-width: 768px) {
  .thumbnail-container {
    width: 90vw;
  }

  .thumbnail-item {
    width: 80px;
  }
}
</style>