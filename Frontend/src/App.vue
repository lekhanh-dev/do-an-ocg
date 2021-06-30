<template>
  <HeaderPage />
  <div class="wrap-router-view">
    <!-- <transition name="fade"> -->
    <router-view></router-view>
    <!-- </transition> -->
  </div>
  <FooterPage />
</template>

<script>
import HeaderPage from "@/components/HeaderPage.vue";
import FooterPage from "@/components/FooterPage.vue";
export default {
  components: {
    HeaderPage,
    FooterPage,
  },
};
</script>
<style lang="scss">
body {
  margin: 0;
  background-color: rgb(247, 247, 247);
}

ul {
  list-style: none;
  padding-left: 0;
}
a {
  text-decoration: none;
}
#app {
  // font-family: Avenir, Helvetica, Arial, sans-serif;
  font-family: "Work Sans", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

$color-primary: black;
$color-secondary: white;

.text-primary {
  color: $color-primary !important;
}

.text-secondary {
  color: $color-secondary !important;
}

@media (max-width: 576px) {
  .container {
    padding: 10px;
  }
}
.wrap-router-view {
  min-height: 100vh;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.4s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
