<script lang="ts">
	import { goto } from '$app/navigation';
	import { authClient } from '$lib/auth-client';
	import type { PageProps } from './$types';

	const { data }: PageProps = $props();
	const session = authClient.useSession();
</script>

<div>
	{#if $session.isPending || $session.isRefetching}
		<p>Loading</p>
	{:else}
		<div>
			<p>
				Name: {data.user.name}
			</p>
			<p>
				JWT: {data.token}
			</p>
			<button
				onclick={async () => {
					await authClient.signOut();
					await goto('/login');
				}}
			>
				Sign Out
			</button>
		</div>
	{/if}
</div>
