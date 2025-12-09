<script lang="ts">
	import { goto } from '$app/navigation';
	import { authClient } from '$lib/auth-client';
	import type { PageProps } from './$types';

	const { data }: PageProps = $props();
</script>

<div>
	<div>
		<p>
			Name: {data.user.name}
		</p>
		<p>
			ID: {data.user.id}
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
</div>
