import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),//为了在github page跑专门改的hash模式 详细见https://router.vuejs.org/zh/guide/essentials/history-mode
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/explore',
      name: 'explore',
      component: () => import('../views/explore.vue')
    },
    {
      path: '/create',
      name: 'create',
      component: () => import('../views/create.vue'),
    },
    {
      path: '/detail/:id',
      name: 'detail',
      component: () => import('../views/DetailView.vue'),
    },
    {
      path: '/finish/:id',
      name: 'finish',
      component: () => import('../views/finish.vue'),
    },
    {
      path: '/travels/:id',
      name: 'travels',
      component: () => import('../views/travels.vue'),
    },
    {
      path: '/alltravels/',
      name: 'alltravels',
      component: () => import('../views/alltravels.vue'),
    },
    {
      path: '/about/:id',
      name: 'about-others',
      component: () => import('../views/AboutView.vue')
    },
  ]
})

export default router
