<script>
    import {Link} from "svelte-navigator";
    import Chart from "../Chart.svelte";
    import {onMount} from "svelte";

    let stats = null;
    let loaded = false
    let sum=0;
    let pieRepartition = [];
    let pieColorOptions = [];
    let oneMonthHistory = [];
    let oneMonthHistoryXAxis = [];
    let perMonthHistory = [];
    let perMonthHistoryXAxis = [];

    let best = "yellowgreen";
    let worst = "greenyellow";
    let equals=false;

    onMount(async () => {
        stats = (await (await fetch("/api/stats")).json()).stats
        if (stats.split.greenyellow > stats.split.yellowgreen) {
            best = "greenyellow";
            worst = "yellowgreen"
        }else if(stats.split.greenyellow===stats.split.yellowgreen){
            equals=true;
        }
        for (const item of Object.entries(stats.split)) {
            pieRepartition.push({name: item[0], value: item[1]})
            pieColorOptions.push(item[0])
            sum+=item[1];
        }

        for (const item of Object.entries(stats.per_last_30_days)) {
            oneMonthHistory.push({
                name: item[0],
                data: item[1],
                type: 'line',
                smooth: true,
                lineStyle: {
                    normal: {
                        color: item[0],
                    }
                },
                itemStyle: {
                    color: item[0],
                }

            })
        }
        oneMonthHistoryXAxis = stats.per_last_30_days_x_axis

        for (const item of Object.entries(stats.per_month)) {
            perMonthHistory.push({
                name: item[0],
                data: item[1],
                type: 'line',
                smooth: true,
                lineStyle: {
                    normal: {
                        color: item[0],
                    }
                },
                itemStyle: {
                    color: item[0],
                }

            })
        }
        perMonthHistoryXAxis = stats.per_month_x_axis


        loaded = true
    })

    function oneMonthChart() {
        const res = {
            legend: {
                orient: 'horizontal',
                left: 0,
                top: 10
            },
            xAxis: {data: oneMonthHistoryXAxis}, yAxis: {type: "value"}, series: oneMonthHistory
        }
        return res;
    }

    function perMonthChart() {
        const res = {
            legend: {
                orient: 'horizontal',
                left: 0,
                top: 10
            },
            xAxis: {data: perMonthHistoryXAxis}, yAxis: {type: "value"}, series: perMonthHistory
        }
        return res;
    }

</script>

<div class="card bg-base-100 shadow-xl container p-4 m-4 mx-auto">
    <div class="w-full">
        <Link to="/">
            <div class="btn btn-ghost">← Back to vote</div>
        </Link>
    </div>

    <h1 class="text-5xl font-bold">Results</h1>
    {#if loaded}
        <div>
            <div class="flex flex-col md:flex-row">
                <div class="w-full md:w-1/2">
                    <h2 class="text-2xl">Which one is the best? <span class="text-xs">(according to {sum} votes)</span></h2>
                    <div class="w-full  h-56 md:h-96">
                        <Chart option={{color: pieColorOptions,series: [{type: 'pie',data: pieRepartition}]}}/>
                    </div>
                </div>
                <div class="w-full md:w-1/2">
                    <h2 class="text-2xl">Cumulative sum for last 30 days</h2>
                    <div class="w-full h-56 md:h-96">
                        <Chart option={oneMonthChart()}/>
                    </div>
                </div>
            </div>
            <h2 class="text-2xl">All-time trend</h2>
            <div class=" h-56 md:h-96">
                <Chart option={perMonthChart()}/>

            </div>

            <div class="grid h-56 md:h-96 place-items-center">
                <p>
                    {#if (equals)}
                        <b class="text-4xl" style="color: black;">No one</b> wins 😭.
                    {:else }
                    <b class="text-4xl" style="color: {best};">🎉 {best} 🎉</b> wins.

                    {/if}
                </p>
            </div>


        </div>
    {:else }
        Loading...
    {/if}

</div>