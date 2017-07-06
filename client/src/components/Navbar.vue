<template>
  <nav class="nav has-shadow">
    <div class="container">
      <div class="nav-left">
        <router-link
          :to="{name: 'Hello'}"
          class="nav-item">
          <img src="https://i.imgur.com/CH034bi.png" alt="The Zoo">
        </router-link>
        <router-link
          :to="{ name: 'CohortList'}"
          exact-active-class="is-active"
          class="nav-item is-tab is-hidden-mobile">
          Cohorts
        </router-link>
        <router-link
          :to="{ name: 'UserList'}"
          exact-active-class="is-active"
          class="nav-item is-tab is-hidden-mobile">
          Users
        </router-link>
        <a class="nav-item is-tab is-hidden-mobile">Forums</a>
        <a class="nav-item is-tab is-hidden-mobile">Projects</a>
      </div>
      <span class="nav-toggle" v-on:click="collapse">
        <span></span>
        <span></span>
        <span></span>
      </span>
      <div class="nav-right nav-menu" v-bind:class="{'is-active': collapsed}">
        <router-link
          :to="{name: 'Hello'}"
          class="nav-item is-tab is-hidden-tablet"
          exact-active-class="is-active">
          Home
        </router-link>
        <router-link
          :to="{name: 'CohortList'}"
          class="nav-item is-tab is-hidden-tablet"
          exact-active-class="is-active">
          Cohorts
        </router-link>
        <router-link
          :to="{name: 'UserList'}"
          class="nav-item is-tab is-hidden-tablet"
          exact-active-class="is-active">
          Users
        </router-link>
        <a class="nav-item is-tab is-hidden-tablet">Forums</a>
        <a class="nav-item is-tab is-hidden-tablet">Projects</a>

        <template v-if="loggedUser.username">
          <a class="nav-item is-tab" @click="reloadUser">{{ loggedUser.username }}</a>
        </template>
        <template v-else>
          <router-link
            :to="{name: 'SignIn'}"
            class="nav-item is-tab">
            Login
          </router-link>
          <router-link
            :to="{name: 'SignUp'}"
            class="nav-item is-tab">
            Register
          </router-link>
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
      return this.$store.state.loggedUser
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

      let { username } = this.$store.state.loggedUser

      this.$router.push({name: 'ShowUser', params: { username }})
      this.$store.dispatch('LOAD_USER_DATA', username)
    }
  }
}
</script>
