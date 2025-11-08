import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/index',
    component: () => import('../views/layout/layout.vue'),
    children: [
      {
        path: '/index',
        name: 'Index',
        component: () => import('../views/Index.vue')
      },
      {
        path: '/list',
        name: 'List',
        component: () => import('../views/List.vue')
      },
      {
        path: 'setting',
        name: 'Setting',
        component: () => import('../views/Setting.vue')
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/install',
    name: 'Install',
    component: () => import('../views/Install.vue')
  }
]
const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
