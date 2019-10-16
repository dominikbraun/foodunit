<template>
  <div class="row m-0 h-100">
    <SideNav />
    <SideCart v-bind="" />
  </div>
</template>

<script>
import SideNav from './components/SideNav.vue'
import SideCart from './components/SideCart.vue';

export default {
  name: 'app',
  components: {
      SideCart,
      SideNav
  }
}
</script>
