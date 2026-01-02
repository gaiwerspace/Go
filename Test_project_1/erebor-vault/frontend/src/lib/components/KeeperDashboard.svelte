<script lang="ts">
import {query} from "@urql/svelte";import VaultHistory from "./VaultHistory.svelte";interface Props{keeperId:string}let {keeperId}:Props=$props();let selectedVaultId:string|null=$state(null);
const keeperQuery=query({query:`query GetKeeper($id:ID!){keeper(id:$id){id name vaults{id balance createdAt}totalBalance}}`,variables:{id:keeperId}});
function fmt(d:string){return new Date(d).toLocaleDateString("en-US",{year:"numeric",month:"short",day:"numeric"})}
</script>
<div>{#if $keeperQuery.fetching}<div class="text-center py-12"><div class="animate-spin rounded-full h-12 w-12 border-b-2 border-erebor-gold mx-auto"></div></div>
{:else if $keeperQuery.error}<div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">{$keeperQuery.error.message}</div>
{:else if $keeperQuery.data?.keeper}{@const k=$keeperQuery.data.keeper}
<div class="bg-white rounded-lg shadow-lg p-8 mb-6"><div class="flex justify-between"><div><h2 class="text-3xl font-bold mb-2">{k.name}</h2><p class="text-gray-600">ID:{k.id}</p></div><div class="text-right"><p class="text-sm text-gray-600">Total</p><p class="text-4xl font-bold text-erebor-gold">{k.totalBalance.toFixed(2)}</p><p class="text-sm text-gray-500">gold</p></div></div></div>
<h3 class="text-2xl font-bold text-erebor-light mb-4">Vaults</h3>
{#if k.vaults.length===0}<div class="bg-white rounded-lg shadow p-6 text-center text-gray-500">No vaults</div>{:else}
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">{#each k.vaults as v}<div onclick={()=>selectedVaultId=v.id} class="bg-white rounded-lg shadow-lg p-6 hover:shadow-xl cursor-pointer" class:border-2={selectedVaultId===v.id} class:border-erebor-gold={selectedVaultId===v.id}><div class="flex justify-between mb-4"><div><p class="text-sm text-gray-600">Vault</p><p class="font-mono text-lg font-semibold">#{v.id}</p></div><span class="text-3xl">🏺</span></div><div class="mb-3"><p class="text-sm text-gray-600">Balance</p><p class="text-3xl font-bold text-erebor-gold">{v.balance.toFixed(2)}</p></div><div class="pt-3 border-t"><p class="text-xs text-gray-500">{fmt(v.createdAt)}</p></div></div>{/each}</div>{/if}
{#if selectedVaultId}<div class="mt-6"><VaultHistory vaultId={selectedVaultId}/></div>{/if}
{/if}</div>
