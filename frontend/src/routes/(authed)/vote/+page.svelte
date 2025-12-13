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
	{#if query.isPending}
		<LoaderCircleIcon class="animate-spin" />
	{:else if query.isError}
		<Alert.Root variant="destructive">
			<CircleAlertIcon />
			<Alert.Title>Unable to retrieve vote matchup at this time.</Alert.Title>
		</Alert.Root>
	{:else}
		<div class="flex flex-col items-center gap-8">
			<div class="flex flex-col items-center gap-8 md:flex-row">
				<VoteCharacterCard character={query.data.char1} />
				<span class="text-2xl font-bold">VS</span>
				<VoteCharacterCard character={query.data.char2} />
			</div>
			<Button variant="secondary" class="w-fit" onclick={() => query.refetch()}>
				<RefreshCcwIcon />
				Skip
			</Button>
		</div>
	{/if}
</div>
