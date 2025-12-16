<script lang="ts">
	import anilistIcon from '$lib/assets/anilist.svg';
	import CharacterImage from '$lib/components/character-image.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Tooltip from '$lib/components/ui/tooltip/index';
	import { format } from 'date-fns';
	import { HeartIcon, HeartOffIcon } from 'lucide-svelte';
	import Markdown from 'svelte-exmarkdown';
	import type { HTMLAnchorAttributes } from 'svelte/elements';
	import type { PageProps } from './$types';
	import { createMutation } from '@tanstack/svelte-query';
	import { invalidateAll } from '$app/navigation';
	import { loginDialogState } from '$lib/states.svelte';
	import { kyClientAuthed } from '$lib/ky';
	import { toast } from 'svelte-sonner';

	const { data }: PageProps = $props();

	const character = $derived(data.character);

	const likeMutation = createMutation(() => ({
		mutationFn: async () => {
			if (character.liked) {
				await kyClientAuthed.delete('likes', {
					json: {
						characterId: character.id
					}
				});
			} else {
				await kyClientAuthed.post('likes', {
					json: {
						characterId: character.id
					}
				});
			}
		},
		onSuccess: async () => {
			toast.success(`You have ${character.liked ? 'unliked' : 'liked'} this character.`);
			await invalidateAll();
		}
	}));
</script>

<div class="py-4">
	<div class="flex flex-col items-center gap-8 md:flex-row md:items-start">
		<CharacterImage src={character.image} name={character.name} class="h-[200px] w-[150px]" />
		<div class="w-full">
			<div class="flex items-center gap-4">
				<h1 class="text-2xl font-bold">
					{character.name}
				</h1>
				<Tooltip.Provider delayDuration={0}>
					<Tooltip.Root>
						<Tooltip.Trigger>
							{#snippet child({ props })}
								<a
									{...props}
									href={`https://anilist.co/character/${character.anilistId}`}
									target="_blank"
									class="inline-flex"
								>
									<img src={anilistIcon} alt="anilist-logo" class="h-6 w-6" />
								</a>
							{/snippet}
						</Tooltip.Trigger>
						<Tooltip.Content>
							<p>Link to Anilist</p>
						</Tooltip.Content>
					</Tooltip.Root>
				</Tooltip.Provider>
				<Button
					variant="destructive"
					class="ml-auto"
					loading={likeMutation.isPending}
					onclick={() => {
						if (data.user) {
							likeMutation.mutate();
						} else {
							loginDialogState.open = true;
						}
					}}
				>
					{#if character.liked}
						<HeartOffIcon class="mr-2" />
					{:else}
						<HeartIcon class="mr-2" />
					{/if}
					<span>
						{character.likes}
					</span>
				</Button>
			</div>
			<div class="mt-4 grid grid-cols-[100px_4px_auto] gap-x-8">
				{#if character.age && character.age.endsWith('-')}
					<b>Initial Age</b>
					<b>:</b>
					<p>{character.age.substring(0, character.age.length - 1)}</p>
				{:else if character.age}
					<b>Age</b>
					<b>:</b>
					<p>{character.age}</p>
				{/if}
				{#if character.birthYear && character.birthMonth && character.birthDay}
					<b>Date of Birth</b>
					<b>:</b>
					<p>
						{format(
							new Date(character.birthYear, character.birthMonth, character.birthDay),
							'dd MMM yyyy'
						)}
					</p>
				{:else if character.birthYear && character.birthMonth}
					<b>Date of Birth</b>
					<b>:</b>
					<p>
						{format(new Date(character.birthYear, character.birthMonth, 0), 'MMM yyyy')}
					</p>
				{:else if character.birthYear}
					<b>Year of Birth</b>
					<b>:</b>
					<p>
						{format(new Date(character.birthYear, 0, 0), 'yyyy')}
					</p>
				{:else if character.birthMonth && character.birthDay}
					<b>Birthday</b>
					<b>:</b>
					<p>
						{format(new Date(0, character.birthMonth, character.birthDay), 'dd MMM')}
					</p>
				{:else if character.birthMonth}
					<b>Birth Month</b>
					<b>:</b>
					<p>
						{format(new Date(0, character.birthMonth, 0), 'MMMM')}
					</p>
				{/if}
				{#if character.bloodType}
					<b>Blood Type</b>
					<b>:</b>
					<p>
						{character.bloodType}
					</p>
				{/if}
				{#if character.gender}
					<b>Gender</b>
					<b>:</b>
					<p>
						{character.gender}
					</p>
				{/if}
			</div>
		</div>
	</div>
	{#if character.description}
		<article class="prose mt-6 max-w-full text-muted-foreground dark:prose-invert">
			<Markdown md={character.description} {a} />
		</article>
	{/if}
	<p class="mt-6 text-xl font-semibold">Animes:</p>
	<ul class="mt-2 ml-4 list-disc">
		{#each character.animes as anime}
			<li>{anime}</li>
		{/each}
	</ul>
</div>

{#snippet a(props: HTMLAnchorAttributes)}
	<a {...props} target="_blank">
		{@render props.children?.()}
	</a>
{/snippet}
