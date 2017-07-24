<template>
<nav class="navbar">
  <div class="navbar-brand">
    <router-link
      :to="{name: 'Hello'}"
      class="navbar-item">
      <img src="/static/images/logo.png" alt="Chingu Developers Network">
    </router-link>

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
        <a class="navbar-item" @click="reloadUser">
          <figure class="image is-32x32">
            <img :src="userGravatar" class="avatar">
          </figure>
        </a>
        <router-link :to="{ name: 'Settings' }" class="navbar-item">Settings</router-link>
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
import { gravatar } from '@/components/utils'

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
    },

    userGravatar (state) {
      let { email } = state.loggedUser
      return gravatar(email)
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

.navbar-item img {
  max-height: none;
}

.navbar-brand .navbar-item img {
  max-height: 1.75rem;
} 
</style>
