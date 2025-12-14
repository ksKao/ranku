<script lang="ts">
	import { onMount } from 'svelte';
	import { env } from '$env/dynamic/public';
	import type { PageProps } from './$types';
	import { leaderboardSchema } from '$lib/schemas/leaderboard.schema';
	import { flip } from 'svelte/animate';

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

<ul class="mx-auto max-w-4xl space-y-2">
	{#each leaderboard.slice(0, 10) as char, i (char.id)}
		<li animate:flip class="flex justify-between">
			<p class="w-6">
				{i + 1}
			</p>
			<p class="grow">
				{char.name}
			</p>
			<p>
				{char.score}
			</p>
		</li>
	{/each}
</ul>
