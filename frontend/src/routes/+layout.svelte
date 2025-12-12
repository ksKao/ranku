<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import favicon from '$lib/assets/favicon.svg';
	import { authClient } from '$lib/auth-client';
	import { Button } from '$lib/components/ui/button';
	import {
		DropdownMenu,
		DropdownMenuContent,
		DropdownMenuGroup,
		DropdownMenuItem,
		DropdownMenuTrigger
	} from '$lib/components/ui/dropdown-menu';
	import { MenuIcon } from 'lucide-svelte';
	import { type LayoutProps } from './$types';
	import './layout.css';
	import LoginDialog from './login-dialog.svelte';
	import { Toaster } from '$lib/components/ui/sonner/index.js';
	import { toast } from 'svelte-sonner';

	const { data, children }: LayoutProps = $props();
	let loginDialogOpen = $state(false);

	type NavItem = (
		| {
				text: string;
		  }
		| { image: string; alt: string }
	) &
		({ link: string } | { onClick: () => void });

	const navItems: NavItem[] = $derived([
		{ text: 'Vote', link: '/vote' },
		{ text: 'Search', onClick: () => {} },
		{ image: favicon, alt: 'logo', link: '/' },
		{ text: 'Profile', link: '/profile' },
		{
			text: data.session ? 'Log Out' : 'Login',
			onClick: async () => {
				if (data.session) {
					await authClient.signOut();
					await invalidateAll();
					toast.success('Logged out successfully');
				} else {
					loginDialogOpen = true;
				}
			}
		}
	]);
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<Toaster richColors position="top-center" />
<LoginDialog
	open={loginDialogOpen}
	onOpenChange={(open) => {
		loginDialogOpen = open;
	}}
	loginForm={data.loginForm}
	registerForm={data.registerForm}
/>
<header class="sticky top-0 z-50 bg-background">
	<div class="mx-auto flex w-full items-center justify-between gap-8 px-6 py-7 md:max-w-7xl">
		<div
			class="hidden flex-1 items-center gap-8 font-medium text-muted-foreground md:flex md:justify-center lg:gap-16"
		>
			{#each navItems as navItem}
				<a
					href={'link' in navItem ? navItem.link : ''}
					class="hover:text-primary"
					onclick={'onClick' in navItem
						? (e) => {
								e.preventDefault();
								navItem.onClick();
							}
						: undefined}
				>
					{#if 'text' in navItem}
						{navItem.text}
					{:else}
						<img class="h-8 w-8" src={navItem.image} alt={navItem.alt} />
					{/if}
				</a>
			{/each}
		</div>

		<div class="ml-auto flex items-center gap-6">
			<DropdownMenu>
				<DropdownMenuTrigger class="md:hidden">
					{#snippet child({ props })}
						<Button {...props} variant="outline" size="icon">
							<MenuIcon />
							<span class="sr-only">Menu</span>
						</Button>
					{/snippet}
				</DropdownMenuTrigger>
				<DropdownMenuContent class="w-48" align="end">
					<DropdownMenuGroup>
						{#each navItems as navItem}
							{#if 'text' in navItem}
								<DropdownMenuItem
									onSelect={() => {
										console.log('selected');
									}}
								>
									{navItem.text}
								</DropdownMenuItem>
							{/if}
						{/each}
					</DropdownMenuGroup>
				</DropdownMenuContent>
			</DropdownMenu>
		</div>
	</div>
</header>
<main class="px-6">
	{@render children()}
</main>
