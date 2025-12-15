<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';
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
	import { loginDialogState, searchCommandState } from '$lib/states.svelte';
	import type { Session, User } from 'better-auth';
	import { MenuIcon } from 'lucide-svelte';
	import { toast } from 'svelte-sonner';
	import ThemeSwitcher from './theme-switcher.svelte';
	import SearchCommand from './search-command.svelte';

	type Props = {
		user: User | undefined;
		session: Session | undefined;
	};

	const props: Props = $props();

	type NavItem = (
		| {
				text: string;
		  }
		| { image: string; alt: string }
	) &
		({ link: string } | { onClick: () => void });

	const navItems: NavItem[] = $derived([
		{
			text: 'Vote',
			onClick: () => {
				if (props.user) {
					goto('/vote');
				} else {
					loginDialogState.open = true;
				}
			}
		},
		{
			text: 'Search',
			onClick: () => {
				searchCommandState.open = true;
			}
		},
		{ image: favicon, alt: 'logo', link: '/' },
		{
			text: 'Profile',
			onClick: () => {
				if (props.user) {
					goto('/profile');
				} else {
					loginDialogState.open = true;
				}
			}
		},
		{
			text: props.session ? 'Log Out' : 'Login',
			onClick: async () => {
				if (props.session) {
					await authClient.signOut();
					await invalidateAll();
					toast.success('Logged out successfully');
				} else {
					loginDialogState.open = true;
				}
			}
		}
	]);
</script>

<SearchCommand />
<header class="sticky top-0 z-50 bg-background">
	<div class="relative flex w-full items-center justify-between gap-8 py-7">
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
						<div class="flex items-center gap-4 font-bold text-foreground">
							<img class="h-8 w-8" src={navItem.image} alt={navItem.alt} />
							<span>Leaderboard</span>
						</div>
					{/if}
				</a>
			{/each}
		</div>

		<div class="mr-auto md:hidden">
			<DropdownMenu>
				<DropdownMenuTrigger>
					{#snippet child({ props })}
						<Button {...props} variant="outline" size="icon">
							<MenuIcon />
							<span class="sr-only">Menu</span>
						</Button>
					{/snippet}
				</DropdownMenuTrigger>
				<DropdownMenuContent class="w-48" align="start">
					<DropdownMenuGroup>
						<DropdownMenuItem
							onSelect={() => {
								goto('/');
							}}
						>
							Home
						</DropdownMenuItem>
						{#each navItems as navItem}
							{#if 'text' in navItem}
								<DropdownMenuItem
									onSelect={() => {
										if ('link' in navItem) {
											goto(navItem.link);
										} else {
											navItem.onClick();
										}
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
		<div class="absolute top-1/2 right-0 -translate-y-1/2">
			<ThemeSwitcher />
		</div>
	</div>
</header>
