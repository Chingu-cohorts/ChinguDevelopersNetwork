<template>
<article class="media" v-if="loggedUser.id">
  <figure class="media-left">
    <p class="image is-64x64">
      <img :src="userGravatar">
    </p>
  </figure>
  <div class="media-content">
    <markdown-editor v-model="comment.content" ref="markdownEditor"></markdown-editor>
    <div class="field">
      <p class="control">
        <button class="button" @click="saveComment">Post comment</button>
      </p>
    </div>
  </div>
</article>
</template>

<script>
import { markdownEditor } from 'vue-simplemde'
import { gravatar } from '@/components/utils'

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

      this.$store.dispatch('CREATE_FORUM_COMMENT', comment).then(() => {
        this.$store.dispatch('LOAD_FORUM_POST', postId)
        this.comment.content = ''
      }).catch(err => {
        console.error(err)
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
