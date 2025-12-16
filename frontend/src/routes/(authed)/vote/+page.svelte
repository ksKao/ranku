<script lang="ts">
	import * as Alert from '$lib/components/ui/alert/index';
	import Button from '$lib/components/ui/button/button.svelte';
	import { kyClientAuthed } from '$lib/ky';
	import { matchupSchema } from '$lib/schemas/matchup.schema';
	import { createMutation, createQuery } from '@tanstack/svelte-query';
	import { CircleAlertIcon, LoaderCircleIcon, RefreshCcwIcon } from 'lucide-svelte';
	import { toast } from 'svelte-sonner';
	import VoteCharacterCard from './vote-character-card.svelte';

	const query = createQuery(() => ({
		queryKey: ['getVoteMatchup'],

		queryFn: async () => {
			const json = await kyClientAuthed.get('votes/matchup').json();

			return matchupSchema.parse(json);
		},

		refetchInterval: false
	}));

	const createVoteMutation = createMutation(() => ({
		mutationFn: async ({
			forCharacterId,
			againstCharacterId
		}: {
			forCharacterId: string;
			againstCharacterId: string;
		}) => {
			await kyClientAuthed.post('votes', {
				json: {
					forCharacterId,
					againstCharacterId
				}
			});
		},
		onSuccess: () => {
			query.refetch();
		},
		onError: () => {
			toast.error('Failed to vote for character. Please try again later.');
		}
	}));

	const char1 = $derived(query.data?.char1);
	const char2 = $derived(query.data?.char2);
</script>

<div class="flex h-full w-full items-center justify-center">
	{#if query.isPending || query.isFetching || query.isRefetching || query.isLoading}
		<LoaderCircleIcon class="animate-spin" />
	{:else if query.isError}
		<Alert.Root variant="destructive" class="w-fit">
			<CircleAlertIcon />
			<Alert.Title>Unable to retrieve vote matchup at this time.</Alert.Title>
		</Alert.Root>
	{:else if !char1 || !char2}
		<Alert.Root variant="default" class="w-fit">
			<CircleAlertIcon />
			<Alert.Title>
				You have voted for all of the character matchups! You will be able to vote again once more
				characters are added.
			</Alert.Title>
		</Alert.Root>
	{:else}
		<div class="flex flex-col items-center gap-8">
			<div class="grid grid-flow-col grid-rows-[auto_auto_auto_auto] place-items-center gap-4">
				{#key char1.id}
					<VoteCharacterCard
						character={char1}
						loading={createVoteMutation.isPending}
						onVote={() => {
							createVoteMutation.mutate({
								forCharacterId: char1.id,
								againstCharacterId: char2.id
							});
						}}
					/>
				{/key}
				<div></div>
				<div class="my-4 flex items-center justify-center md:px-8">
					<span class="text-2xl font-bold">VS</span>
				</div>
				<div></div>
				<div></div>
				{#key char2}
					<VoteCharacterCard
						character={char2}
						loading={createVoteMutation.isPending}
						onVote={() => {
							createVoteMutation.mutate({
								forCharacterId: char2.id,
								againstCharacterId: char1.id
							});
						}}
					/>
				{/key}
			</div>
			<Button variant="secondary" class="w-fit" onclick={() => query.refetch()}>
				<RefreshCcwIcon />
				Skip
			</Button>
		</div>
	{/if}
</div>
