import http from '@/utils/request'


// 上传图片
export const uploadImage = (data) => {
  return http({
    url: "/base/uploadImage",
    method: "POST",
    data,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

// 删除图片
export const deleteImage = (name) => {
  return http({
    url: "/base/deleteImage",
    method: "POST",
    data: { name: name },
  })
}

// 获取阿里云密钥
export const getAliyunSecret = () => {
  return http({
    url: "/base/getAliyunSecret",
    method: "GET",
  })
}

// 保存阿里云密钥
export const saveAliyunSecret = (data) => {
  return http({
    url: "/base/saveAliyunSecret",
    method: "POST",
    data,
  })
}