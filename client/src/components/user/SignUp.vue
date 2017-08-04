<template>
<section class="section signup">
  <div class="columns">
    <div class="container">
      <div class="column is-half is-offset-one-quarter signup-container">

        <h1 class="title has-text-centered">Time to make history</h1>

        <div class="form-container">
          <p class="has-text-centered description">Take a new path, join a community like no one, do the unimaginable.</p>

          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="text" placeholder="First name" v-model="user.first_name">
              <span class="icon is-small is-left">
                <i class="fa fa-address-book"></i>
              </span>
            </p>
            <p class="help is-warning" v-if="!user.first_name">Optional</p>
          </div>

          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="text" placeholder="Last name" v-model="user.last_name">
              <span class="icon is-small is-left">
                <i class="fa fa-grav"></i>
              </span>
            </p>
            <p class="help is-warning" v-if="!user.last_name">Optional</p>
          </div>

          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="text" placeholder="Username" v-model="user.username" required>
              <span class="icon is-small is-left">
                <i class="fa fa-user"></i>
              </span>
            </p>
            <p class="help is-primary" v-if="!user.username">Required</p>
          </div>

          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="text" placeholder="Email" v-model="user.email" required>
              <span class="icon is-small is-left">
                <i class="fa fa-envelope"></i>
              </span>
            </p>
            <p class="help is-primary" v-if="!user.email">Required</p>
          </div>

          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="password" placeholder="Password" v-model="user.password" required>
              <span class="icon is-small is-left">
                <i class="fa fa-lock"></i>
              </span>
            </p>
            <p class="help is-primary" v-if="!user.password">Required</p>
          </div>

          <p class="has-text-centered">{{ error }}</p>

          <spinner v-if="loading"></spinner>

          <div class="field" v-if="!loading">
            <p class="control">
              <button class="button is-primary" type="submit" @click="registerUser">
                Register
              </button>
            </p>
          </div>

          <p class="has-text-centered">By registering you agree to our <router-link :to="{ name: 'TermsOfService' }">Terms of Service</router-link></p>

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
  name: 'sign-up',

  data () {
    return {
      user: {
        first_name: '',
        last_name: '',
        username: '',
        email: '',
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
    registerUser (e) {
      e.preventDefault()

      this.loading = true

      let user = this.user
      let isValid = this.validateUser(user)

      if (isValid) {
        this.handleRegistration(user)
      } else {
        this.error = 'Username, email and password are required'
      }

      this.loading = false
    },

    handleRegistration (user) {
      http.post('/users', {
        ...user
      }).then(res => {
        this.$router.push({name: 'SignIn'})
        console.log(res)
      }).catch(err => {
        if (err.response.data.message) {
          this.error = err.response.data.message
        } else {
          console.error(err)
        }
      })
    },

    // If you know of a better way to do it
    // please let me know
    // P.S. this is validated server side so don't worry
    validateUser (user) {
      return user.email !== '' && user.username !== '' && user.password !== ''
    }
  }
}
</script>

<style scoped>
.signup {
  background: url(/static/images/signup-bg.jpg) no-repeat center center fixed;
  background-size: cover;
  height: 100%;
}

.signup-container {
  background-color: rgba(0, 0, 0, .3);
  border-radius: 0.5em;
  margin-top: 2em;
  color: #f9f9f9;
}

.signup-container h1 {
  font-weight: 700;
  text-transform: uppercase; 
  color: #f9f9f9;
}

.signup-container .fa {
  color: #f9f9f9;
}

.signup-container input {
  background-color: rgba(255, 255, 255, .15);
  border: none;
  box-shadow: none;
  border-radius: 0;
  color: #f9f9f9;
  text-shadow: -2px 0px 1px #000;
}

.signup-container input:focus {
  background-color: rgba(255, 255, 255, .3);
  border: 1px solid #15df89;
}

.signup-container input::-webkit-input-placeholder {
  color: #f9f9f9;
}

.signup-container .form-container {
  width: 80%;
  margin: 0 auto;
}

.signup-container .form-container .description {
  margin-bottom: 1em;
}

.signup-container .form-container .button {
  display: flex;
  margin: 0 auto;
}

.control.has-icons-left .input,
.control.has-icons-left .select select {
  padding-left: 2.4em;
}
</style>
