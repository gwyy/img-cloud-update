import { createRouter, createWebHistory } from 'vue-router'
import { useTokenStore } from '@/store'
import { isInstall } from '@/api/system'

const routes = [
  {
    path: '/',
    redirect: '/index',
    component: () => import('../views/layout/layout.vue'),
    children: [
      {
        path: '/index',
        name: 'Index',
        component: () => import('../views/Index.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: '/list',
        name: 'List',
        component: () => import('../views/List.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: 'setting',
        name: 'Setting',
        component: () => import('../views/Setting.vue'),
        meta: { requiresAuth: true ,isInstall: true }
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: { requiresAuth: false ,isInstall: true }
    },
  {
    path: '/install',
    name: 'Install',
    component: () => import('../views/Install.vue'),
    meta: { requiresAuth: false,isInstall: false }
  }
]
const router = createRouter({
  history: createWebHistory(),
  routes
})

//导航守卫
router.beforeEach(async (to, from, next) => { 
  //如果不需要安装，直接放行
  if (to.meta.isInstall === false) {
    return next();
  }
  // 获取tokenStore实例
  const tokenStore = useTokenStore()
 
  //判断是否安装
  const res = await isInstall()
  if (res.code === 0) {
      if (res.data.isInstall === "false") {
        tokenStore.logout()
        next({ name: 'Install' ,query: { redirect: to.fullPath }}); // 去安装
    } else {
        // 如果路由不需要认证，直接放行
        if (!to.meta.requiresAuth) {
          return next();
        }
      //判断是否登录，登录后初始化表单
      if (!tokenStore.isLogin()) {  
        next({ name: 'Login' ,query: { redirect: to.fullPath }}); // 未登录，跳转到登录页，
      }
    }
  } else { //安装报错
    //无脑退出登录
    tokenStore.logout()
    alert(res.message)
    next({ name: 'Install',query: { redirect: to.fullPath } }); // 去安装
  }
  next();
})


export default router
