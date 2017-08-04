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

          <p class="has-text-centered">{{ error }}</p>

          <spinner v-if="loading"></spinner>

          <div class="field" v-if="!loading">
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

import { http } from '@/api'
import Spinner from '@/components/misc/Spinner'

export default {
  name: 'sign-in',

  data () {
    return {
      user: {
        username: '',
        password: ''
      },
      error: '',
      loading: false
    }
  },

  components: {
    Spinner
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

      this.loading = true

      let {username, password} = this.user

      let user = {
        username,
        password
      }

      // If the user introduced an username and a password we can proceed
      let isUserValid = this.verifyUser(user)

      if (isUserValid) {
        this.handleLogin(user)
      } else {
        this.error = 'Username and password are required'
        this.loading = false
      }
    },

    handleLogin (user) {
      http.post('/users/login', {
        ...user
      }).then(res => {
        // The login was successful, set the token and redirect
        this.loading = false
        localStorage.setItem('token', res.data.token)
        window.location = '/'
      }).catch(err => {
        // Something went wrong
        if (err.response) {
          let { code, message } = err.response.data

          // 404 no user with that username
          // 401 wrong password
          if (code === 400) {
            this.error = message
          } else if (code === 401) {
            this.error = message
          }

          console.log(err)
        } else {
          console.error(err)
        }

        this.loading = false
      })
    },

    // Returns true if the user has valid data
    verifyUser (user) {
      if (user.username !== '' && user.password !== '') {
        return true
      }

      return false
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
