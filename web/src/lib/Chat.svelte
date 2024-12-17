<script>
  import { onMount } from "svelte";
  import Message from "./Message.svelte";

   export let user
   let messages = []

   onMount(()=>{
      let d = setInterval(()=>{
         fetch(`/api/messages`).then(async (res)=>{
            if (res.status==200){
               messages = await res.json()
               // sort messages after the
            }
         })
      }, 500)

      return ()=>{
         clearInterval(d)
      }
   })

   let sendContent
   function sendMessage(e) {
      if (e.key=="Enter") {
         fetch(`/api/messages`, {
            method: "POST",
            body: JSON.stringify({
               msg: sendContent
            })
         }).then(async (res)=>{
            if (res.status==200){
               res = await res.json()
               let msg = res.message
               messages = [...messages, msg]
               sendContent = ""
            }
         })
      }
   }
</script>

<div class="min-h-[2dvh]"></div>
<section class="flex flex-col items-center justify-center w-screen min-h-[95dvh] bg-gray-100 text-gray-800">
<div class="flex flex-col flex-grow w-full max-w-xl overflow-hidden bg-white rounded-lg shadow-xl">
		<div class="flex flex-col flex-grow h-0 p-4 overflow-auto">

         {#each messages as msg}
            <Message messageauthor={msg.author_id} selfid={user?.id} time={msg.sent_at} content={msg.content} messageauthorname={msg.author_name}/>
         {/each}

		</div>

		<div class="p-4 bg-gray-300">
			<input bind:value={sendContent} on:keypress={sendMessage} class="flex items-center w-full h-10 px-3 text-sm rounded" type="text" placeholder="Type your message…">
		</div>
	</div>
</section>
