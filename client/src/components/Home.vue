<template>
  <div class="home">
    <section class="banner">
      <div class="container has-text-centered">

        <h1 class="title is-size-1">Chingu</h1>
        <p class="subtitle is-size-4">We believe in free, high quality environments where you can work on your skills and become and rock star.</p>

        <div class="block">
          <a class="button is-primary is-large" style="color: #fff;">
            <span class="icon">
              <i class="fa fa-fire"></i>
            </span>
            <span>Apply</span>
          </a>
        </div>

      </div> <!-- container -->
    </section>

    <section class="hero is-dark">
      <div class="hero-body">
        <div class="container has-text-centered">
          <h1 class="title">How it works</h1>
          <p class="subtitle">We pair you with people at the same skill level, you build projects, gain experience, you become a rock star.</p>

          <div class="columns">

            <div class="column is-one-third contains-icon">
              <img src="/static/images/icons/svg/leaf.svg" onerror="/static/images/icons/png/leaf.png">
              <h2 class="is-size-2">1st</h2>
              <p>You apply to join our amazing community.</p>
            </div> <!-- column -->

            <div class="column is-one-third contains-icon">
              <img src="/static/images/icons/svg/leaves.svg" onerror="/static/images/icons/png/leaves.png">
              <h2 class="is-size-2">2nd</h2>
              <p>If you're accepted, you will join a cohort session and be paired in a team.</p>
            </div> <!-- column -->
            
            <div class="column is-one-third contains-icon">
              <img src="/static/images/icons/svg/tree.svg" onerror="/static/images/icons/png/tree.png">
              <h2 class="is-size-2">3rd</h2>
              <p>You build amazing projects with your team, then you decide if you want to continue with us or get a job (even both at the same time!)</p>
            </div> <!-- column -->

          </div> <!-- columns -->

        </div> <!-- container -->
      </div> <!-- hero-body -->
    </section>

    <section class="section">
      <div class="container">
        <h2 class="title">Battle-tested method</h2>
        <p class="subtitle">Thousands of developers have joined our community, hundreds have gotten jobs, here's what they say about their experience:</p>

        <div class="columns">

          <div class="column is-one-third">
            <div class="card">
              <div class="card-content">
                <p class="quote">
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
                <p class="quote">
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
                <p class="quote">
                  “Learn with other developers, you start as noob and become a pro with a family of devs.”
                </p>
                <p class="subtitle">
                  trion129
                </p>
              </div>
            </div>
          </div>

        </div> <!-- columns -->

      </div> <!-- container -->
    </section>

    <section class="hero is-dark">
      <div class="hero-body">
        <div class="container has-text-centered">
          <h2 class="title">Share your love</h2>
          <p class="subtitle">Running this community takes thousands of hours from our volunteers, allow us to continue doing what we love</p>

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
            <a class="button is-primary is-outlined is-medium">
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
export default {
  name: 'home'
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

.banner {
  color: #fff;
  background-color: #000;
  background: url(/static/images/home.jpg) no-repeat center center fixed;
  background-size: cover;
  height: 50vh;
}

.banner .title, 
.banner .subtitle {
  color: #fff;
}

.banner .title {
  text-shadow: 3px 3px #000;
}

.banner .subtitle {
  text-shadow: 1px 1px #000;
}

.banner .container {
  padding-top: 2rem;
}

.contains-icon img {
  display: block;
  margin: 0 auto;
  width: 256px;
  height: 256px;
}
</style>
