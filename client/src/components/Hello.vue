<template>
  <div>
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">Cohorts</h1>
          <h2 class="subtitle">Join a community and take your skills to the next level</h2>

          <div class="columns is-multiline row">
            <individual-cohort
              v-for="cohort in cohorts"
              :cohort="cohort"
              :key="cohort.id"
            ></individual-cohort>
          </div>

          <div class="block has-text-centered">
            <a class="button is-primary is-inverted is-outlined is-large">
              <span class="icon">
                <i class="fa fa-fire"></i>
              </span>
              <span>Apply</span>
            </a>
          </div>

        </div>
      </div>
    </section>

    <section class="section">
      <div class="container">
        <h1 class="title">What devs are saying</h1>
        <h2 class="subtitle">
          Hundreds of developers have joined <strong>Chingu</strong>, here's what they say about their experience
        </h2>

        <div class="columns">
          
          <div class="column is-one-third">
            <div class="card">
              <div class="card-content">
                <p class="title quote">
                  “An amazing collaborative experience beyond my expectations.”
                </p>
                <p class="subtitle">
                  periman2
                </p>
              </div>
            </div>
          </div>

          <div class="column is-one-third">
            <div class="card">
              <div class="card-content">
                <p class="title quote">
                  “I honestly don't believe I've seen any collaboration environment produce this much overall "learning" ever, really.”
                </p>
                <p class="subtitle">
                  p1xt
                </p>
              </div>
            </div>
          </div>

          <div class="column is-one-third">
            <div class="card">
              <div class="card-content">
                <p class="title quote">
                  “Learn with other developers, you start as noob and become a pro with a family of devs.”
                </p>
                <p class="subtitle">
                  trion129
                </p>
              </div>
            </div>
          </div>

        </div>
      </div>
    </section>

    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container has-text-centered">
          <h1 class="title">Share your love</h1>
          <h2 class="subtitle">Running this community is expensive, help us with a donation</h2>

          <div class="field">
            <label class="label">Amount</label>
            <p class="control">
              <span class="select">
                <select>
                  <option>$5</option>
                  <option>$10</option>
                  <option>$25</option>
                  <option>$50</option>
                  <option>$100</option>
                </select>
              </span>
            </p>
          </div>

          <div class="block has-text-centered">
            <a class="button is-primary is-inverted is-outlined is-medium">
              <span class="icon">
                <i class="fa fa-check"></i>
              </span>
              <span>Donate</span>
            </a>
          </div>

        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import IndividualCohort from './cohort/IndividualCohort'

export default {
  name: 'hello',

  components: {
    IndividualCohort
  },

  computed: mapState([
    'cohorts'
  ]),

  mounted () {
    this.$store.dispatch('LOAD_COHORTS_LIST')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.quote {
  font-size: 1.5rem;
}

.field .label {
  color: #fff;
}

.field .control {
  text-align: center;
}
</style>
