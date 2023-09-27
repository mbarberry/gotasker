<template>
  <div class="flex flex-col min-h-screen justify-between">
    <div class="shadow-2xl p-8 m-8 basis-4/5">
      <Tasker msg="Welcome to Go Tasker!" />
    </div>

    <footer class="basis-1/5 h-52 bg-gopher flex justify-center items-center ">
      <div class="flex gap-2">
        <p class="attribution">Built and maintained with</p>
        <svg xmlns="http://www.w3.org/2000/svg" fill="white" height="1.5em" viewBox="0 0 512 512">
          <path
            d="M47.6 300.4L228.3 469.1c7.5 7 17.4 10.9 27.7 10.9s20.2-3.9 27.7-10.9L464.4 300.4c30.4-28.3 47.6-68 47.6-109.5v-5.8c0-69.9-50.5-129.5-119.4-141C347 36.5 300.6 51.4 268 84L256 96 244 84c-32.6-32.6-79-47.5-124.6-39.9C50.5 55.6 0 115.2 0 185.1v5.8c0 41.5 17.2 81.2 47.6 109.5z" />
        </svg>
        <a class="attribution" href="https://github.com/mikebarberry" title="my github" target="_blank">Mike Barberry</a>
      </div>
    </footer>
  </div>
</template>

<script>
import Tasker from './components/Tasker.vue';

export default {
  name: 'App',
  components: {
    Tasker,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
