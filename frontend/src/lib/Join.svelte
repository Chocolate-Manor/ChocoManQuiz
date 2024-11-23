<script>
    import Button from "./Button.svelte";

    let { joinPin = $bindable(""), handleJoin } = $props();
  
  </script>
  
  <div class="flex items-center justify-center h-screen bg-gray-100">
    <div class="w-full max-w-sm p-6 bg-white border border-gray-300 rounded-lg shadow-md">
      <input
        type="text"
        placeholder="Game PIN"
        bind:value={joinPin}
        class="w-full px-4 py-2 mb-4 text-lg border border-gray-300 rounded focus:ring-2 focus:ring-purple-500 focus:outline-none"
      />
      <Button onclick={handleJoin}>Enter</Button>
    </div>
  </div>