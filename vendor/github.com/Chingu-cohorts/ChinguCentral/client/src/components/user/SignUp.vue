<template>
<div>
  <section class="hero is-primary">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">Sign Up</h1>
        <h2 class="subtitle">Register to start applying to cohorts</h2>
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
            <p v-if="user.username.length < 4" class="help is-danger">
              Username must be at least 4 characters long
            </p>
          </div>
          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="text" placeholder="Email" v-model="user.email">
              <span class="icon is-small is-left">
                <i class="fa fa-envelope"></i>
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
            <p v-if="user.password.length < 6" class="help is-danger">
              Password must be at least 6 characters long
            </p>
          </div>
          <div class="field">
            <p class="control">
              <button class="button is-primary is-outlined" type="submit" v-on:click="registerUser">
                Register
              </button>
            </p>
          </div>
          <p>By registering you agree to our <a>Terms of Service</a>.</p>
        </div>
      </div>
    </div>
  </section>
</div>
</template>

<script>
export default {
  name: 'sign-up',

  data () {
    return {
      user: {
        username: '',
        email: '',
        password: ''
      }
    }
  },

  methods: {
    registerUser (e) {
      e.preventDefault()
      let { username, email, password } = this.user

      let user = {
        username,
        email,
        password
      }

      this.$store.dispatch('POST_REGISTRATION_DATA', user).then(() => {
        this.$router.push({name: 'SignIn'})
      }).catch(err => {
        console.error(err)
      })
    }
  }
}
</script>
