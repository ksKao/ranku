<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import * as Form from '$lib/components/ui/form/index';
	import { Input } from '$lib/components/ui/input';
	import { registerSchema, type RegisterSchema } from '$lib/schemas/register.schema';
	import { loginDialogState } from '$lib/states.svelte';
	import { toast } from 'svelte-sonner';
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';

	type Props = {
		form: SuperValidated<Infer<RegisterSchema>>;
	};

	const props: Props = $props();

	const form = superForm(props.form, {
		validators: zod4Client(registerSchema),
		onUpdate: async ({ result }) => {
			if (result.type === 'success') {
				await invalidateAll();
				loginDialogState.open = false;
				toast.success('Account registered');
			} else if (result.type === 'failure') {
				if ('message' in result.data) toast.error(result.data.message);
			}
		}
	});

	const { form: formData, enhance } = form;
</script>

<form method="POST" action="/register" use:enhance>
	<div class="my-4 space-y-4">
		<Form.Field {form} name="email">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Email</Form.Label>
					<Input
						{...props}
						bind:value={$formData.email}
						placeholder="test@example.com"
						autocomplete="email"
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
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
						autocomplete="new-password"
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="confirmPassword">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Confirm Password</Form.Label>
					<Input
						{...props}
						bind:value={$formData.confirmPassword}
						placeholder="●●●●●●●●"
						type="password"
						autocomplete="new-password"
					/>
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<Form.Button class="mt-4 w-full">Register</Form.Button>
</form>
