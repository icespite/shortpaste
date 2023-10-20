import {
  createRouter,
  createWebHistory,
  type NavigationGuard,
  type NavigationGuardNext,
  type RouteLocationNormalized
} from 'vue-router'

import LoginView from '@/views/LoginView.vue'
import LinkView from '@/views/LinkView.vue'
import TextView from '@/views/TextView.vue'
import FileView from '@/views/FileView.vue'
import MainView from '@/views/MainView.vue'
import NotFoundView from '@/views/NotFoundView.vue'

import { store } from '@/store'

const routes = [
  {
    path: '/login',
    name: 'login',
    component: LoginView,
    meta: { requiresAuth: false }
  },
  {
    path: '/',
    name: 'main',
    redirect: '/link',
    component: MainView,
    meta: { requiresAuth: true },
    children: [
      {
        path: 'link',
        name: 'links',
        component: LinkView
      },
      {
        path: 'text',
        name: 'texts',
        component: TextView
      },
      {
        path: 'file',
        name: 'files',
        component: FileView
      }
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFoundView
  }
]

const isAuthenticatedGuard: NavigationGuard = (
  to: RouteLocationNormalized,
  _from: RouteLocationNormalized,
  next: NavigationGuardNext
): void => {
  const isAuthenticated = store.getters.isAuthenticated

  if (to.matched.some((record) => record.meta.requiresAuth) && !isAuthenticated) {
    // Redirect to /login if not logged in
    next('/login')
  } else if (!to.matched.some((record) => record.meta.requiresAuth) && isAuthenticated) {
    // Redirect to / if already logged in
    next('/')
  } else {
    next()
  }
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

router.beforeEach(isAuthenticatedGuard)

export default router
