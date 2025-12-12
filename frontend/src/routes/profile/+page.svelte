<script lang="ts">
	import * as Empty from '$lib/components/ui/empty/index.js';
	import { LogInIcon } from 'lucide-svelte';
	import type { PageProps } from '../$types';

	const { data }: PageProps = $props();
</script>

{#if !data.session || !data.user}
	<div class="flex h-full w-full items-center justify-center">
		<Empty.Root>
			<Empty.Header>
				<Empty.Media variant="icon">
					<LogInIcon />
				</Empty.Media>
				<Empty.Title>Unauthorized</Empty.Title>
				<Empty.Description>You must be logged in to view your profile.</Empty.Description>
			</Empty.Header>
		</Empty.Root>
	</div>
{:else}
	<h1>Profile</h1>
{/if}
