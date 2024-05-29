import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
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
      path: '/about/:id',
      name: 'about-others',
      component: () => import('../views/AboutView.vue')
    },
  ]
})

export default router
