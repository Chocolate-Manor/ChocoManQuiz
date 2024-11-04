<script lang="ts">
    export let quiz: {_id: string, name: string};
</script>

<div class="bg-white rounded-lg shadow-lg p-6 mb-6 transition-transform transform hover:scale-105">
    <h2 class="text-3xl font-bold text-center text-gray-800">{quiz.name}</h2> <!-- Big title -->
    <p class="text-lg text-gray-600 text-center mt-2">Join the fun!</p> <!-- Subtitle or description -->
</div>