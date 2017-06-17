<template>
  <section class="hero">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">
          Users
        </h1>
        <h2 class="subtitle">
          We got hundreds of members
        </h2>

        <div class="columns is-multiline row">
          
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'user-list',

  computed: mapState([
    'users'
  ]),

  mounted () {
    this.$store.dispatch('LOAD_USERS_LIST')
  }
}
</script>
