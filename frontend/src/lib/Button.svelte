<script>
  let {onclick, children} = $props();
</script>

<button
  class="px-6 py-2 rounded-lg bg-gradient-to-r from-indigo-500 to-purple-600 text-white font-semibold text-sm md:text-base shadow-lg hover:from-purple-600 hover:to-indigo-500 transition duration-200 ease-in-out transform hover:scale-105 active:scale-100 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
 {onclick}>
<!-- {text} -->
{@render children?.()}
</button>
