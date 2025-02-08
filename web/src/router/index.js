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
      },{
        path: 'setting',
        name: 'Setting',
        component: () => import('../views/Setting.vue')
      }
    ]
  }
]
const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
