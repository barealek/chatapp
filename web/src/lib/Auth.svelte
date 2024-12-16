<script>
  import { error, json } from "@sveltejs/kit";

   let navn
   let kode
   let errortekst
   function registrer() {
      fetch(`/api/auth/signup`, {
         body: JSON.stringify({
            user: navn,
            password: kode
         }),
         method: "POST"
      }).then(async (res) => {
         if (res.status == 200) {
            errortekst = ""
            window.location.reload()
         } else if (res.status == 403) {

            errortekst = await res.text()
         }
      })
   }
   function login() {
      fetch(`/api/auth/login`, {
         body: JSON.stringify({
            user: navn,
            password: kode
         }),
         method: "POST"
      }).then(async (res) => {
         if (res.status == 200) {
            errortekst = ""
            window.location.reload()
         } else if (res.status == 401) {

            errortekst = await res.text()
         }
      })
   }
</script>

<div class="mt-7 bg-white border border-gray-200 rounded-xl shadow-sm">
  <div class="p-4 sm:p-7">
    <div class="text-center">
      <h1 class="block text-2xl font-bold text-gray-800">Log ind</h1>
    </div>

    <div class="mt-5">
      <!-- Form -->
      <form>
        <div class="grid gap-y-4">
          <!-- Form Group -->
          <div>
            <label for="email" class="block text-sm mb-2">Navn</label>
            <div class="relative">
              <input bind:value={navn} class="py-3 px-4 block w-full border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" required aria-describedby="email-error">
            </div>
          </div>
          <!-- End Form Group -->

          <!-- Form Group -->
          <div>
            <div class="flex justify-between items-center">
              <label for="password" class="block text-sm mb-2">Kode</label>
            </div>
            <div class="relative">
              <input bind:value={kode} type="password" class="py-3 px-4 block w-full border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none" required aria-describedby="password-error">
            </div>
          </div>
          <!-- End Form Group -->

          <button on:click={login} type="submit" class="w-full py-3 px-4 inline-flex justify-center items-center gap-x-2 text-sm font-medium rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 focus:outline-none focus:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none">Log ind</button>
          <button on:click={registrer} type="submit" class="text-gray-900 w-full py-3 px-4 inline-flex justify-center items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-white hover:bg-neutral-200 focus:outline-none focus:bg-neutral-300 disabled:opacity-50 disabled:pointer-events-none">Registrer</button>
          {#if errortekst !== "" && errortekst != null}
          <p class="text-red-500">{errortekst}</p>
          {/if}
        </div>
      </form>
      <!-- End Form -->
    </div>
  </div>
</div>
