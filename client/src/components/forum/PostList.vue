<template>
<div>
  <section class="hero is-medium is-dark is-bold">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">Forum</h1>
        <h2 class="subtitle">Discuss ideas that matter</h2>
      </div>
    </div>
  </section>
  <div class="container">
    <div class="columns forum-container">
      <div class="column is-8">
        <individual-post></individual-post>
        <individual-post></individual-post>
      </div>
      <div class="column is-4">
        <a class="box button is-primary is-outlined is-medium new-post">
          <span class="icon is-small">
            <i class="fa fa-check"></i>
          </span>
          <span>New Post</span>
        </a>
        <div class="box member-of-the-day">
          <h2 class="has-text-centered is-primary">Member of the day</h2>
          <img class="image is-128x128" src="https://gravatar.com/avatar/3196ff4ca2acf9464de8a7c5b5e62957?s=256">
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import { mapState } from 'vuex'
import IndividualPost from './IndividualPost'

export default {
  name: 'PostList',

  components: {
    IndividualPost
  },

  computed: mapState([
    'posts'
  ])

  /*
  mounted () {
    this.$store.dispatch('LOAD_POSTS_LIST')
  }
  */
}
</script>

<style scoped>
.forum-container {
  margin-top: 2em;
  margin-bottom: 2em;
}

.is-4 .box h2 {
  font-weight: 700;
}

.member-of-the-day img {
  border-radius: 50%;
  margin: 0 auto;
}

.new-post {
  display: flex;
}
</style>
