<script lang="ts">
	import * as Empty from '$lib/components/ui/empty/index';
	import { page } from '$app/state';
	import { BanIcon, GlobeIcon } from 'lucide-svelte';
</script>

<div class="flex h-full w-full items-center justify-center">
	<Empty.Root>
		<Empty.Header>
			{#if page.status === 404}
				<Empty.Media variant="icon">
					<GlobeIcon />
				</Empty.Media>
				<Empty.Title class="text-xl">Oops!</Empty.Title>
				<Empty.Description>We can't seem to find the page you're looking for.</Empty.Description>
			{:else}
				<Empty.Media variant="icon">
					<BanIcon />
				</Empty.Media>
				<Empty.Title class="text-xl">Oops!</Empty.Title>
				<Empty.Description>An unexpected error occured.</Empty.Description>
			{/if}
		</Empty.Header>
	</Empty.Root>
</div>
