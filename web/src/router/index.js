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
      component: () => import('../views/myself.vue')
    },
    {
      path: '/explore',
      name: 'explore',
      component: () => import('../views/explore.vue')
    },
    {
      path: '/create',
      name: 'create',
      component: () => import('../views/CreatePage.vue'),
    },
    {
      path: '/createplan/:id',
      name: 'createplan',
      component: () => import('../views/Createplan.vue'),
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
      path: '/alltravelshave/',
      name: 'alltravelshave',
      component: () => import('../views/alltravelshave.vue'),
    },
    {
      path: '/alltravels/',
      name: 'alltravels',
      component: () => import('../views/alltravels.vue'),
    },
    {
      path: '/about/:id',
      name: 'about-others',
      component: () => import('../views/userpage.vue')
    },
    {
      path: '/itinerarys/:id',
      name: 'itinerary',
      component: () => import('../views/itinerary.vue'),
      meta: { keepAlive: false }
    },
    {
      path: '/myitinerarys/:id',
      name: 'myitinerary',
      component: () => import('../views/myitinerary.vue'),
      meta: { keepAlive: false }
    },
    {
      path: '/partys/:id',
      name: 'party',
      component: () => import('../views/party.vue'),
      meta: { keepAlive: false }
    },
    {
      path: '/merplan/:id',
      name: 'merplan',
      component: () => import('../views/merplan.vue'),
      meta: { keepAlive: false }
    },
    {
      path: '/member/:id',
      name: 'member',
      component: () => import('../views/member.vue'),

    },
    {
      path: '/ifollow/:id',
      name: 'ifollow',
      component: () => import('../views/ifollow.vue'),

    },
    {
      path: '/followme/:id',
      name: 'followme',
      component: () => import('../views/followme.vue'),

    },
  ]
})

export default router
