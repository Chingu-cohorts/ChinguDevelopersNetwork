<template>
<div>
  <section class="hero user-profile-background">
    <div class="hero-body">
      <div class="container has-text-centered">
        <figure class="image is-256x256">
          <img :src="userGravatar" class="avatar" :alt="currentUser.username">
        </figure>
        <h1 class="title" v-if="currentUser.first_name">{{ currentUser.first_name }} {{ currentUser.last_name }}</h1>
        <h2 class="subtitle">@{{ currentUser.username }}</h2>
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

  <section class="section feed">
    <div class="container">
      <div class="columns is-centered">

        <div class="column is-5">
          
          <h2>About</h2>
          <div class="about">
            <p>{{ userAbout }}</p>
          </div>

          <h2>Skills</h2>
          <div class="skills">
            <span class="tag is-medium">Go</span>
            <span class="tag is-medium">Vue.js</span>
            <span class="tag is-medium">Ruby on Rails</span>
            <span class="tag is-medium">Elixir</span>
            <span class="tag is-medium">JavaScript</span>
            <span class="tag is-medium">C/C++</span>
          </div>

        </div>

        <div class="column is-5">

          <h2>Projects</h2>
          <div class="projects">

          </div>

          <h2>Posts</h2>
          <div class="posts">

          </div>

        </div>
        
      </div> <!-- columns -->
    </div> <!-- container -->
  </section>
</div>
</template>

<script>
import { gravatar } from '@/utils'

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

<style lang="scss" scoped>
$font: 'Raleway', sans-serif;
$light_blue: #17c4fc;
$darker_blue: #1fa2ff;
$text_blue: #1dabfe;
$text_color: #fff;
$gray: #e9edef;

h1, h2, h3, h4, h5, h6 {
  font-family: $font;
}

.user-profile-background {
  background-image: linear-gradient(135deg, $light_blue 0%, $darker_blue 100%);
}

.feed h2 {
  color: $text_blue;
  font-size: 1.6rem;
  font-weight: 700;
  padding: 1rem 0;
}

.feed .tag {
  color: $text_color;
  background-color: $text_blue;
  margin: 0.5rem 0.5rem 0.5rem 0;
}

.feed .about {
  background: $gray;
  padding: 1rem;
}

.hero-body .title, .subtitle {
  color: $text_color;
}

.hero-body .fa {
  color: $text_color;
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
</style>
