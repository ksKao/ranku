<script lang="ts">
	import { goto } from '$app/navigation';
	import { authClient } from '$lib/auth-client';

	let loading = $state(false);
	let username = $state('');
	let email = $state('');
	let password = $state('');
</script>

<h1>Register</h1>

<form
	onsubmit={async (e) => {
		e.preventDefault();
		await authClient.signUp.email(
			{
				email,
				username,
				password,
				name: username
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
	<input bind:value={email} placeholder="email" />
	<input bind:value={password} placeholder="password" />
	<button disabled={loading}>
		{loading ? 'Loading' : 'Register'}
	</button>
</form>
<a href="/login">Link to login</a>
