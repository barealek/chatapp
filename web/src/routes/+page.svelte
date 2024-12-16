<script>
  import Auth from "$lib/Auth.svelte";
  import Chat from "$lib/Chat.svelte";
  import Header from "$lib/Header.svelte";
  import { onMount } from "svelte";

  let user
  let userFetched = false

  onMount(() => {
   fetch("/api/session").then(async (res) => {
      if (res.status == 200) {
         user = await res.json()
      }
      userFetched = true
   })
  })
</script>

{#if userFetched}
   {#if user == null}
      <div class="h-screen w-screen flex items-center justify-center">
         <Auth />
      </div>
   {:else}
      <Header name={user.name} />
      <Chat user={user} />
   {/if}

{:else}
<div class="h-screen w-screen flex items-center justify-center">
   <div
   class="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-current border-e-transparent align-[-0.125em] text-primary motion-reduce:animate-[spin_1.5s_linear_infinite]"
   role="status">
   <span
      class="!absolute !-m-px !h-px !w-px !overflow-hidden !whitespace-nowrap !border-0 !p-0 ![clip:rect(0,0,0,0)]"
      >Loading...</span
   >
   </div>

</div>
{/if}
