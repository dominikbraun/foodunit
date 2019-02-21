import Vue from 'vue'
import Router from 'vue-router'
import FoodUnit from '@/components/main'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'FoodUnit',
      component: FoodUnit
    }
  ]
})
