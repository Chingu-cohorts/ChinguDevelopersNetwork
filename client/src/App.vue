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
import Navbar from '@/components/Navbar'
import MainFooter from '@/components/MainFooter'

let token = localStorage.getItem('token')

export default {
  name: 'app',

  components: {
    Navbar,
    MainFooter
  },

  mounted () {
    if (token) {
      this.$store.dispatch('LOAD_LOGGED_USER_DATA')
    }
  }
}
</script>

<style>
#app {
  font-family: 'Nunito', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

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

.avatar {
  border-radius: 100%;
  -webkit-border-radius: 100%;
  -moz-border-radius: 100%;
}
</style>
