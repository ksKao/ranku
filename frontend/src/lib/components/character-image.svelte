<script lang="ts">
	import { cn } from '$lib/utils';
	import { ImageOffIcon } from 'lucide-svelte';

	type Props = {
		src: string;
		name: string;
		class?: string;
	};

	const props: Props = $props();
</script>

<div class={cn(`relative overflow-hidden rounded-md`, props.class)}>
	{#if props.src.endsWith('default.jpg')}
		<div class="absolute top-0 left-0 flex h-full w-full items-center justify-center bg-muted">
			<ImageOffIcon class="h-1/2 w-1/2 text-muted-foreground" />
		</div>
	{:else}
		<img src={props.src} alt={`${props.name} Image`} class="absolute top-0 left-0 w-full" />
	{/if}
</div>
