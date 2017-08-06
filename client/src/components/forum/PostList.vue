<template>
<div>
  <section class="hero is-small is-dark is-bold">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">Forum</h1>
        <h2 class="subtitle">Discuss ideas that matter</h2>
      </div>
    </div>
  </section>
  <div class="container">
    <div class="columns forum-container">

      <div class="column is-4">
        <template v-if="loggedUser.username">
          <router-link
            :to="{ name: 'NewPost' }"
            class="box button is-primary is-outlined is-medium new-post"
          >New Post</router-link>
        </template>
        <div class="box member-of-the-day">
          <h2 class="has-text-centered is-primary">Important</h2>
          <p>During the beta period try to break as many things as you can.</p>
        </div>
      </div>


      <div class="column is-8">
        <individual-post
          v-for="post in posts"
          :post="post"
          :key="post.id">
        </individual-post>
      </div>

    </div>
  </div>
</div>
</template>

<script>
import IndividualPost from '@/components/forum/IndividualPost'

export default {
  name: 'PostList',

  data () {
    return {
      interval: null
    }
  },

  components: {
    IndividualPost
  },

  computed: {
    posts () {
      return this.$store.state.forum.posts
    },

    loggedUser () {
      return this.$store.state.user.loggedUser
    }
  },

  methods: {
    loadData () {
      this.$store.dispatch('LOAD_POSTS_LIST')
    }
  },

  mounted () {
    this.$store.dispatch('LOAD_POSTS_LIST')

    this.interval = setInterval(() => this.loadData(), 20000)
  },

  beforeDestroy () {
    clearInterval(this.interval)
  }
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
