<script lang="ts">
	import { authClient } from '$lib/auth-client';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Alert from '$lib/components/ui/alert/index';
	import * as Card from '$lib/components/ui/card/index';
	import * as Empty from '$lib/components/ui/empty/index';
	import * as Field from '$lib/components/ui/field/index';
	import * as Item from '$lib/components/ui/item/index.js';
	import { Input } from '$lib/components/ui/input';
	import { createMutation } from '@tanstack/svelte-query';
	import { CircleAlertIcon, HeartOffIcon, LoaderCircleIcon, LogInIcon } from 'lucide-svelte';
	import { toast } from 'svelte-sonner';
	import type { PageProps } from './$types';
	import CharacterImage from '$lib/components/character-image.svelte';
	import { kyClientAuthed } from '$lib/ky';
	import { invalidateAll } from '$app/navigation';

	const { data }: PageProps = $props();
	const passwordValues = $state({ currentPassword: '', newPassword: '' });

	const updatePasswordMutation = createMutation(() => ({
		mutationFn: async () => {
			const results = await authClient.changePassword({
				...passwordValues
			});

			if (results.error) throw new Error(results.error.message);
		},
		onSuccess: () => {
			toast.success('Password updated successfully');
			passwordValues.currentPassword = '';
			passwordValues.newPassword = '';
		},
		onError: (e) => {
			toast.error(e.message);
		}
	}));

	const unlikeMutation = createMutation(() => ({
		mutationFn: async (characterId: string) => {
			await kyClientAuthed.delete('likes', {
				json: {
					characterId: characterId
				}
			});
		},
		onSuccess: async () => {
			toast.success('You have unliked the character successfully!');
			await invalidateAll();
		},
		onError: () => {
			toast.error('Failed to unlike character');
		}
	}));
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
	<div class="grid gap-4 md:grid-cols-2">
		<Card.Root>
			<Card.Header>
				<Card.Title>Account Information</Card.Title>
			</Card.Header>
			<Card.Content>
				<p><b>Username: </b> {data.user.name}</p>
				<p><b>Email: </b> {data.user.email}</p>
			</Card.Content>
		</Card.Root>
		<Card.Root>
			<Card.Header>
				<Card.Title>Change Password</Card.Title>
			</Card.Header>
			<Card.Content>
				<form
					class="space-y-4"
					onsubmit={(e) => {
						e.preventDefault();
						console.log('submit');
						updatePasswordMutation.mutate();
					}}
				>
					<!-- Put this here to stop warning: https://stackoverflow.com/questions/48150635/autocomplete-new-password-ignored-by-chrome-63-in-windows?rq=3 -->
					<input type="text" name="email" autocomplete="username" class="hidden" />
					<Field.Field>
						<Field.Label for="current-password">Current Password</Field.Label>
						<Input
							bind:value={passwordValues.currentPassword}
							id="current-password"
							type="password"
							placeholder="●●●●●●●●"
							autocomplete="current-password"
						/>
					</Field.Field>
					<Field.Field>
						<Field.Label for="new-password">New Password</Field.Label>
						<Input
							bind:value={passwordValues.newPassword}
							id="new-password"
							type="password"
							placeholder="●●●●●●●●"
							autocomplete="new-password"
						/>
					</Field.Field>
					<Button
						type="submit"
						class="w-full"
						disabled={!passwordValues.currentPassword || !passwordValues.newPassword}
						loading={updatePasswordMutation.isPending}
					>
						Update
					</Button>
				</form>
			</Card.Content>
		</Card.Root>
		<Card.Root class="md:col-span-2">
			<Card.Header>
				<Card.Title>Liked Characters</Card.Title>
			</Card.Header>
			<Card.Content>
				{#if !data.likedCharacters}
					<Alert.Root variant="destructive" class="w-full">
						<CircleAlertIcon />
						<Alert.Title>Failed to get liked characters</Alert.Title>
					</Alert.Root>
				{:else if !data.likedCharacters.length}
					<p class="text-center text-muted-foreground">No liked characters</p>
				{:else}
					<ul class="space-y-2">
						{#each data.likedCharacters as character}
							<li>
								<Item.Root variant="outline">
									<Item.Content>
										<div class="flex flex-row gap-4">
											<CharacterImage
												class="h-8 w-8 rounded-sm"
												src={character.image}
												name={character.name}
											/>
											<Item.Title>
												<a href={`/characters/${character.id}`}>
													{character.name}
												</a>
											</Item.Title>
										</div>
									</Item.Content>
									<Item.Actions>
										<Button
											variant="ghost"
											disabled={unlikeMutation.isPending}
											onclick={() => unlikeMutation.mutate(character.id)}
										>
											{#if unlikeMutation.isPending}
												<LoaderCircleIcon class="animate-spin" />
											{:else}
												<HeartOffIcon />
											{/if}
										</Button>
									</Item.Actions>
								</Item.Root>
							</li>
						{/each}
					</ul>
				{/if}
			</Card.Content>
		</Card.Root>
	</div>
{/if}
