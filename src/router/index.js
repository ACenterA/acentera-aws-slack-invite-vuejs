import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/views/layout/Layout'
export const constantRouterMap = [
  {
    path: '/',
    component: Layout,
    redirect: 'home',
    children: [
      {
        path: '/',
        component: () => import('@/views/home/index'),
        name: 'Home',
        hidden: true,
        meta: {
          title: 'home', noCache: true
        }
      }
    ]
  }
]

export const asyncRouterMap = [
  { path: '*', redirect: '/', hidden: true }
]

const createRouter = () => new Router({
  mode: 'history',
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
  // routes: []
})

const router = createRouter()

export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // the relevant part
  return newRouter
}

window.ResetRouter = resetRouter

export default router
