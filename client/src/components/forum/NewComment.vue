<template>
<div>
  <article class="media" v-if="loggedUser.id">
    <figure class="media-left">
      <p class="image is-64x64">
        <img :src="userGravatar" class="avatar" :alt="loggedUser.username">
      </p>
    </figure>
    <div class="media-content">
      <markdown-editor v-model="comment.content" ref="markdownEditor"></markdown-editor>
      <div class="field">
        <p class="control">
          <button class="button is-success is-outlined" @click="saveComment">Post comment</button>
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
      }
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
    saveComment (e) {
      e.preventDefault()

      let postId = this.$route.params.id

      let { content } = this.comment

      let comment = {
        content,
        post_id: postId
      }

      http.post(`/posts/${comment.post_id}/comments`, {
        content: comment.content
      }).then(res => {
        console.log(res.data)
        this.$store.dispatch('LOAD_FORUM_POST', postId)
        this.comment.content = ''
      }).catch(err => {
        if (err.response) {
          console.log(err.response.data)
        } else {
          console.error(err)
        }
      })
    }
  }
}
</script>

<style scoped>
article {
  margin-bottom: 2em;
}
</style>
