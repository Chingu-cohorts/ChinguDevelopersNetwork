<template>
<nav class="navbar">
  <div class="navbar-brand">
    <router-link
      :to="{name: 'Hello'}"
      class="navbar-item">
      <img src="https://i.imgur.com/CH034bi.png" alt="The Zoo">
    </router-link>

    <a class="navbar-item is-hidden-desktop" href="https://github.com/Chingu-cohorts/ChinguCentral" target="_blank">
      <span class="icon" style="color: #333;">
        <i class="fa fa-github"></i>
      </span>
    </a>

    <a class="navbar-item is-hidden-desktop" href="https://twitter.com/ChinguCollabs" target="_blank">
      <span class="icon" style="color: #55acee;">
        <i class="fa fa-twitter"></i>
      </span>
    </a>

    <div class="navbar-burger burger" data-target="navMenu" @click="collapse">
      <span></span>
      <span></span>
      <span></span>
    </div>
  </div>

  <div id="navMenu" class="navbar-menu" v-bind:class="{'is-active': collapsed}">
    <div class="navbar-start">
      <router-link
        :to="{ name: 'CohortList'}"
        exact-active-class="is-active"
        class="navbar-item">Cohorts</router-link>
      <router-link
        :to="{ name: 'UserList'}"
        exact-active-class="is-active"
        class="navbar-item">Users</router-link>
      <router-link
        :to="{ name: 'PostList'}"
        exact-active-class="is-active"
        class="navbar-item">Forum</router-link>
    </div>

    <div class="navbar-end">
      <template v-if="loggedUser.username">
        <a class="navbar-item" @click="reloadUser">{{ loggedUser.username }}</a>
        <a class="navbar-item" @click="logout">Logout</a>
      </template>
      <template v-else>
        <router-link
          :to="{name: 'SignIn'}"
          class="navbar-item">Login</router-link>
        <router-link
          :to="{name: 'SignUp'}"
          class="navbar-item">Register</router-link>
      </template>
    </div>
  </div>
</nav>
</template>

<script>
export default {
  name: 'navbar',

  data () {
    return {
      collapsed: false
    }
  },

  computed: {
    loggedUser () {
      return this.$store.state.user.loggedUser
    }
  },

  methods: {
    collapse (e) {
      this.collapsed = !this.collapsed
    },

    // You will see, young padawan, only changing the route
    // doesn't actually re render the component, so we need
    // to change the route and manually dispatch the mutation
    // in order to get the current route data
    reloadUser (e) {
      e.preventDefault()

      let { username } = this.$store.state.user.loggedUser

      this.$router.push({name: 'ShowUser', params: { username }})
      this.$store.dispatch('LOAD_USER_DATA', username)
    },

    logout (e) {
      this.$store.dispatch('LOGOUT_USER')
      window.location = '/'
    }
  }
}
</script>

<style scoped>
.navbar {
  border-bottom: 1px solid #f9f9f9;
}
</style>
