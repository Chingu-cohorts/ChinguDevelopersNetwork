<template>
  <h1>{{ stuff() }}</h1>
</template>

<script>
export default {
  name: 'ShowPost',

  methods: {
    stuff () {
      let { id, slug } = this.$route.params
      return id + slug
    }
  }
}
</script>
