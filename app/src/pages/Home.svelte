<script>
    import {Link, navigate} from "svelte-navigator";

    let error = null;

    async function saveVote(kw) {
        try{
       const response= await fetch('/api/store', {
            method: 'POST', // or 'PUT'
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                choice: kw,

            }),
        })
        if (response.status === 200) {
               navigate("stats");
               return
        }
        const content=await response.json()
            error=content.message;
        }catch (e){
            error="Sorry, something went wrong 😭"
        }

    }
</script>
<div class="flex h-full flex-row justify-center items-center">

    <div class="card bg-base-100 shadow-xl">

        <div class="card-body items-center text-center">
            <h1 class="card-title">⚡️⚡️⚡️ Which one do you prefer? ⚡️⚡️⚡️</h1>
            <div class="text-error">
                &nbsp;
                {#if error}

                    {error}
                {/if}
            </div>

            <div class="card-actions flex justify-end items-center">
                <button class="btn btn-primary  " on:click={()=>{saveVote("greenyellow")}}>Greenyellow</button>
                <span class="text-5xl font-bold">VS</span>
                <button class="btn btn-secondary" on:click={()=>{saveVote("yellowgreen")}}>Yellowgreen</button>
            </div>
            <span class=" text-xs">For more fair results, please consider voting once.</span>
            <Link to="stats">
                <div class="btn btn-ghost">see results</div>
            </Link>
        </div>
    </div>
</div>