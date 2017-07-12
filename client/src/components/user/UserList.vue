<template>
  <section class="hero">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">
          Users
        </h1>
        <h2 class="subtitle">
          We got an incredible community
        </h2>

        <div class="columns is-multiline row">
          <individual-user
            v-for="user in users"
            :user="user"
            :key="user.id"
          ></individual-user>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import IndividualUser from './IndividualUser'

export default {
  name: 'user-list',

  components: {
    IndividualUser
  },

  computed: {
    users () {
      return this.$store.state.user.users
    }
  },

  mounted () {
    this.$store.dispatch('LOAD_USERS_LIST')
  }
}
</script>
