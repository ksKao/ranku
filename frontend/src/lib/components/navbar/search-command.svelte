<script lang="ts">
	import * as Command from '$lib/components/ui/command/index.js';
	import { kyClient } from '$lib/ky';
	import { characterSchema } from '$lib/schemas/character.schema';
	import { searchCommandState } from '$lib/states.svelte';
	import { createQuery } from '@tanstack/svelte-query';
	import { LoaderCircleIcon } from 'lucide-svelte';
	import { Debounced } from 'runed';
	import { z } from 'zod/v4';
	import CharacterImage from '$lib/components/character-image.svelte';

	let searchValue = $state('');
	const debouncedSearch = new Debounced(() => searchValue, 1000);

	const searchQuery = createQuery(() => ({
		queryKey: ['searchCharacters', debouncedSearch],
		queryFn: async () => {
			if (!debouncedSearch.current) return [];

			const json = await kyClient('characters', {
				searchParams: {
					query: debouncedSearch.current
				}
			}).json();

			const parsed = z.array(characterSchema).parse(json);

			return parsed;
		}
	}));
</script>

<Command.Dialog bind:open={searchCommandState.open} shouldFilter={false}>
	<Command.Input placeholder="Start typing to search for a character..." bind:value={searchValue} />
	<Command.List>
		{#if debouncedSearch.pending || searchQuery.isPending || searchQuery.isFetching}
			<Command.Loading>
				<div class="flex items-center justify-center p-4">
					<LoaderCircleIcon class="animate-spin" />
				</div>
			</Command.Loading>
		{:else if searchQuery.error}
			<Command.Empty class="text-destructive">Error! Unable to fetch at this time</Command.Empty>
		{:else}
			<Command.Empty>No Results</Command.Empty>
			{#each searchQuery.data as character (character.id)}
				<Command.LinkItem
					href={`/characters/${character.id}`}
					class="gap-4"
					onSelect={() => (searchCommandState.open = false)}
				>
					<CharacterImage
						src={character.image}
						name={character.name}
						class="h-12 w-12 rounded-sm"
					/>
					<div class="flex flex-col">
						<span class="text-lg font-semibold">
							{character.name}
						</span>
						<span class="text-muted-foreground">
							{character.anime}
						</span>
					</div>
				</Command.LinkItem>
			{/each}
		{/if}
	</Command.List>
</Command.Dialog>
