<template>
  <div id="app">
    <navbar></navbar>
    <transition name="fade">
      <router-view></router-view>
    </transition>
    <main-footer></main-footer>
  </div>
</template>

<script>
import Navbar from './components/Navbar'
import MainFooter from './components/MainFooter'

export default {
  name: 'app',
  components: {
    Navbar,
    MainFooter
  }
}
</script>

<style>
.fade-enter-active, .fade-leave-active {
  transition-property: opacity;
  transition-duration: .25s;
}

.fade-enter-active {
  transition-delay: .25s;
}

.fade-enter, .fade-leave-active {
  opacity: 0
}
</style>
