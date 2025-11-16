import axios from 'axios' // 引入axios

const http = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 10000, // 请求超时时间（单位：毫秒）
})



// 请求拦截器
http.interceptors.request.use(
  (config) => {
    // 在发送请求前做些什么，例如添加 token
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
http.interceptors.response.use(
  (response) => {
    // 对响应数据做处理
    return response.data; // 直接返回 data，简化后续处理
  },
  (error) => {
    // 处理错误，例如 401 未授权
    if (error.response?.status === 401) {
      // 跳转登录页或提示
      console.error('未授权，请重新登录');
    }
    return Promise.reject(error);
  }
);



export default http;