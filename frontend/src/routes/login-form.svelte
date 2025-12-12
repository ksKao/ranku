<script lang="ts">
	import * as Form from '$lib/components/ui/form/index';
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import { loginSchema, type LoginSchema } from './schema';
	import { Input } from '$lib/components/ui/input';
	import { invalidateAll } from '$app/navigation';
	import { toast } from 'svelte-sonner';
	import { loginDialogState } from '$lib/states.svelte';

	type Props = {
		form: SuperValidated<Infer<LoginSchema>>;
	};

	const props: Props = $props();

	const form = superForm(props.form, {
		validators: zod4Client(loginSchema),
		onUpdate: async ({ result }) => {
			if (result.type === 'success') {
				await invalidateAll();
				loginDialogState.open = false;
				toast.success('Login success');
			} else if (result.type === 'failure') {
				if ('message' in result.data) toast.error(result.data.message);
			}
		}
	});

	const { form: formData, enhance } = form;
</script>

<form method="POST" action="/login" use:enhance>
	<div class="my-4 space-y-4">
		<Form.Field {form} name="username">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Username</Form.Label>
					<Input
						{...props}
						bind:value={$formData.username}
						placeholder="Username"
						autocomplete="username"
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="password">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Password</Form.Label>
					<Input
						{...props}
						bind:value={$formData.password}
						placeholder="●●●●●●●●"
						type="password"
						autocomplete="current-password"
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<Form.Button class="mt-4 w-full">Login</Form.Button>
</form>
