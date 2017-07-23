<template>
<div>
  <section class="hero user-profile-background" v-bind:class="{ admin : currentUser.is_admin }">
    <div class="hero-body">
      <div class="container has-text-centered">
        <figure class="image is-256x256">
          <img :src="userGravatar" class="avatar" :alt="currentUser.username">
        </figure>
        <h1 class="title" v-if="currentUser.first_name">{{ currentUser.first_name }} {{ currentUser.last_name }}</h1>
        <h2 class="subtitle">{{ currentUser.username }}</h2>
        <a :href="userTwitter" v-if="currentUser.twitter_username">
          <i class="fa fa-twitter"></i>
        </a>
        <a :href="userGithub" v-if="currentUser.github_username">
          <i class="fa fa-github"></i>
        </a>
        <a :href="userMedium" v-if="currentUser.medium_username">
          <i class="fa fa-medium"></i>
        </a>
        <a :href="userLinkedin" v-if="currentUser.linkedin_username">
          <i class="fa fa-linkedin"></i>
        </a>
      </div>
    </div>
  </section>

  <section class="section">
    <div class="container">
      
      <nav class="level">
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Projects Completed</p>
            <p class="title">{{ userCompletedProjects }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Reputation</p>
            <p class="title">{{ userReputationLevel }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Recommended by</p>
            <p class="title">{{ userRecommendations }}</p>
          </div>
        </div>
      </nav>

      <hr>

      <div class="columns feed">

        <div class="column is-one-quarter">
          <div class="box">
            <div class="content">
              <h2>About</h2>
              <p>{{ userAbout }}</p>
            </div>
          </div>
        </div>

        <div class="column">

        </div>
      </div>
    </div>
  </section>
</div>
</template>

<script>
import { reputation, gravatar } from '@/components/utils'

export default {
  name: 'show-user',

  computed: {
    currentUser () {
      return this.$store.state.user.currentUser
    },

    userAbout (state) {
      if (state.currentUser.about) {
        return state.currentUser.about
      }
      return 'The user hasn\'t added information yet.'
    },

    userCompletedProjects (state) {
      if (state.currentUser.projects) {
        return state.currentUser.projects.length
      }
      return '0'
    },

    userReputationLevel (state) {
      let { experience } = state.currentUser
      return reputation(experience)
    },

    userRecommendations (state) {
      if (state.currentUser.recommendations) {
        return state.currentUser.recommendations.length
      }
      return '0'
    },

    userGravatar (state) {
      let { email } = state.currentUser
      return gravatar(email)
    },

    userTwitter (state) {
      let twitterUsername = state.currentUser.twitter_username
      return `https://twitter.com/${twitterUsername}`
    },

    userGithub (state) {
      let githubUsername = state.currentUser.github_username
      return `https://github.com/${githubUsername}`
    },

    userMedium (state) {
      let mediumUsername = state.currentUser.medium_username
      return `https://medium.com/@${mediumUsername}`
    },

    userLinkedin (state) {
      let linkedinUsername = state.currentUser.linkedin_username
      return `https://linkedin.com/${linkedinUsername}`
    }
  },

  mounted () {
    let { username } = this.$route.params
    this.$store.dispatch('LOAD_USER_DATA', username)
  }
}
</script>

<style scoped>
.user-profile-background {
  background-image: linear-gradient(135deg, #15df89 0%, #28a0dc 100%)
}

.admin {
  background-image: linear-gradient(135deg, #000 0%, #6f6f6f 100%)
}

.hero-body .title, .subtitle {
  color: #fff;
}

.hero-body .fa {
  color: #fff;
}

.is-256x256 {
  height: 256px;
  width: 256px;
}

.hero-body figure {
  display: inline-flex;
  justify-content: center;
}

.hero-body figure img {
  box-shadow: 0 3px 2px rgba(0, 0, 0, 0.3);
  transition: all 0.3s ease-in-out;
}

.hero-body figure img:hover {
  border-radius: 0%;
}

a .fa {
  margin: 0 0.3em;
}

hr {
  margin: 0 auto;
  width: 80%;
}

.feed {
  margin-top: 2em;
}

.feed .box {
  background-color: #fbfbfb;
}
</style>
