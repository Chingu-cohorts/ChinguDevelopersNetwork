<template>
  <div class="home">
    <section class="chingu">
      <div class="container has-text-centered">

        <div class="image-header">
          
        </div>
        
        <div class="columns">

          <div class="column is-one-third">
            <h3>join a cohort</h3>
          </div>

          <div class="column is-one-third">
            <h2>meet your teammates</h2>
            <p>Meet the people you'll spend the next few weeks working with on a build-to-learn-project!</p>
          </div>

          <div class="column is-one-third">
            <h3>build your project</h3>
          </div>

        </div> <!-- columns -->

      </div> <!-- container -->
    </section>

    <h2 class="has-text-centered">what devs are saying</h2>

    <section class="section chingu">
      <div class="container">

        <div class="columns">

          <div class="column is-one-third">
            <div class="card">
              <div class="card-content has-text-centered">
                <p class="quote">
                  “An amazing collaborative experience beyond my expectations.”
                </p>
                <p class="subtitle">
                  - periman2
                </p>
              </div>
            </div>
          </div>

          <div class="column is-one-third">
            <div class="card">
              <div class="card-content has-text-centered">
                <p class="quote">
                  “I honestly don't believe I've seen any collaboration environment produce this much overall "learning" ever, really.”
                </p>
                <p class="subtitle">
                  - p1xt
                </p>
              </div>
            </div>
          </div>

          <div class="column is-one-third">
            <div class="card">
              <div class="card-content has-text-centered">
                <p class="quote">
                  “Learn with other developers, you start as noob and become a pro with a family of devs.”
                </p>
                <p class="subtitle">
                  - trion129
                </p>
              </div>
            </div>
          </div>

        </div> <!-- columns -->

      </div> <!-- container -->
    </section>

    <section class="section">
      <div class="container has-text-centered">
        <h2 class="title">share your love</h2>

        <div class="columns is-centered">

          <div class="column is-2-tablet">
            <img src="/static/images/cookie-1.png">
          </div>

          <div class="column is-4-tablet">
            <p class="subtitle">running this community is expensive, help us out with a donation!</p>

            <div class="is-inline-flex">
            <div class="field">
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
              <a class="button">donate</a>
            </div>
            </div>
          </div>

          <div class="column is-2-tablet">
            <img src="/static/images/cookie-2.png">
          </div>
        
        </div><!-- columns -->

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
h2 {
  color: #1dabfe;
  font-size: 2rem;
  font-family: 'Raleway', sans-serif;
  font-weight: 700;
  padding: 2rem 0;
}

.chingu {
  background-image: linear-gradient(135deg, #17c4fc 0%, #1fa2ff 100%);
  color: #fff;
}

.chingu h2, .chingu h3 {
  color: #fff;
}

.chingu h2 {
  font-size: 1.4rem;
  padding: 1rem 0;
}

.chingu h3 {
  font-size: 1rem;
  font-weight: 700;
  padding: 1rem 0;
}

.tabs a {
  color: #fff;
}

.tabs a:hover {
    border-bottom: none;
    color: #fff;
}

.card {
  box-shadow: none;
}

.card .card-content .quote {
  color: #1dabfe;
  font-size: 1.5rem;
}

.card .card-content .subtitle {
  color: gray;
  font-style: italic;
}

.field {
  margin-right: 1rem;
}

.field .label {
  color: #fff;
}

.field .control {
  text-align: center;
}

.button {
  background-color: #1dabfe;
  color: #fff;
}

.select select {
  border: 1px solid #1dabfe;
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
