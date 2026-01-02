<script lang="ts">
import {query} from "@urql/svelte";interface Props{onSelectKeeper:(id:string)=>void}let {onSelectKeeper}:Props=$props();
const keepersQuery=query({query:`query GetKeepers{keepers{id name totalBalance}}`});
</script>
<div><h2 class="text-3xl font-bold text-erebor-light mb-6">Keepers</h2>
{#if $keepersQuery.fetching}<div class="text-center py-12"><div class="animate-spin rounded-full h-12 w-12 border-b-2 border-erebor-gold mx-auto"></div></div>
{:else if $keepersQuery.error}<div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">{$keepersQuery.error.message}</div>
{:else if $keepersQuery.data}<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
{#each $keepersQuery.data.keepers as keeper}<div onclick={()=>onSelectKeeper(keeper.id)} class="bg-white rounded-lg shadow-lg p-6 hover:shadow-xl cursor-pointer hover:border-erebor-gold border-2"><div class="text-4xl mb-3 text-center">👤</div><h3 class="text-xl font-bold text-center mb-2">{keeper.name}</h3><div class="text-center"><p class="text-sm text-gray-600">Balance</p><p class="text-2xl font-bold text-erebor-gold">{keeper.totalBalance.toFixed(2)}</p><p class="text-xs text-gray-500">gold</p></div></div>{/each}
</div>{/if}</div>
