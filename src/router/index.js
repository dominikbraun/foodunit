import Vue from 'vue'
import Router from 'vue-router'
import main from '@/components/main'
import login from '@/components/login'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'FoodUnit',
      component: main
    },
    {
      path: '/login',
      name: 'Login',
      component: login
    }
  ]
})
