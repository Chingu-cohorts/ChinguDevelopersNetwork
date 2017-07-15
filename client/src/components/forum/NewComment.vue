<template>
  <div>
    <template v-if="!loggedUser.id">
      <em>You must login to be able to comment</em>
    </template>
    <template v-else>
      <markdown-editor v-model="comment.content" ref="markdownEditor"></markdown-editor>
    </template>
  </div>
</template>

<script>
import { markdownEditor } from 'vue-simplemde'

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
    }
  }
}
</script>
