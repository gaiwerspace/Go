<script lang="ts">
import {mutation,query} from "@urql/svelte";interface Props{onSuccess?:()=>void}let {onSuccess}:Props=$props();let selectedKeeperId=$state("1");let initialDeposit=$state("");let error=$state("");
const keepersQuery=query({query:`query{keepers{id name}}`});
const createMutation=mutation({query:`mutation CreateVault($keeperID:ID!,$initialDeposit:Float!){createVault(keeperID:$keeperID,initialDeposit:$initialDeposit){id}}`});
async function handleSubmit(e:Event){e.preventDefault();error="";const d=parseFloat(initialDeposit);if(isNaN(d)||d<0){error="Invalid";return}
const r=await createMutation.executeMutation({keeperID:selectedKeeperId,initialDeposit:d});if(r.error)error=r.error.message;else{initialDeposit="";if(onSuccess)onSuccess()}}
</script>
<form onsubmit={handleSubmit} class="space-y-4">{#if error}<div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">{error}</div>{/if}
<div><label class="block text-sm font-medium mb-2">Keeper</label>{#if $keepersQuery.data}<select bind:value={selectedKeeperId} class="w-full px-4 py-2 border rounded-lg">{#each $keepersQuery.data.keepers as k}<option value={k.id}>{k.name}</option>{/each}</select>{/if}</div>
<div><label class="block text-sm font-medium mb-2">Deposit</label><input type="number" step="0.01" min="0" bind:value={initialDeposit} required class="w-full px-4 py-2 border rounded-lg"/></div>
<button type="submit" disabled={$createMutation.fetching} class="w-full bg-erebor-gold text-erebor-dark px-6 py-3 rounded-lg font-semibold hover:bg-yellow-500 disabled:opacity-50">{$createMutation.fetching?"Creating...":"Create"}</button>
</form>
