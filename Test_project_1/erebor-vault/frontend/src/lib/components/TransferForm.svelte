<script lang="ts">
import {mutation,query} from "@urql/svelte";interface Props{onSuccess?:()=>void}let {onSuccess}:Props=$props();let src=$state("");let dst=$state("");let amount=$state("");let error=$state("");
const keepersQuery=query({query:`query{keepers{id name vaults{id balance}}}`});
const transferMutation=mutation({query:`mutation TransferGold($sourceVaultID:ID!,$destinationVaultID:ID!,$amount:Float!){transferGold(sourceVaultID:$sourceVaultID,destinationVaultID:$destinationVaultID,amount:$amount){id}}`});
async function handleSubmit(e:Event){e.preventDefault();error="";if(!src||!dst){error="Select vaults";return}if(src===dst){error="Different vaults";return}
const a=parseFloat(amount);if(isNaN(a)||a<=0){error="Invalid";return}
const r=await transferMutation.executeMutation({sourceVaultID:src,destinationVaultID:dst,amount:a});if(r.error)error=r.error.message;else{src="";dst="";amount="";if(onSuccess)onSuccess()}}
function lbl(k:any,v:any){return`${k.name} - #${v.id} (${v.balance.toFixed(2)})`}
</script>
<form onsubmit={handleSubmit} class="space-y-4">{#if error}<div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">{error}</div>{/if}
{#if $keepersQuery.data}<div><label class="block text-sm font-medium mb-2">From</label><select bind:value={src} required class="w-full px-4 py-2 border rounded-lg"><option value="">Select</option>{#each $keepersQuery.data.keepers as k}{#each k.vaults as v}<option value={v.id}>{lbl(k,v)}</option>{/each}{/each}</select></div>
<div><label class="block text-sm font-medium mb-2">To</label><select bind:value={dst} required class="w-full px-4 py-2 border rounded-lg"><option value="">Select</option>{#each $keepersQuery.data.keepers as k}{#each k.vaults as v}<option value={v.id} disabled={v.id===src}>{lbl(k,v)}</option>{/each}{/each}</select></div>{/if}
<div><label class="block text-sm font-medium mb-2">Amount</label><input type="number" step="0.01" min="0.01" bind:value={amount} required class="w-full px-4 py-2 border rounded-lg"/></div>
<button type="submit" disabled={$transferMutation.fetching} class="w-full bg-blue-600 text-white px-6 py-3 rounded-lg font-semibold hover:bg-blue-700 disabled:opacity-50">{$transferMutation.fetching?"Transferring...":"Transfer"}</button>
</form>
