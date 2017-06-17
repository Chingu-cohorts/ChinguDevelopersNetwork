<template>
  <div id="app">
    <navbar></navbar>
    <router-view></router-view>
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
</style>
