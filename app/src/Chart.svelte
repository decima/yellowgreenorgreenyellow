<script>
    import {onMount} from "svelte";
    import * as echarts from 'echarts';

    let container = null
    let chart = null;
    export let option = {}
    onMount(() => {

        chart = echarts.init(container)
        chart.setOption(option)
    })
    function onResize(){
        chart.resize();

    }

</script>
<svelte:window
        on:resize={onResize}
/>
<div bind:this={container} class="w-full h-full"></div>
