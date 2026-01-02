<script lang="ts">
import {query} from "@urql/svelte";interface Props{vaultId:string}let {vaultId}:Props=$props();
const historyQuery=query({query:`query GetVaultHistory($vaultID:ID!){vaultHistory(vaultID:$vaultID){id amount sourceVaultID destinationVaultID createdAt}}`,variables:{vaultID:vaultId}});
function fmt(d:string){return new Date(d).toLocaleString("en-US",{month:"short",day:"numeric",hour:"2-digit",minute:"2-digit"})}
function isOut(t:any){return t.sourceVaultID===vaultId}
</script>
<div class="bg-white rounded-lg shadow-lg p-6"><h3 class="text-2xl font-bold mb-4">📜 History - #{vaultId}</h3>
{#if $historyQuery.fetching}<div class="text-center py-8"><div class="animate-spin rounded-full h-8 w-8 border-b-2 border-erebor-gold mx-auto"></div></div>
{:else if $historyQuery.error}<div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">{$historyQuery.error.message}</div>
{:else if $historyQuery.data}{#if $historyQuery.data.vaultHistory.length===0}<p class="text-gray-500 text-center py-8">No transfers</p>{:else}
<div class="space-y-3 max-h-96 overflow-y-auto">{#each $historyQuery.data.vaultHistory as t}{@const out=isOut(t)}
<div class="border rounded-lg p-4"><div class="flex justify-between"><div class="flex gap-2"><span class="text-2xl">{out?"📤":"📥"}</span><div><p class="font-semibold text-lg" class:text-red-600={out} class:text-green-600={!out}>{out?"-":"+"}{t.amount.toFixed(2)}</p><p class="text-sm text-gray-600">{out?"To":"From"} #{out?t.destinationVaultID:t.sourceVaultID}</p></div></div><div class="text-right text-xs text-gray-500">{fmt(t.createdAt)}</div></div></div>{/each}</div>{/if}{/if}</div>
