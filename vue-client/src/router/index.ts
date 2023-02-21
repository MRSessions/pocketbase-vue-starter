// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/default/Default.vue'),
    redirect: '/home',
    children: [
      {
        path: '/home',
        name: 'Home',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "home" */ '@/views/Home.vue'),
      },
      {
        path: '/about',
        name: 'About',
        component: () => import(/* webpackChunkName: "about" */ '@/views/About.vue'),
      },
      {
        path: '/contact',
        name: 'Contact',
        component: () => import(/* webpackChunkName: "contact" */ '@/views/Contact.vue'),
      },
      {
        path: '/admin',
        component: () => import('@/App.vue'),
        children: [
          {
            path: '/admin/setup-admin',
            name: 'SetupAdmin',
            component: () => import('@/views/admin/SetupAdmin.vue'),
          },
        ]
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/home',
    // name: 'NotFound',
    // component: () => import('@/views/NotFound.vue'),
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
