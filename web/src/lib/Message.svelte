<script>
  export let messageauthor, selfid, time, content, messageauthorname;

  export let sending = false

  var date = new Date(time)
</script>

{#if messageauthor == selfid}
<div class="flex justify-end w-full max-w-xs mt-2 ml-auto space-x-3">
  <div>
    <div class="p-3 text-white {!sending ? 'bg-blue-600' : 'bg-blue-600/75'} rounded-l-lg rounded-br-lg">
      <p class="text-sm">
         {content}
      </p>
    </div>
    <span class="text-xs leading-none text-gray-500">{messageauthorname} - {date.toLocaleString()}</span>
  </div>
  <!-- <div class="flex-shrink-0 w-10 h-10 bg-gray-300 rounded-full"></div> -->
</div>
{:else}
<div class="flex w-full max-w-xs mt-2 space-x-3">
  <!-- <div class="flex-shrink-0 w-10 h-10 bg-gray-300 rounded-full"></div> -->
  <div>
    <div class="p-3 bg-gray-300 rounded-r-lg rounded-bl-lg">
      <p class="text-sm">
        {content}
      </p>
    </div>
    <span class="text-xs leading-none text-gray-500">{messageauthorname} - {date.toLocaleString()}</span>
  </div>
</div>
{/if}
