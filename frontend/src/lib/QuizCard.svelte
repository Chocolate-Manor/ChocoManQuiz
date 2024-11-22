<script lang="ts">
    import Button from "./Button.svelte";
    //let {quiz, host} = $props();
    let { quiz, onhost }: { quiz: {_id: string, name: string}, onhost:Function } = $props();
</script>

<div class="bg-white rounded-lg shadow-lg p-6 mb-6 transition-transform transform hover:scale-105">
    <div class="flex items-center justify-between">
        <div class="flex-1 text-center">
            <h2 class="text-3xl font-bold text-gray-800">{quiz.name}</h2> <!-- Big title -->
            <p class="text-lg text-gray-600 mt-2">Join the fun!</p> <!-- Subtitle or description -->
        </div>
            <Button onclick={onhost}>Host</Button>
    </div>
</div>

