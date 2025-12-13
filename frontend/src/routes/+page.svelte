<script lang="ts">
	import { onMount } from 'svelte';
	import { env } from '$env/dynamic/public';

	onMount(() => {
		const endpoint = new URL('/leaderboard/live', env.PUBLIC_BACKEND_URL);
		const eventSource = new EventSource(endpoint.href);

		// Listen for the default 'message' event
		eventSource.addEventListener('message', (event: MessageEvent) => {
			console.log('Received message:', event.data);
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

<h1>Hello</h1>
