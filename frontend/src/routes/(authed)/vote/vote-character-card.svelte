<script lang="ts">
	import type { CharacterSchema } from '$lib/schemas/character.schema';

	type Props = {
		character: CharacterSchema;
	};

	const props: Props = $props();
</script>

<div class="flex flex-col items-center gap-2 text-lg font-bold">
	<p>{props.character.name}</p>
	<div class="relative h-[200px] w-[200px] overflow-hidden rounded-md">
		<img
			src={props.character.image}
			alt={`${props.character.name} Image`}
			class="absolute top-0 left-0 w-full"
		/>
	</div>
	<!-- <p class="text-muted-foreground">{props.character.}</p> -->
</div>
