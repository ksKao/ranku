<script lang="ts">
	import { goto } from '$app/navigation';
	import { authClient } from '$lib/auth-client';

	let loading = $state(false);
	let username = $state('');
	let password = $state('');
</script>

<h1>Login</h1>

<form
	onsubmit={async (e) => {
		e.preventDefault();
		await authClient.signIn.username(
			{
				username,
				password
			},
			{
				onRequest: () => {
					loading = true;
				},
				onSuccess: async () => {
					await goto('/');
				},
				onError: (e) => {
					alert(e.error.message);
				},
				onResponse: () => {
					loading = false;
				}
			}
		);
	}}
>
	<input bind:value={username} placeholder="username" />
	<input bind:value={password} placeholder="password" />
	<button disabled={loading}>
		{loading ? 'Loading' : 'Login'}
	</button>
</form>
<a href="/register">Link to register</a>
