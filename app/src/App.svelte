<script>
    import logo from './assets/logo.svg';
    import {Route, Router} from "svelte-navigator";
    import Home from "./pages/Home.svelte";
    import Stats from "./pages/Stats.svelte";

</script>

<main class="w-screen h-screen flex flex-col">
<div class="flex-grow">
       <Router>
           <Route path="" component={Home} primary={false}/>
           <Route path="stats" component={Stats} primary={false}/>
       </Router>
</div>
<footer class="mb-2 w-full text-center">
    made with ❤️ by <a  class="underline text-accent" href="https://twitter.com/d3cima" target="_blank">@d3cima</a>
</footer>
</main>
<style>

    main {
        overflow: auto;
        background: linear-gradient(-90deg,yellowgreen 50%,greenyellow 50%);
    }
</style>
