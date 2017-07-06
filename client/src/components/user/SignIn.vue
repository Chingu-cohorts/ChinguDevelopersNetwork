<template>
<div>
  <section class="hero is-primary">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">Sign In</h1>
        <h2 class="subtitle">Start interacting with other members</h2>
      </div>
    </div>
  </section>

  <section class="section">
    <div class="columns">
      <div class="container">
        <div class="column is-half is-offset-one-quarter">
          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="text" placeholder="Username" v-model="user.username">
              <span class="icon is-small is-left">
                <i class="fa fa-user"></i>
              </span>
            </p>
          </div>
          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="password" placeholder="Password" v-model="user.password">
              <span class="icon is-small is-left">
                <i class="fa fa-lock"></i>
              </span>
            </p>
          </div>
          <div class="field">
            <p class="control">
              <button class="button is-primary is-outlined" type="submit" v-on:click="loginUser">
                Login
              </button>
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</div>
</template>

<script>
export default {
  name: 'sign-in',

  data () {
    return {
      user: {
        username: '',
        password: ''
      }
    }
  },

  methods: {
    loginUser (e) {
      e.preventDefault()

      let {username, password} = this.user

      let user = {
        username,
        password
      }

      this.$store.dispatch('POST_LOGIN_DATA', user).then(() => {
        window.location = '/'
      }).catch(err => {
        console.error(err)
      })
    }
  }
}
</script>
