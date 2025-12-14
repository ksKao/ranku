<script lang="ts">
	import { env } from '$env/dynamic/public';
	import CharacterImage from '$lib/components/character-image.svelte';
	import * as Item from '$lib/components/ui/item/index.js';
	import { leaderboardSchema } from '$lib/schemas/leaderboard.schema';
	import { onMount } from 'svelte';
	import { flip } from 'svelte/animate';
	import type { PageProps } from './$types';

	const { data }: PageProps = $props();

	// svelte-ignore state_referenced_locally
	let leaderboard = $state(data.leaderboard);

	onMount(() => {
		const endpoint = new URL('/leaderboard/live', env.PUBLIC_BACKEND_URL);
		const eventSource = new EventSource(endpoint.href);

		// Listen for the default 'message' event
		eventSource.addEventListener('message', (event: MessageEvent) => {
			const parsed = leaderboardSchema.safeParse(JSON.parse(event.data));

			if (parsed.data) {
				leaderboard = parsed.data;
			}
		});

		// Handle errors and connection state
		eventSource.addEventListener('error', (event: Event) => {
			console.error('EventSource error:', event);
		});

		// Cleanup the connection when the component is destroyed
		return () => {
			eventSource.close();
		};
	});
</script>

<ul class="mx-auto max-w-4xl space-y-2 py-4">
	{#each leaderboard as char, i (char.id)}
		<li animate:flip>
			<Item.Root variant="outline">
				<Item.Media class="h-5 w-5">
					{i + 1}
				</Item.Media>
				<Item.Content>
					<div class="flex flex-row gap-4">
						<CharacterImage class="h-8 w-8 rounded-sm" src={char.image} name={char.name} />
						<Item.Title>{char.name}</Item.Title>
					</div>
				</Item.Content>
				<Item.Actions>
					{char.score}
				</Item.Actions>
			</Item.Root>
		</li>
	{/each}
</ul>
