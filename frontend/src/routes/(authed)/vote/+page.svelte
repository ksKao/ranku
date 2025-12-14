<script lang="ts">
	import * as Alert from '$lib/components/ui/alert/index';
	import Button from '$lib/components/ui/button/button.svelte';
	import { kyClientAuthed } from '$lib/ky';
	import { matchupSchema } from '$lib/schemas/matchup.schema';
	import { createQuery } from '@tanstack/svelte-query';
	import { CircleAlertIcon, LoaderCircleIcon, RefreshCcwIcon } from 'lucide-svelte';
	import VoteCharacterCard from './vote-character-card.svelte';

	const query = createQuery(() => ({
		queryKey: ['getVoteMatchup'],

		queryFn: async () => {
			const json = await kyClientAuthed.get('votes/matchup').json();

			return matchupSchema.parse(json);
		},

		refetchInterval: false
	}));
</script>

<div class="flex h-full w-full items-center justify-center">
	{#if query.isPending || query.isFetching || query.isRefetching || query.isLoading}
		<LoaderCircleIcon class="animate-spin" />
	{:else if query.isError}
		<Alert.Root variant="destructive" class="w-fit">
			<CircleAlertIcon />
			<Alert.Title>Unable to retrieve vote matchup at this time.</Alert.Title>
		</Alert.Root>
	{:else if !query.data.char1 || !query.data.char2}
		<Alert.Root variant="default" class="w-fit">
			<CircleAlertIcon />
			<Alert.Title>
				You have voted for all of the character matchups! You will be able to vote again once more
				characters are added.
			</Alert.Title>
		</Alert.Root>
	{:else}
		<div class="flex flex-col items-center gap-8">
			<div class="grid grid-cols-1 md:grid-cols-3">
				{#key query.data.char1.id}
					<VoteCharacterCard character={query.data.char1} />
				{/key}
				<div class="my-4 flex items-center justify-center">
					<span class="text-2xl font-bold">VS</span>
				</div>
				{#key query.data.char2.id}
					<VoteCharacterCard character={query.data.char2} />
				{/key}
			</div>
			<Button variant="secondary" class="w-fit" onclick={() => query.refetch()}>
				<RefreshCcwIcon />
				Skip
			</Button>
		</div>
	{/if}
</div>
