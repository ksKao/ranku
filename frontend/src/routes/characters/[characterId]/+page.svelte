<script lang="ts">
	import CharacterImage from '$lib/components/character-image.svelte';
	import type { PageProps } from './$types';
	import { format } from 'date-fns';
	import anilistIcon from '$lib/assets/anilist.svg';
	import * as Tooltip from '$lib/components/ui/tooltip/index';
	import Markdown from 'svelte-exmarkdown';
	import type { HTMLAnchorAttributes } from 'svelte/elements';

	const { data }: PageProps = $props();
</script>

<div class="py-4">
	<div class="flex flex-col items-center gap-8 md:flex-row md:items-start">
		<CharacterImage
			src={data.character.image}
			name={data.character.name}
			class="h-[200px] w-[150px]"
		/>
		<div>
			<div class="flex items-center gap-4">
				<h1 class="text-2xl font-bold">
					{data.character.name}
				</h1>
				<Tooltip.Provider delayDuration={0}>
					<Tooltip.Root>
						<Tooltip.Trigger>
							{#snippet child({ props })}
								<a
									{...props}
									href={`https://anilist.co/character/${data.character.anilistId}`}
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
			</div>
			<div class="mt-4 grid grid-cols-[auto_4px_auto] gap-x-8">
				{#if data.character.age && data.character.age.endsWith('-')}
					<b>Initial Age</b>
					<b>:</b>
					<p>{data.character.age.substring(0, data.character.age.length - 1)}</p>
				{:else if data.character.age}
					<b>Age</b>
					<b>:</b>
					<p>{data.character.age}</p>
				{/if}
				{#if data.character.birthYear && data.character.birthMonth && data.character.birthDay}
					<b>Date of Birth</b>
					<b>:</b>
					<p>
						{format(
							new Date(
								data.character.birthYear,
								data.character.birthMonth,
								data.character.birthDay
							),
							'dd MMM yyyy'
						)}
					</p>
				{:else if data.character.birthYear && data.character.birthMonth}
					<b>Date of Birth</b>
					<b>:</b>
					<p>
						{format(new Date(data.character.birthYear, data.character.birthMonth, 0), 'MMM yyyy')}
					</p>
				{:else if data.character.birthYear}
					<b>Year of Birth</b>
					<b>:</b>
					<p>
						{format(new Date(data.character.birthYear, 0, 0), 'yyyy')}
					</p>
				{:else if data.character.birthMonth && data.character.birthDay}
					<b>Birthday</b>
					<b>:</b>
					<p>
						{format(new Date(0, data.character.birthMonth, data.character.birthDay), 'dd MMM')}
					</p>
				{:else if data.character.birthMonth}
					<b>Birth Month</b>
					<b>:</b>
					<p>
						{format(new Date(0, data.character.birthMonth, 0), 'MMMM')}
					</p>
				{/if}
				{#if data.character.bloodType}
					<b>Blood Type</b>
					<b>:</b>
					<p>
						{data.character.bloodType}
					</p>
				{/if}
				{#if data.character.gender}
					<b>Gender</b>
					<b>:</b>
					<p>
						{data.character.gender}
					</p>
				{/if}
			</div>
		</div>
	</div>
	{#if data.character.description}
		<article class="prose mt-6 max-w-full text-muted-foreground dark:prose-invert">
			<Markdown md={data.character.description} {a} />
		</article>
	{/if}
	<p class="mt-6 text-xl font-semibold">Animes:</p>
	<ul class="mt-2 ml-4 list-disc">
		{#each data.character.animes as anime}
			<li>{anime}</li>
		{/each}
	</ul>
</div>

{#snippet a(props: HTMLAnchorAttributes)}
	<a {...props} target="_blank">
		{@render props.children?.()}
	</a>
{/snippet}
