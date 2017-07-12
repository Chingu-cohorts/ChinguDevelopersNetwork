<template>
  <div>
    <h1>{{ currentCohort.name }}</h1>
  </div>
</template>

<script>
export default {
  name: 'show-cohort',

  computed: {
    currentCohort () {
      return this.$store.state.cohort.currentCohort
    }
  },

  mounted () {
    let { name } = this.$route.params
    this.$store.dispatch('LOAD_COHORT_DATA', name)
  }
}
</script>
