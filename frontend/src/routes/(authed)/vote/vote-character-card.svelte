<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import { type CharacterSchema } from '$lib/schemas/character.schema';
	import { ImageOffIcon } from 'lucide-svelte';

	type Props = {
		character: CharacterSchema;
		loading: boolean;
		onVote: () => void;
	};

	const props: Props = $props();
</script>

<p class="text-center">{props.character.name}</p>
<div class="relative h-[150px] w-[150px] overflow-hidden rounded-md md:h-[200px] md:w-[200px]">
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
<Button onclick={props.onVote}>Vote</Button>
