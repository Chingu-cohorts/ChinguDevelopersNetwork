<template>
<div>
  <section class="hero is-dark is-bold">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">New Post</h1>
        <h2 class="subtitle">Share your new great idea</h2>
      </div>
    </div>
  </section>

  <div class="container new-post">
    <div class="columns">
      <div class="column is-half is-offset-one-quarter">
        <div class="field">
          <label class="label">Title</label>
          <p class="control">
            <input class="input" type="text" v-model="post.title" placeholder="I got an awesome idea!">
          </p>
        </div>

        <markdown-editor v-model="post.content" ref="markdownEditor"></markdown-editor>

        <div class="field is-grouped">
          <p class="control">
            <button class="button is-primary" @click="savePost">Submit</button>
          </p>
          <p class="control">
            <button class="button is-link">Cancel</button>
          </p>
        </div>
      </div>
    </div>

  </div>
</div>
</template>

<script>
import { markdownEditor } from 'vue-simplemde'

export default {
  name: 'NewPost',

  data () {
    return {
      post: {
        title: '',
        content: ''
      }
    }
  },

  components: {
    markdownEditor
  },

  computed: {
    loggedUser () {
      return this.$store.state.user.loggedUser
    }
  },

  methods: {
    savePost (e) {
      let { title, content } = this.post

      let post = {
        title,
        content
      }

      this.$store.dispatch('CREATE_FORUM_POST', post).then(() => {
        this.$router.push({ name: 'PostList' })
      }).catch(err => {
        console.error(err)
      })
    }
  },

  mounted () {
    if (!this.loggedUser.username) {
      this.$router.push({ name: 'SignIn' })
    }
  }
}
</script>

<style scoped>
.new-post {
  margin-top: 2em;
  margin-bottom: 2em;
}
</style>
