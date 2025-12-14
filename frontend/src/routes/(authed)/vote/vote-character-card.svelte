<script lang="ts">
	import { characterSchema, type CharacterSchema } from '$lib/schemas/character.schema';
	import { ImageOffIcon } from 'lucide-svelte';

	type Props = {
		character: CharacterSchema;
	};

	const props: Props = $props();
</script>

<div class="flex flex-col items-center gap-2 text-lg font-bold">
	<p>{props.character.name}</p>
	<div class="relative h-[200px] w-[200px] overflow-hidden rounded-md">
		{#if props.character.image.endsWith('default.jpg')}
			<div class="absolute top-0 left-0 flex h-full w-full items-center justify-center bg-muted">
				<ImageOffIcon class="text-muted-foreground" />
			</div>
		{:else}
			<img
				src={props.character.image}
				alt={`${props.character.name} Image`}
				class="absolute top-0 left-0 w-full"
			/>
		{/if}
	</div>
	<p class="max-w-[200px] text-center text-muted-foreground">{props.character.anime}</p>
</div>
