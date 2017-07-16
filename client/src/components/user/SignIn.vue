<template>
<section class="section signin">
  <div class="columns">
    <div class="container">
      <div class="column is-half is-offset-one-quarter signin-container">
        <h1 class="title has-text-centered">Login</h1>
        <div class="form-container">
          <p class="has-text-centered description">Start interacting with the community.</p>
          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="text" placeholder="Username" v-model="user.username" required>
              <span class="icon is-small is-left">
                <i class="fa fa-user"></i>
              </span>
            </p>
          </div>
          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="password" placeholder="Password" v-model="user.password" required>
              <span class="icon is-small is-left">
                <i class="fa fa-lock"></i>
              </span>
            </p>
          </div>
          <div class="field">
            <p class="control">
              <button class="button is-primary" type="submit" @click="loginUser">
                Login
              </button>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>
</template>

<script>
import { mapState } from 'vuex'

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

  computed: mapState([
    'loggedUser'
  ]),

  mounted () {
    if (this.loggedUser) {
      this.$router.push({ name: 'Hello' })
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

<style scoped>
.signin {
  background: url(/static/images/signin-bg.jpg) no-repeat center center fixed;
  background-size: cover;
  height: 100%;
  min-height: 80vh;
}

.signin-container {
  background-color: rgba(0, 0, 0, .3);
  border-radius: 0.5em;
  margin-top: 2em;
  color: #f9f9f9;
}

.signin-container h1 {
  font-weight: 700;
  text-transform: uppercase; 
  color: #f9f9f9;
}

.signin-container .fa {
  color: #f9f9f9;
}

.signin-container input {
  background-color: rgba(255, 255, 255, .15);
  border: none;
  box-shadow: none;
  border-radius: 0;
  color: #f9f9f9;
  text-shadow: -2px 0px 1px #000;
}

.signin-container input:focus {
  background-color: rgba(255, 255, 255, .3);
  border: 1px solid #15df89;
}

.signin-container input::-webkit-input-placeholder {
  color: #f9f9f9;
}

.signin-container .form-container {
  width: 80%;
  margin: 0 auto;
}

.signin-container .form-container .description {
  margin-bottom: 1em;
}

.signin-container .form-container .button {
  display: flex;
  margin: 0 auto;
}

.control.has-icons-left .input,
.control.has-icons-left .select select {
  padding-left: 2.4em;
}
</style>
