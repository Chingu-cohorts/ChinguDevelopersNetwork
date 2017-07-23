<template>
<div>
  <article class="media" v-if="loggedUser.id">
    <figure class="media-left is-hidden-mobile">
      <p class="image is-64x64">
        <img :src="userGravatar" class="avatar" :alt="loggedUser.username">
      </p>
    </figure>
    <div class="media-content">
      <markdown-editor v-model="comment.content" ref="markdownEditor"></markdown-editor>
      <div class="field">
        <p class="control" v-if="!loading">
          <button class="button is-success is-outlined" @click="handleCommentSubmit">Post comment</button>
        </p>
      </div>
    </div>
  </article>
  <p v-else>
    You must be logged in to be able to interact with others.
  </p>
</div>
</template>

<script>
import { markdownEditor } from 'vue-simplemde'
import { gravatar } from '@/components/utils'
import { http } from '@/api'

export default {
  name: 'new-comment',

  components: {
    markdownEditor
  },

  data () {
    return {
      comment: {
        content: ''
      },
      loading: false
    }
  },

  computed: {
    loggedUser () {
      return this.$store.state.user.loggedUser
    },

    userGravatar () {
      let { email } = this.loggedUser
      return gravatar(email)
    }
  },

  methods: {
    // Starts the process of saving a comment
    handleCommentSubmit (e) {
      e.preventDefault()

      let { content } = this.comment

      // If there's actually something in the comment
      if (content !== '') {
        this.loading = true

        let postId = this.$route.params.id

        let comment = {
          content,
          post_id: postId
        }

        // With the data in place, we can send the POST request
        this.saveComment(comment, postId)
      }
    },

    // Gets the comment and a postId and does a
    // POST request to the API endpoint
    saveComment (comment, postId) {
      http.post(`/posts/${comment.post_id}/comments`, {
        content: comment.content
      }).then(res => {
        this.$store.dispatch('LOAD_FORUM_POST', postId)
        this.resetState()
      }).catch(err => {
        console.error(err)
      })
    },

    // Reset everything to its original state
    resetState () {
      this.comment.content = ''
      this.loading = false
    }
  }
}
</script>

<style scoped>
article {
  margin-bottom: 2em;
}
</style>
