<script lang="ts">
  import svelteLogo from './assets/svelte.svg'
  import viteLogo from '/vite.svg'
  import Counter from './lib/Counter.svelte'
  import Button from './lib/Button.svelte';
  import QuizCard from './lib/QuizCard.svelte';
  import Join from './lib/Join.svelte';

  let pin = $state("");
  let msg = $state("");
  let quizzes: {_id: string, name: string}[] = $state([]);

  async function getQuizzes(){
    let response = await fetch("http://localhost:3000/api/quizzes");
    if(!await response.ok){
      alert("Failed!");
      return;
    }

    let json = await response.json();
    quizzes = json;
    console.log(json);
  }
  
  function join(){
    let websocket = new WebSocket("ws://localhost:3000/ws");
    websocket.onopen = () => {
      websocket.send(`join:${pin}`);
    };
  }

  function host(quiz){
    let websocket = new WebSocket("ws://localhost:3000/ws");
    websocket.onopen = () => {
      websocket.send(`host:${quiz.id}`);
    };

    websocket.onmessage = (event) => {
      msg = event.data;
    }
  }
</script>

<Button onclick={getQuizzes}>getQuizzes</Button>
Message: {msg}

<Join bind:joinPin={pin} handleJoin={join}></Join>

{#each quizzes as quiz}
  <QuizCard onhost={() => host(quiz)} {quiz}/>
{/each}


