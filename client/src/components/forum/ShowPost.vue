<template>
  <div class="container">
    <template v-if="!currentPost.title">
      <spinner></spinner>
    </template>
    <template v-else>
      <div class="forum-post">
        <div class="columns">
          <div class="column is-hidden-mobile is-2-desktop">
            <div class="box has-text-centered user-box">
              <figure class="image is-128x128">
                <img :src="userGravatar">
              </figure>
              <div>
                <a class="username">{{ currentPost.user.username }}</a>
              </div>
            </div>
          </div>

          <div class="column is-12-mobile is-10-desktop">
            <div class="box">
              <div class="content">
                <h1 class="has-text-centered">{{ currentPost.title }}</h1>
                <hr>
                <div v-html="content">{{ currentPost.content }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>

    <new-comment></new-comment>
  </div>
</template>

<script>
import MarkdownIt from 'markdown-it'
import NewComment from '@/components/forum/NewComment'
import Spinner from '@/components/Spinner'
import { gravatar } from '@/components/utils'

export default {
  name: 'ShowPost',

  components: {
    NewComment,
    Spinner
  },

  methods: {
    stuff () {
      let { id, slug } = this.$route.params
      return id + slug
    }
  },

  computed: {
    currentPost () {
      return this.$store.state.forum.currentPost
    },

    userGravatar (state) {
      let { email } = state.currentPost.user
      return gravatar(email)
    },

    content (state) {
      let md = new MarkdownIt()
      return md.render(state.currentPost.content)
    }
  },

  mounted () {
    let { id } = this.$route.params
    this.$store.dispatch('LOAD_FORUM_POST', id)
  }
}
</script>

<style scoped>
hr {
  width: 40%;
  margin: 0 auto;
}

.forum-post {
  margin-top: 1em;
  margin-bottom: 1em;
}

figure img {
  border-radius: 50%;
  margin: 0.5em 0;
}

.user-box {
  border-radius: 0;
  box-shadow: none;
}

.username {
  font-weight: 700;
}
</style>
