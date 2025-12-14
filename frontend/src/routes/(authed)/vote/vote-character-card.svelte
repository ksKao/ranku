<script lang="ts">
	import CharacterImage from '$lib/components/character-image.svelte';
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
<CharacterImage
	class="h-[150px] w-[150px] md:h-[200px] md:w-[200px]"
	src={props.character.image}
	name={props.character.name}
/>
<p class="max-w-[200px] text-center text-muted-foreground">{props.character.anime}</p>
<Button onclick={props.onVote}>Vote</Button>
