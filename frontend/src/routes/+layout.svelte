<script lang="ts">
	import { browser } from '$app/environment';
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/navbar/navbar.svelte';
	import { Toaster } from '$lib/components/ui/sonner/index.js';
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
	import { ModeWatcher } from 'mode-watcher';
	import { type LayoutProps } from './$types';
	import './layout.css';
	import LoginDialog from './login-dialog.svelte';

	const { data, children }: LayoutProps = $props();

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				enabled: browser
			}
		}
	});
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<QueryClientProvider client={queryClient}>
	<ModeWatcher />
	<Toaster richColors position="top-center" />
	<LoginDialog loginForm={data.loginForm} registerForm={data.registerForm} />

	<main class="mx-auto flex h-screen w-screen flex-col px-6 md:max-w-7xl">
		<Navbar session={data.session} user={data.user} />
		<main class="grow">
			{@render children()}
		</main>
	</main>
</QueryClientProvider>
