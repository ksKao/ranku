<script lang="ts">
	import * as FormPrimitive from 'formsnap';
	import { cn, type WithoutChild } from '$lib/utils.js';

	let {
		ref = $bindable(null),
		class: className,
		errorClasses,
		children: childrenProp,
		single = true,
		...restProps
	}: WithoutChild<FormPrimitive.FieldErrorsProps> & {
		single?: boolean;
		errorClasses?: string | undefined | null;
	} = $props();
</script>

<FormPrimitive.FieldErrors
	bind:ref
	class={cn('text-sm font-medium text-destructive', className)}
	{...restProps}
>
	{#snippet children({ errors, errorProps })}
		{#if childrenProp}
			{@render childrenProp({ errors, errorProps })}
		{:else if !single}
			{#each errors as error (error)}
				<div {...errorProps} class={cn(errorClasses)}>{error}</div>
			{/each}
		{:else}
			<div {...errorProps} class={cn(errorClasses)}>{errors[0]}</div>
		{/if}
	{/snippet}
</FormPrimitive.FieldErrors>
